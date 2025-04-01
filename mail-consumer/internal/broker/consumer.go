package broker

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"mail-consumer/internal/mail"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl"
	"github.com/twmb/franz-go/pkg/sasl/scram"
)

type Consumer struct {
	broker string
	log    *logrus.Logger
	sender *mail.Sender
}

func NewConsumer(broker string, log *logrus.Logger, sender *mail.Sender) *Consumer {
	return &Consumer{broker: broker, log: log, sender: sender}
}

func (c *Consumer) Consume() {
	c.log.Info("Connecting to broker: ", c.broker)

	// Configure SASL with SCRAM-SHA-256
	scramMechanism := scram.Auth{
		User: os.Getenv("SASL_USERNAME"),
		Pass: os.Getenv("SASL_PASSWORD"),
	}.AsSha256Mechanism()

	client, err := kgo.NewClient(
		kgo.SeedBrokers(c.broker),
		kgo.ConsumeTopics("emails"),
		kgo.ConsumerGroup("mail-consumer-group"),                 // Specify a consumer group
		kgo.SASL(sasl.Mechanism(scramMechanism)),                 // Use SCRAM-SHA-256
		kgo.DialTLSConfig(&tls.Config{InsecureSkipVerify: true}), // Enable TLS
		kgo.AutoCommitMarks(),                                    // Enable auto-commit for offsets
	)
	if err != nil {
		c.log.Fatal("Failed to connect to Redpanda: ", err)
	}
	defer client.Close()

	c.log.Info("Starting to poll for messages from topic 'emails'")
	for {
		c.log.Debug("Polling fetches...")
		fetches := client.PollFetches(context.Background())
		c.log.Debug("Fetch complete.")

		fetches.EachError(func(topic string, partition int32, err error) {
			c.log.Error("Failed to fetch messages from topic ", topic, " partition ", partition, ": ", err)
		})
		fetches.EachRecord(func(record *kgo.Record) {
			// Parse the record value as JSON
			var emailData struct {
				To      string `json:"to"`
				Subject string `json:"subject"`
				Message string `json:"message"`
			}

			err := json.Unmarshal(record.Value, &emailData)
			if err != nil {
				c.log.Error("Invalid message format: ", err)
				return
			}

			// Validate all required fields
			if emailData.To == "" || emailData.Subject == "" || emailData.Message == "" {
				c.log.Error("Missing required fields in message: ", string(record.Value))
				return
			}

			c.log.Info("Received message: ", string(record.Value))

			// Automatically send the email
			err = c.sender.SendEmail(emailData.To, emailData.Subject, emailData.Message)
			if err != nil {
				c.log.Error("Failed to send email: ", err)
			} else {
				c.log.Info("Email sent successfully for message: ", string(record.Value))
				// Mark the record as processed
				client.MarkCommitRecords(record)
			}
		})

		// Commit offsets for all processed records
		if err := client.CommitUncommittedOffsets(context.Background()); err != nil {
			c.log.Error("Failed to commit offsets: ", err)
		}
	}
}
