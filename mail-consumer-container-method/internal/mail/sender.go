package mail

import (
	"net/smtp"

	"github.com/sirupsen/logrus"
)

type Sender struct {
	host   string
	port   string
	apiKey string
	log    *logrus.Logger
}

func NewSender(host, port, apiKey string, log *logrus.Logger) *Sender {
	log.Info("Initializing SMTP sender with host: ", host, " and port: ", port)
	return &Sender{host: host, port: port, apiKey: apiKey, log: log}
}

func (s *Sender) SendEmail(to string, subject, message string) error {
	addr := s.host + ":" + s.port
	toList := []string{to}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")

	s.log.Info("Sending email to ", to, " via ", s.host, ":", s.port)

	client, err := smtp.Dial(addr)
	if err != nil {
		s.log.Error("SMTP connection error: ", err)
		return err
	}
	defer client.Close()

	// Disable STARTTLS for Mailhog
	if err := client.Hello("localhost"); err != nil {
		s.log.Error("SMTP Hello error: ", err)
		return err
	}

	if err := client.Mail("from@example.com"); err != nil {
		s.log.Error("SMTP Mail error: ", err)
		return err
	}

	for _, recipient := range toList {
		if err := client.Rcpt(recipient); err != nil {
			s.log.Error("SMTP Rcpt error: ", err)
			return err
		}
	}

	wc, err := client.Data()
	if err != nil {
		s.log.Error("SMTP Data error: ", err)
		return err
	}
	_, err = wc.Write(msg)
	if err != nil {
		s.log.Error("SMTP Write error: ", err)
		return err
	}
	err = wc.Close()
	if err != nil {
		s.log.Error("SMTP Close error: ", err)
		return err
	}

	s.log.Info("Email sent successfully to ", to)
	return nil
}
