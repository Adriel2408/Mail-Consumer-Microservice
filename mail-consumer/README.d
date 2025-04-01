# Mail Consumer

## Setup
1. Install Go, Docker, and Docker Compose.
2. Run `docker-compose up --build`.

## Changing for another service
- Step 1: Change the Redpanda Broker in .env to its respective key for the cluster that will be used
- Step 2: Change SASL_Username and password to respective user information from the user in the Redpanda cluster in .env
- Step 3: Change the topic name in consumer.go line 38 to respective topic name that the data will be consumed from
- Step 4: Ensure the correct SMTP is used for respective service with correct credentials
