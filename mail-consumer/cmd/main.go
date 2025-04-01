package main

import (
	"mail-consumer/internal/broker"
	"mail-consumer/internal/config"
	"mail-consumer/internal/logger"
	"mail-consumer/internal/mail"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	log := logger.NewLogger(cfg.LogDir)

	// Log Redpanda broker for debugging
	log.Info("Redpanda Broker: ", cfg.RedpandaBroker)

	sender := mail.NewSender(cfg.MailtrapHost, cfg.MailtrapPort, cfg.MailtrapAPIKey, log)
	consumer := broker.NewConsumer(cfg.RedpandaBroker, log, sender)

	// Start consuming messages and automatically send emails
	consumer.Consume()
}
