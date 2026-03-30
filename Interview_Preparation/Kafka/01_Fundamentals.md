
# Kafka Cluster – What’s Inside

A Kafka cluster is made up of one or more brokers that work together to provide scalability, reliability, and fault tolerance. The cluster itself contains:

# Brokers

Each broker is a Kafka server storing partitions of topics.

Brokers handle producer writes and consumer reads.

# Topics & Partitions

Topics exist logically across the cluster.

Partitions are distributed across multiple brokers.

# Leaders and Followers

Each partition has one leader broker and zero or more follower replicas.

Leaders handle reads/writes; followers replicate data for durability.

# Replication & ISR

Replication factor defines how many copies of each partition exist.

In-Sync Replicas (ISR) are replicas fully caught up and eligible to become leader if needed.

# ZooKeeper / KRaft (Cluster Metadata Manager)

ZooKeeper (or KRaft in newer versions) manages:

Broker metadata (who is in the cluster)

Leader election for partitions

Cluster configuration

# What is NOT inside the cluster

Producers and Consumers are external clients.

Producers send data to the cluster.

Consumers read data from the cluster.

Consumer groups exist logically, not as part of the cluster itself.

Quick Visual
             Producer
                 |
                 v
      +---------------------+
      |      Kafka Cluster  |
      |                     |
      |  Broker 1           |
      |    P0 Leader        |
      |    P1 Follower      |
      |                     |
      |  Broker 2           |
      |    P0 Follower      |
      |    P1 Leader        |
      |                     |
      |  Broker 3           |
      |    P0 Follower      |
      |    P1 Follower      |
      +---------------------+
                 |
                 v
            Consumer Group


✅ Key takeaway:

Inside cluster: Brokers, partitions, leaders/followers, replication, ISR, ZooKeeper/KRaft.

Outside cluster: Producers and consumers (clients), consumer groups (logical).

Offset = position in the Kafka stream of messages.

Commit = action by the consumer to mark a message as processed and update Kafka with the last successfully processed offset



# How Kafka Distributes 6 Partitions on 3 Brokers
    
Let's look at your 6 Partitions and 3 Brokers example with a Replication Factor of 2 (meaning 1 Leader + 1 Follower for every partition).

1. The Leader and Follower Map
Kafka tries to spread the Leaders and Followers so that if one broker fails, you don't lose everything. Here is how Kafka might distribute them:


Partition,   Leader (Active),  Follower (Backup)
Partition 0,  Broker 1,         Broker 2
Partition 1,  Broker 2,         Broker 3
Partition 2,  Broker 3,         Broker 1
Partition 3,  Broker 1,         Broker 3
Partition 4,  Broker 2,         Broker 1
Partition 5,  Broker 3,         Broker 2