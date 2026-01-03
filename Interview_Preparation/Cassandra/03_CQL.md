✅ Must-Know (Basic → Intermediate) Topics — Enough for Interview
1. Core Concepts

What is Cassandra?
→ NoSQL, distributed, column-family-based database

Difference between Cassandra and RDBMS

CAP theorem — Cassandra is AP (Available + Partition tolerant)

Key features: scalability, fault tolerance, no single point of failure

2. Data Model

Keyspace, Table (Column Family), Partition Key, Clustering Key

Primary key structure (Partition + Clustering keys)

CQL (Cassandra Query Language) basics

How data is distributed using consistent hashing

3. Architecture

Node, Cluster, Data Center

Gossip protocol

Snitch and Replication Strategy

Write path: Commit log → Memtable → SSTable

Read path: Bloom filter → Key cache → Row cache → SSTable

4. Replication & Consistency

Replication factor

Consistency levels (ONE, QUORUM, ALL)

Tunable consistency concept

Difference between eventual consistency and strong consistency

5. Compaction & Repair

What is compaction and why it’s needed

What is hinted handoff and read repair

6. Performance / Admin

How Cassandra achieves high availability

Why writes are fast in Cassandra

What happens when a node fails

Simple backup and restore strategies