📬 AWS SQS — Interview Questions & Answers
🌱 Basic Level

1. What is AWS SQS?
SQS (Simple Queue Service) is a fully managed message queuing service that allows decoupling of components of a distributed application. Producers send messages to a queue, and consumers retrieve them asynchronously.

2. What are the types of SQS queues?

Standard Queue: Unlimited throughput, at-least-once delivery, and messages may arrive out of order.

FIFO Queue: First-In-First-Out delivery with exactly-once processing; ensures order is preserved.

3. What is a message in SQS?
A message is the unit of data sent between producer and consumer, up to 256 KB in size.

4. How long can messages stay in a queue?
SQS supports a retention period of 1 minute to 14 days.

5. How do you access SQS?
Via AWS Console, CLI, SDKs, or REST API.

⚙️ Intermediate Level

1. How do you send a message to SQS?
Use SendMessage API, CLI command (aws sqs send-message) or SDKs. Include the message body and optional attributes.

2. How do you receive and delete messages?
Use ReceiveMessage to fetch messages and DeleteMessage after processing to remove them from the queue.

3. What is a visibility timeout?
A period during which a received message is hidden from other consumers. Prevents multiple consumers from processing the same message simultaneously.

4. How do you handle message retries?
If a consumer fails to process a message, it becomes visible again after the visibility timeout. You can also configure Dead Letter Queues (DLQs) for failed messages.

5. How do you secure SQS?
Use IAM policies to control access and optionally enable server-side encryption (SSE-KMS) to protect message data.

6. How do you scale SQS consumers?
Multiple consumers can process messages concurrently; for FIFO queues, concurrency is limited to maintain order.

7. How can SQS integrate with Lambda?
Lambda functions can be triggered automatically by new messages in SQS (event source mapping).

8. How do you ensure message ordering in SQS?
Use a FIFO queue and message group IDs to preserve order within each group.

9. How do you monitor SQS?
Use CloudWatch metrics like ApproximateNumberOfMessages, NumberOfMessagesSent, NumberOfMessagesReceived, NumberOfMessagesDeleted, and ApproximateAgeOfOldestMessage.

10. What is the difference between SQS and SNS?

SQS: Pull-based, decouples producers and consumers, messages are stored until processed.

SNS: Push-based, fan-out notifications to multiple subscribers in real-time.

🚀 Advanced Level

1. How do you implement a reliable decoupled architecture using SQS?
Use SQS queues between services to buffer requests, prevent service overload, and handle retries with DLQs for failed messages.

2. How do you integrate SQS with ECS, EKS, or Lambda in CI/CD pipelines?
Applications running in containers or serverless functions consume messages from SQS for processing asynchronously, enabling event-driven deployments.

3. How do you optimize SQS cost and throughput?

Use batching with SendMessageBatch and ReceiveMessage

Use short polling for low traffic and long polling for higher efficiency

Combine with FIFO queues only where ordering is required

4. How do you secure cross-account SQS access?
Set queue policies to allow specific AWS accounts to send or receive messages without sharing credentials.

5. How do you handle delayed messages?
Set the DelaySeconds parameter for individual messages or for the queue to postpone delivery.

6. How do you implement DLQs effectively?
Configure DLQ for messages that fail to process after a maximum number of retries to allow for later inspection and processing.

7. How do you ensure message duplication does not occur?
Use FIFO queues with deduplication IDs to ensure exactly-once message processing.

8. How do you integrate SQS with monitoring and alerts?
Use CloudWatch alarms on queue depth, age of oldest messages, and failure metrics to trigger alerts via SNS or Lambda.

9. How do you migrate messages between queues or accounts?
Use scripts, Lambda functions, or AWS Data Pipeline to read messages from one queue and send them to another queue.

10. What are best practices for SQS in production?

Use FIFO queues only if ordering is critical

Enable encryption for sensitive messages

Monitor metrics actively

Use DLQs for failed message handling

Combine with SNS for fan-out patterns if needed