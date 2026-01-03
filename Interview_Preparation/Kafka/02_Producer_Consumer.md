# Kafka Workflow (Producer → Cluster → Consumer)

# Producer sends messages

A producer application sends messages to a Kafka topic.

If a key is specified, Kafka determines the partition for the message.

# Kafka Cluster receives the message

The cluster consists of multiple brokers.

Each partition has a leader broker and follower replicas.

The leader handles the write request from the producer.

Followers replicate the data from the leader for durability.

# Consumer Group reads messages

Consumers subscribe to a topic as part of a consumer group.

Each partition is read by only one consumer in the group.

Consumers track offsets to know which messages they’ve processed.

# Key Points in the Workflow

Producers → Partition leader → Followers replicate → Consumers

Fault tolerance: If a leader fails, a follower in ISR becomes the new leader.

Ordering: Kafka guarantees message order within a partition, not across partitions.

Decoupling: Producers and consumers work independently; slow consumers do not block producers.