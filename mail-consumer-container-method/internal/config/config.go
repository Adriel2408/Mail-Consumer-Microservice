package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedpandaBroker string
	MailtrapHost   string
	MailtrapPort   string
	MailtrapAPIKey string
	LogDir         string
}

func LoadConfig(envFile string) (Config, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return Config{}, err
	}

	return Config{
		RedpandaBroker: os.Getenv("REDPANDA_BROKER"),
		MailtrapHost:   os.Getenv("SMTP_HOST"), // Ensure this matches the .env key
		MailtrapPort:   os.Getenv("SMTP_PORT"), // Ensure this matches the .env key
		MailtrapAPIKey: os.Getenv("MAILTRAP_API_KEY"),
		LogDir:         os.Getenv("LOG_DIR"),
	}, nil
}
