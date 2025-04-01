# Mail Consumer Specification

## Purpose
A microservice to consume email events from Redpanda and send them via SMTP.

## Technologies
- Event Broker: Redpanda
- Language: Golang
- Mail Testing: Mailhog
- Virtualization: Docker
- Data Format: JSON

## Requirements
- Consume events from "emails" topic.
- Send emails with retry on failure (TBD).
- Log errors in JSON format to `logs/log_YYYYMMDD.log`.
- Load config from `.env`.