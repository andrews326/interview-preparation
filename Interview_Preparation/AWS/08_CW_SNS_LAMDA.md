☁️ AWS CloudWatch — Interview Questions & Answers
🌱 Basic Level

1. What is CloudWatch?
CloudWatch is a monitoring service that provides metrics, logs, and alarms for AWS resources and applications.

2. What are CloudWatch metrics?
Metrics are quantitative data about AWS resources, like CPU usage, memory utilization, and network traffic.

3. What are CloudWatch logs?
Logs capture events and activities from AWS services, applications, or custom sources.

4. What is a CloudWatch alarm?
An alarm watches a metric and triggers an action (like sending a notification) when a threshold is breached.

5. How do you access CloudWatch?
Through AWS Management Console, AWS CLI, SDKs, or CloudWatch API.

⚙️ Intermediate Level

1. How do you monitor EC2 instances using CloudWatch?
Enable detailed monitoring or use CloudWatch Agent to track CPU, memory, disk, and network metrics.

2. How can CloudWatch logs be analyzed?
Use CloudWatch Logs Insights to query and visualize logs.

3. How do you automate responses to alarms?
Attach actions like sending an SNS notification, triggering a Lambda function, or auto-scaling EC2 instances.

4. What are custom metrics in CloudWatch?
Metrics sent by applications or scripts using CloudWatch API to monitor non-AWS resources or custom events.

5. How do you create dashboards in CloudWatch?
Use the CloudWatch console or API to visualize metrics, logs, and alarms in widgets on a single dashboard.

🚀 Advanced Level

1. How do you integrate CloudWatch with CI/CD pipelines?
Monitor deployment success, error rates, or resource usage and trigger automated notifications or rollback actions via Lambda.

2. How do you optimize CloudWatch cost?
Use metrics aggregation, log retention policies, and filter patterns to avoid storing unnecessary data.

3. How do you monitor Kubernetes clusters using CloudWatch?
Install CloudWatch Agent on nodes or use Container Insights to capture container and pod metrics.

4. How do you troubleshoot applications with CloudWatch logs?
Query logs with CloudWatch Logs Insights, set alarms on error patterns, and integrate with SNS or Lambda for automatic alerts.

5. What are best practices for CloudWatch?
Use tagging, custom namespaces, structured logs, dashboards for monitoring, and alerts for critical resources.

🔔 AWS SNS — Interview Questions & Answers
🌱 Basic Level

1. What is SNS?
SNS (Simple Notification Service) is a messaging service that sends notifications to subscribers via email, SMS, HTTP/S, or Lambda.

2. What is a topic in SNS?
A topic is a logical access point where publishers send messages that are delivered to all subscribers.

3. Who can publish messages to an SNS topic?
Publishers can be applications, AWS services, or users with appropriate permissions.

4. Who can receive messages from an SNS topic?
Subscribers include email addresses, phone numbers, SQS queues, HTTP endpoints, or Lambda functions.

5. How is SNS different from SQS?
SNS = push-based notifications to multiple subscribers.
SQS = pull-based queue for decoupling messages and processing asynchronously.

⚙️ Intermediate Level

1. How do you secure an SNS topic?
Use IAM policies, topic policies, and encryption (SSE-KMS) to control access and protect messages.

2. How do you send notifications from CloudWatch alarms?
Attach an SNS topic to the alarm; when the alarm triggers, the message is sent to all subscribers.

3. How do you integrate SNS with Lambda?
Subscribe a Lambda function to the topic; SNS triggers the Lambda function automatically when a message is published.

4. How do you handle message retries in SNS?
SNS automatically retries delivery to HTTP/S endpoints for up to several hours. For SQS endpoints, messages remain until processed or expire.

5. What are the use cases of SNS?
Monitoring alerts, CI/CD notifications, event-driven workflows, fan-out message distribution to multiple subscribers.

🚀 Advanced Level

1. How do you implement a fan-out architecture using SNS and SQS?
Publish messages to an SNS topic, which then delivers them to multiple SQS queues for parallel processing.

2. How do you ensure message durability in SNS?
Enable message encryption (SSE-KMS) and use SQS as a subscriber to persist messages reliably.

3. How do you integrate SNS with CloudFormation or Terraform?
Define SNS topics, subscriptions, and policies as resources in IaC templates for automated infrastructure setup.

4. How do you secure cross-account SNS access?
Use topic policies to allow specific AWS accounts to publish or subscribe.

5. How is SNS used in CI/CD pipelines?
Send notifications on build status, deployment completion, or failure alerts to teams via email, Slack, or Lambda triggers.

⚡ AWS Lambda — Interview Questions & Answers
🌱 Basic Level

1. What is AWS Lambda?
Lambda is a serverless compute service that runs code in response to events without provisioning or managing servers.

2. How do you trigger Lambda functions?
Triggers can include S3 uploads, SNS messages, API Gateway requests, CloudWatch events, or DynamoDB streams.

3. How is Lambda billed?
Billing is based on the number of requests and the compute time (GB-seconds) your code consumes.

4. What is the maximum execution time of a Lambda function?
The maximum timeout is 15 minutes per invocation.

5. What programming languages does Lambda support?
Node.js, Python, Java, Go, Ruby, .NET Core, and custom runtimes.

⚙️ Intermediate Level

1. How do you handle environment variables in Lambda?
You can define key-value pairs in the Lambda configuration to pass runtime settings securely.

2. How do you deploy Lambda functions in CI/CD pipelines?
Use AWS CLI, SDKs, or tools like Jenkins/CodePipeline to package and deploy functions automatically.

3. How do you integrate Lambda with S3 and SNS?
Set S3 events or SNS topics as triggers so Lambda executes automatically when an object is uploaded or a message is published.

4. How do you monitor Lambda execution?
Use CloudWatch logs and metrics (invocations, errors, duration, throttles).

5. How do you handle versioning in Lambda?
Enable versions and aliases to manage multiple releases and support safe deployments or rollbacks.

🚀 Advanced Level

1. How do you implement CI/CD with Lambda and API Gateway?
Deploy code via pipelines (Jenkins, CodePipeline) using CloudFormation/Terraform templates, test with automated scripts, and version functions with aliases for production deployment.

2. How do you manage concurrency in Lambda?
Set reserved concurrency limits to control the number of simultaneous executions and prevent downstream overload.

3. How do you secure Lambda functions?
Use IAM roles, environment variable encryption, VPC integration, and KMS for sensitive data.

4. How do you handle errors and retries?
Use Dead Letter Queues (DLQ), built-in retries for asynchronous invocations, and CloudWatch alarms to monitor failures.

5. How do you optimize Lambda cost?
Use efficient code, right-size memory allocation, avoid unnecessary invocations, and batch processes where possible.