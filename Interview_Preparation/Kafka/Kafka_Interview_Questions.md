# Kafka Interview Questions

## 1. Fundamentals
1. What is Apache Kafka and why is it used? 

Answer: Kafka is a distributed event streaming platform used for building real-time data pipelines and streaming applications.
It allows high-throughput, low-latency data transfer between producers and consumers.

2. Difference between Kafka and traditional message queues?  

| Kafka                           | Traditional MQ            |
| ------------------------------- | ------------------------- |
| Stores messages for a long time | Deletes after consumption |
| Can re-read messages            | Usually can’t re-read     |
| High throughput                 | Lower throughput          |
| Scales horizontally             | Limited scalability       |

# high-throughput ?
Specifically, it refers to the ability of a system to handle a large number of operations in a given time.(speed of operation like api request and response between time)

3. Explain Kafka’s architecture (Producer, Broker, Consumer, Topic). 

Producer: Sends messages to Kafka topics.

Consumer: Reads messages from topics.

Broker: Kafka server storing messages.

Topic: Named stream of messages.

Partition: Subset of a topic for parallelism and scalability.

4. What is a Kafka topic?  

Answer: A topic is a category or feed name to which messages are sent. Consumers subscribe to topics to read messages.

5. What is a partition and why do we use it?  
 Answer: A partition is a subdivision of a topic.

Used to parallelize consumption.

Ensures scalability across multiple brokers.

Each partition has an ordered, immutable log.

#  Without Partition Key (Round-robin method):

If you don’t provide a partition key, Kafka will just distribute the messages evenly across the available partitions. Kafka doesn't care about the content of the messages. It just tries to balance the load.

6. What is an offset in Kafka?  

Answer: Offset is the position of a message in a partition. Consumers track offsets to know which messages they have read.


7. What is a broker?  

Answer: A broker is a Kafka server that stores messages and serves clients (producers/consumers). Multiple brokers form a Kafka cluster.

8. What is a cluster in Kafka?  

Answer: A cluster is a group of Kafka brokers working together to provide scalability and fault tolerance.

9. What is a Producer?  

Answer: A producer is a client application that sends messages to Kafka topics.

10. What is a Consumer?  

Answer: A Consumer is a client application that read messages from the  Kafka topics.

 Consumers can be part of a Consumer Group to share load.

11. What is a Consumer Group?  

Answer: A consumer group is a set of consumers that work together to read messages from a topic.

Each partition is consumed by only one consumer in the group, ensuring load balancing.

Multiple consumer groups can consume the same topic independently.


12. What is a leader and follower in Kafka?  

Each partition has one leader and zero or more followers.

Leader: Handles all read/write requests for the partition.

Followers: Replicate the leader’s data to ensure fault tolerance.

If the leader fails, a follower becomes the new leader.


13. How does Kafka achieve fault tolerance?  

Answer:

Kafka replicates partitions across multiple brokers.

If a broker fails, followers with the same data take over.

Uses In-Sync Replicas (ISR) to track healthy replicas.


14. What is replication factor?  

Answer: Number of copies of a partition across brokers.

Ensures data durability and availability if a broker fails.

Example: Replication factor = 3 → 1 leader + 2 followers.


15. What is retention period in Kafka?

the retention period is the amount of time that Kafka retains messages in a topic before they are deleted.

Answer: Time Kafka keeps messages in a topic.

Messages can be deleted after this period even if they aren’t consumed.

Helps manage disk space.

## 2. Producer & Consumer Mechanics
16. How does a producer decide which partition to send a message to?  

If a key is provided, Kafka hashes the key to choose a partition.

If no key, messages are sent in round-robin to partitions.

17. What are producer acknowledgments (acks=0, 1, all)?  

acks=0: Producer does not wait for acknowledgment. Fast but not reliable.

acks=1: Waits for leader to acknowledge. Moderate reliability.

acks=all (or -1): Waits for all in-sync replicas to acknowledge. Most relia

18. What is message batching and compression in Kafka?  

Answer:

Batching: Sends multiple messages together for efficiency.

Compression: Reduces network usage (gzip, snappy, lz4).

Helps increase throughput and reduce latency.


Batching refers to grouping multiple messages (or records) into a single request (or batch) that can be sent to Kafka brokers.

Compression is a technique used to reduce the size of messages before they are sent over the network or stored on disk.

19. What happens when a consumer fails in a consumer group? 

Kafka rebalances the partitions among remaining consumers.

Failed consumer’s partitions are reassigned to others to continue processing.

20. How does Kafka handle backpressure (slow consumers)?  

Kafka handles backpressure (slow consumers) by using buffers and flow control mechanisms to manage the rate at which data is read and written. Here's a simple breakdown:

Producer-side buffering:
Consumer-side buffering:
Offset tracking:

Kafka retains messages in the topic (up to retention limit).

Slow consumers can read at their own pace without blocking producers.

This decouples producers and consumers, unlike traditional queues.

21. How does Kafka ensure exactly-once delivery?  

Kafka uses idempotent producers and transactions.

Producers assign unique IDs to messages, so even if they retry, duplicates aren’t written.

Consumers use transactional reads to ensure each message is processed exactly once.

22. What are delivery semantics (at-most-once, at-least-once, exactly-once)?  

At-most-once: Message may be lost, never duplicated.

At-least-once: Message may be delivered multiple times, never lost.

Exactly-once: Message is delivered once, no loss or duplication.

23. How do you manually commit offsets?  

In consumer code, call commitSync() or commitAsync() to record the last read offset.

Helps resume consumption after failure without losing or duplicating messages.

24. What’s the difference between auto-commit and manual commit? 

| Auto-commit                                      | Manual commit                                         |
| ------------------------------------------------ | ----------------------------------------------------- |
| Kafka automatically commits offsets at intervals | Developer decides when to commit                      |
| Easier to use                                    | Provides more control & prevents duplicate processing |
| May commit before processing                     | Safer for exactly-once processing                     |


25. What is offset reset policy (earliest/latest)?  

Determines where a consumer starts reading if no offset is found.

earliest: Start from the beginning of the partition.

latest: Start from new messages only.


## 3. Kafka Architecture & Storage
26. How does Kafka store data on disk?  

Messages are stored in segment files on disk.

Each partition is a log; new messages are appended.

Efficient sequential writes improve throughput.

27. What is a segment file?  

A segment file is A chunk of messages in a partition.

Kafka rotates segment files based on size or time.

Helps with efficient storage management.

28. How does Kafka handle large volumes of data efficiently?  

Uses append-only logs (sequential writes).

Zero-copy mechanism reduces memory copying.

Data can be replicated and distributed across brokers.

29. What is zero-copy mechanism in Kafka?  

Allows Kafka to send data from disk to network without copying to user space.

Reduces CPU usage and increases throughput.

30. What is Kafka’s write-ahead log (WAL)?  

Each partition is a log of messages stored on disk before acknowledgment.

Ensures durability; even if broker crashes, committed messages are not lost.

31. Explain how replication works between brokers.  

Each partition has one leader and followers.

Followers replicate messages from the leader.

If the leader fails, one follower takes over as the new leader, ensuring fault tolerance.

32. What is ISR (In-Sync Replica)?  

ISR is the set of replicas fully caught up with the leader.

Only replicas in ISR can become leader during failure.

Ensures data consistency and durability.

33. What happens when a broker in ISR goes down?  

Leader detects failure.

A new leader is chosen from the remaining in-sync replicas.

Consumers continue reading without data loss.

34. How does leader election happen in Kafka?  

The controller broker detects broker failure.

Among ISR replicas, a new leader is elected automatically.

Leadership ensures no message loss and cluster stability.

35. What is KRaft mode (vs Zookeeper)?  

KRaft (Kafka Raft metadata mode) replaces Zookeeper.

Manages broker metadata and leader election internally.

Simplifies deployment and removes Zookeeper dependency.

## 4. Kafka Reliability & Performance
36. How does Kafka guarantee message durability? 

Replication ensures multiple copies of data.

Messages are written to disk before acknowledgment.

ISR ensures only fully replicated data is acknowledged.

37. How can you improve throughput in Kafka?  

Increase partition count.

Use batching and compression.

Tune producer and consumer configurations (linger.ms, buffer size).

Scale brokers horizontally.

38. What’s the impact of replication factor on performance?  

Higher replication = better fault tolerance, but slightly lower throughput due to extra replication.

Lower replication = faster, but risk of data loss.

39. What is message ordering in Kafka?  

Kafka guarantees message order within a single partition.

Messages across partitions are not ordered.

40. Can Kafka lose messages? How do you prevent it?  

Kafka can lose messages if replication factor = 1 or acks=0.

To prevent loss:

Set replication factor ≥ 2

Use acks=all

Ensure in-sync replicas (ISR) are maintained

41. What happens if a producer sends faster than Kafka can handle?  

Kafka uses internal buffers to store messages temporarily.

If buffers are full, the producer can:

    Block until space is available

    Throw an exception (if configured with max.block.ms)

42. What’s the typical latency in Kafka?  

Kafka usually has 1–10 ms latency for single messages.

Batching increases throughput but can slightly increase latency.

43. How does Kafka handle large messages?  

Large messages can be split into smaller chunks or increase message.max.bytes.

Producers and consumers must adjust configurations (fetch.max.bytes, max.request.size).

Avoid very large messages in production; prefer external storage and send references.

44. How can you monitor Kafka performance?  

Monitor consumer lag, throughput, disk usage, replication status.

Metrics can be collected via JMX and visualized in Prometheus + Grafana.

45. What tools are used to monitor Kafka?  

Kafka Manager / CMAK

Confluent Control Center

Prometheus + Grafana

Burrow for consumer lag

## 5. Advanced / Tricky Scenario Questions
46. What happens when all brokers fail?  

Kafka cannot serve producers or consumers.

Messages in ISR are safe, but unreplicated messages may be lost.

Once brokers recover, Kafka restores leadership from replicas.


47. Can two consumers in the same group read the same partition?  

No. Within a consumer group, each partition is assigned to only one consumer.

Multiple consumers in the same group read different partitions for load balancing.

48. What happens when replication factor = 1 and the broker fails?  

Partition data may be lost since there’s no replica.

Always use replication factor ≥ 2 for reliability.

49. What’s the difference between partition leader and replicas?  

Leader: Handles all reads/writes.

Replicas (followers): Copy leader’s data for fault tolerance.

Followers cannot serve clients unless they become the leader.

50. What if producer sends data but doesn’t receive an acknowledgment?  

Producer may retry sending depending on retries config.

Using idempotent producer ensures duplicates are not created.

If retries fail, message may be lost if acks=0; otherwise, delivery is guaranteed.

51. How to achieve high availability in Kafka?  

Use replication factor ≥ 2.

Ensure multiple brokers in the cluster.

Maintain In-Sync Replicas (ISR).

Deploy multiple Zookeeper/KRaft nodes for metadata management.

52. What’s the difference between Kafka Streams and Kafka Connect?  

Kafka Streams: A library for building stream processing applications.

Kafka Connect: A framework for moving data in/out of Kafka from databases, files, or other systems.

53. What are Kafka Connectors?  

Pre-built or custom plugins for Kafka Connect.

Examples: JDBC Source Connector, S3 Sink Connector.

Helps integrate Kafka with external systems easily.

54. Explain the role of Schema Registry.  

Stores Avro, JSON, or Protobuf schemas for messages.

Ensures schema compatibility between producers and consumers.

Helps prevent data format errors in pipelines.

55. What’s the difference between Kafka and Flume / Spark Streaming?  

| Kafka                           | Flume / Spark Streaming                        |
| ------------------------------- | ---------------------------------------------- |
| Distributed log platform        | Data ingestion/processing tools                |
| Stores messages for long time   | Processes messages in-memory or near real-time |
| Decouples producers & consumers | Focused on ETL or batch processing             |


56. Explain exactly-once semantics in Kafka.  

Guarantees each message is processed exactly once by consumer.

Achieved via idempotent producers + transactions in Kafka Streams.

57. How does Kafka handle rebalancing of consumer groups?  

When a consumer joins/leaves a group, Kafka reassigns partitions among consumers.

This ensures load balancing but may temporarily pause consumption.

58. How can message duplication occur in Kafka?  

Occurs if:

    Producer retries without idempotence

    Consumer commits offset after processing failure

Use idempotent producers and manual offset commit to prevent duplication.

59. What happens when a topic reaches its retention time?  

Kafka deletes messages older than the configured retention period.

Helps manage disk usage automatically.

60. What happens when a partition is full or disk runs out of space? 

Producer cannot write new messages → may throw exceptions.

Kafka stops accepting writes until disk space is available.

Use log compaction, retention policies, or increase disk capacity to avoid this.



## 6. Kafka in Real Projects (DevOps + Cloud)
61. How do you secure Kafka (authentication & authorization)?  

Authentication: Use SASL/SSL to verify producers and consumers.

Authorization: Use ACLs (Access Control Lists) to control who can read/write to topics.

Helps prevent unauthorized access in production environments.

62. What is SSL and SASL in Kafka?  

SSL (Secure Sockets Layer): Encrypts data over the network.

SASL (Simple Authentication and Security Layer): Provides authentication mechanism for users.

Combined, they ensure secure and authenticated communication.

63. How to deploy Kafka on Kubernetes / Docker?  

Use Docker images for Kafka brokers.

Deploy using Kubernetes StatefulSets for stable network IDs.

Use Persistent Volumes (PV) for data storage.

Configure headless service for broker discovery.

64. How to set up Kafka monitoring in AWS or GCP?  

Collect metrics using JMX exporters.

Use Prometheus + Grafana dashboards.

Monitor broker health, consumer lag, disk usage, and throughput.

65. How do you scale Kafka horizontally?  

Add more brokers to the cluster.

Increase partitions of topics to distribute load.

Ensure replication factor is adjusted for new brokers.

66. What happens when a new broker joins the cluster?  

Kafka rebalances partitions to include the new broker.

Reduces load on existing brokers.

Helps increase cluster capacity and throughput.

67. What metrics are critical to monitor in production Kafka?  

Consumer lag

Replication status & ISR

Throughput (messages/sec)

Disk usage

Request/response latency

Under-replicated partitions

68. How to troubleshoot consumer lag?

Check consumer offsets vs latest offsets.

Increase consumer parallelism or processing speed.

Optimize fetch size and batch processing.

Ensure no network bottlenecks.

69. How do you ensure schema compatibility between producers and consumers?  

Use Schema Registry to store message schemas.

Enforce backward/forward compatibility rules.

Prevents consumers from breaking when producers change message structure.

70. What are dead-letter queues (DLQs) in Kafka?  

DLQs store messages that cannot be processed successfully.

Helps debug or reprocess problematic messages without affecting main pipeline.


## 7. Coding / Real-world Scenario Questions
71. Write a Go/Python/Java producer to send messages to Kafka.  
72. How do you consume messages from Kafka in your language?  
73. What will you do if your consumer lags behind producer for hours?  

Check consumer performance (processing speed).

Increase consumer instances or partitions.

Optimize message processing logic.

If necessary, skip/skip/replay messages depending on SLA.

74. How do you ensure data consistency between microservices using Kafka?  

Use idempotent producers to avoid duplicates.

Implement exactly-once processing with transactions.

Use schema registry to maintain data format consistency.

75. How would you migrate data from one Kafka cluster to another?  

Use MirrorMaker or Kafka Connect.

Set up source and target clusters.

Mirror topics and offsets to the new cluster.

Test consumer behavior on the target cluster.

76. What happens when multiple producers write to the same partition?  

Kafka preserves order within a partition.

Messages are appended sequentially.

Use keys if specific ordering is required per key.

77. How would you ensure ordering for a specific key (e.g., user ID)?  

Always send messages with the same key.

Kafka ensures messages with the same key go to the same partition, preserving order.

78. How can Kafka be used for event-driven microservices?  

Microservices publish events to topics.

Other services subscribe to relevant topics.

Decouples services and enables asynchronous communication.

79. How to handle poison messages (bad data)?  

Detect and redirect to a Dead Letter Queue (DLQ).

Log the error and alert operators.

Optionally reprocess after fixing the data.


80. How to replay messages in Kafka for debugging?

Consumers can reset offsets (earliest or specific offset).

Read messages again from the topic.

Useful for reprocessing events or debugging failed processing.
