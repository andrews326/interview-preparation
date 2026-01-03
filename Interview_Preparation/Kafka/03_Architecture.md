# Kafka Architecture – One-Liner Cheat Sheet

Producer: Sends messages to Kafka topics.

Topic: Named stream of messages; logical grouping.

Partition: Subdivision of a topic; maintains message order.

Broker: Kafka server that stores data and serves clients.

Consumer: Reads messages from topics.

Consumer Group: Set of consumers sharing topic partitions for load balancing.

Leader (Partition): Handles all reads/writes for that partition.

Follower (Partition): Replicates leader data for fault tolerance.

Replication Factor: Number of copies of a partition across brokers.

ISR (In-Sync Replica): Followers fully caught up with leader; eligible for leadership.

ZooKeeper/KRaft: Manages cluster metadata, leader election, and broker membership.

Offset is like a "bookmark" or "position marker" that tells Kafka where the consumer is in the stream of messages. It marks the position of the last message the consumer has read from a particular partition.

Commit is when the consumer officially marks the message as processed by saving the offset. It’s like saying, "I’ve read and successfully processed this message, and now I’m at this position (offset)."