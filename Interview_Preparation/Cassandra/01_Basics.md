💠 Cassandra Interview Questions (Part 1: Basics to Intermediate)

1. What is Apache Cassandra?
A highly scalable, distributed NoSQL database designed to handle large amounts of data across many servers with no single point of failure.

2. What type of database is Cassandra?
It’s a NoSQL, column-family (wide-column) store, not relational.

3. Who originally developed Cassandra?
Developed at Facebook to power their Inbox Search feature; later open-sourced and maintained by the Apache Software Foundation.

4. What is a Cluster in Cassandra?
A cluster is a collection of nodes (servers) that hold data in a distributed manner.

5. What is a Node?
A node is a single machine (physical or virtual) in the cluster that stores part of the data.

6. What is a Keyspace?
A keyspace is the top-level namespace in Cassandra, similar to a database in RDBMS.
It defines replication strategies and replication factors.

7. What is a Table (Column Family)?
A table stores data in rows and columns, grouped by a primary key.
Each table belongs to a keyspace.

8. What is a Partition Key?
The partition key determines which node stores a particular row. It decides data distribution.

9. What is a Clustering Key?
Defines the order of data within a partition.

10. What is a Primary Key?
Combination of partition key + clustering key(s) — uniquely identifies each row.

11. What query language does Cassandra use?
CQL (Cassandra Query Language) — similar to SQL but without joins or foreign keys.

12. What is a Replica?
A copy of data stored on different nodes to ensure fault tolerance.

13. What is the Replication Factor (RF)?
The number of nodes that will store a copy of each piece of data.
Example: RF=3 means each data item is stored on 3 different nodes.

14. What are the Replication Strategies?

SimpleStrategy: For single data center setups.

NetworkTopologyStrategy: For multi–data center setups.

15. What is Consistency Level?
Defines how many replicas must acknowledge a read/write for it to be considered successful.
Examples: ONE, TWO, QUORUM, ALL.

16. What is a Write Path in Cassandra?

 1. The write request is sent to the coordinator node, which decides which node(s) will store the data based on the replication factor (RF) and the partition key.

 2. The data is then stored in the memtable of the responsible node(s), and simultaneously, it is written to the commit log to ensure data durability.

 3. The data is replicated to other nodes based on the replication factor (RF), ensuring multiple copies of the data are stored across different nodes.

 4. When the memtable becomes full, its data is flushed to disk in the form of an SSTable.

 5. If any replica is down, the coordinator creates a hint. Once the downed node is available again, the hint is used to store the missed data on that node.

17. What is a Read Path in Cassandra?

Query checks Memtable first.

Then Bloom Filter + SSTables for data.

Data is merged and returned to the client.

Cached if necessary.

1. Visualizing the Read Path (High-level):

2. Client Request → Coordinator Node

3. Coordinator finds relevant replica nodes based on partition key

4. Coordinator sends request to replicas based on consistency level

5. Replicas check memtables and SSTables for the data

6. Bloom filters and SSTable indexes help quickly locate data

7. If data is inconsistent across replicas, read repair is triggered

8. Data is merged from multiple sources (memtable, SSTables)

9. Coordinator sends the final result back to the client

18. What is a Commit Log?
A durable log where all write operations are recorded before being written to Memtable — ensures recovery if node crashes.

19. What is a Memtable?
An in-memory data structure that stores recently written data; flushed to disk as an SSTable when full.

20. What is an SSTable?
Sorted String Table — immutable file on disk that stores flushed data from Memtables.

21. What is Compaction?
The process of merging multiple SSTables into one to reduce disk usage and improve read performance.

22. What is a Bloom Filter?
A memory-efficient data structure used to check whether a key might exist in an SSTable, reducing disk lookups.

23. What is a Gossip Protocol?
A peer-to-peer communication mechanism that Cassandra nodes use to share information about themselves and other nodes (like heartbeat, status, etc.).

24. What is a Hint in Cassandra?
If a node is down, a hint is stored temporarily by another node, so it can deliver missed writes once the node comes back up (Hinted Handoff).

25. What is a Tombstone?
A marker placed on deleted data to ensure consistency during replication and compaction.

26. What is Lightweight Transaction (LWT)?
A way to achieve compare-and-set semantics in Cassandra using the IF clause.
Example: INSERT INTO users (id, name) VALUES (1, 'Andre') IF NOT EXISTS;

27. What is Data Modeling in Cassandra?
It’s query-based design — you model tables based on queries you need, not normalized relationships.

28. What is Denormalization in Cassandra?
Data is duplicated across tables to optimize for fast reads and avoid joins.

29. What are Secondary Indexes?
Indexes created on non-primary key columns to allow querying by those columns (but not recommended for large datasets).

30. What are Materialized Views?
Predefined query results stored as separate tables that auto-sync with the base table (used for alternate query patterns).