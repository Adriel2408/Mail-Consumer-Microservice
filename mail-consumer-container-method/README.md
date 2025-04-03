# Mail Consumer

## Setup
1. Install Go, Docker, and Docker Compose.
2. Run `docker-compose up --build`.

## Changing for another service (Ignore - Doesn't Matter For Container Testing)
- Step 1: Change the Redpanda Broker in .env to its respective key for the cluster that will be used
- Step 2: Change SASL_Username and password to respective user information from the user in the Redpanda cluster in .env
- Step 3: Change the topic name in consumer.go line 28 to respective topic name that the data will be consumed from
- Step 4: Ensure the correct SMTP is used for respective service with correct credentials

## Test Functionality
- Step 1: Load Up Docker
- Step 2: Run "docker exec -it mail-consumer-redpanda-1 rpk topic create emails" in your terminal to create a topic in your cluster, whatever name you put for your topic which in this case is emails will have to match the name in line 28 on consumer.go
- Step 3: Run "docker exec -it mail-consumer-redpanda-1 rpk topic produce emails" in your terminal
- Step 4: Type in your message, ensure it is in json format and follows this format - {"to": "", "subject": "", "message": "" }
- Step 5: Run "docker exec -it mail-consumer-redpanda-1 rpk topic consume emails" to check if the event was consumed and deleted. If nothing appears under this line then that means it was properly consumed and deleted from the topic