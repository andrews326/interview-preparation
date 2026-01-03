💠 Cassandra Interview Questions (Part 2: Advanced Topics)
1. Explain the architecture of Cassandra.

Cassandra uses a peer-to-peer distributed architecture — no master or slave nodes.
Each node in the cluster is equal and communicates via Gossip Protocol.

Data distribution: Based on consistent hashing (token ring).

Each node is responsible for a range of tokens (data).

Replication ensures redundancy and fault tolerance.

2. What is a Token and Token Ring?

Each row is assigned a token based on its partition key (using a hash function).

The token ring is a circular structure that maps tokens to nodes.

This ensures even distribution of data across the cluster.

3. What is Virtual Node (vNode)?

Instead of each node managing one big token range, it manages multiple small ranges (vNodes).
This improves:

Load balancing

Faster node recovery and data rebalancing

4. Explain Cassandra’s replication mechanism.

Cassandra replicates data across nodes based on:

Replication factor (RF) – number of copies per data item.

Replication strategy – determines where replicas are placed.

SimpleStrategy for single data centers.

NetworkTopologyStrategy for multiple data centers.

5. What is Consistency Level in detail?

Defines how many replicas must acknowledge a read/write before success.
Common levels:

ONE → Fast, but less consistent.

QUORUM → Balance between consistency and speed.

ALL → Strong consistency, but slower.

Formula for strong consistency:

R + W > RF


Where:

R = read consistency level

W = write consistency level

RF = replication factor

6. What happens when a node goes down?

Writes are stored as hints on other nodes (Hinted Handoff).

When the node recovers, the hints are replayed to it.

Read Repair and Anti-Entropy Repair keep replicas consistent.

7. What is Read Repair?

When a read detects outdated data among replicas, Cassandra automatically updates the stale replicas with the latest version in the background.

8. What is Anti-Entropy Repair (Full Repair)?

A manual or scheduled process that compares and syncs SSTables between replicas to fix inconsistencies across nodes.

9. What is the difference between Read Repair and Anti-Entropy Repair?
Feature	Read Repair	Anti-Entropy Repair
Trigger	On read request	Scheduled manually
Scope	Specific rows	Entire dataset
Performance	Lightweight	Heavy process
Use case	Fix small inconsistencies	Full sync across cluster
10. What is Hinted Handoff?

A mechanism to handle temporary node failures.
When a node is down, the coordinator stores a “hint” (the missed write). Once the node comes back, it replays these hints.

11. Explain the Write Path in detail.

Data received by Coordinator node.

Written to Commit Log (for durability).

Written to Memtable (in-memory).

When Memtable is full → Flushed to SSTable on disk.

Background processes like compaction merge SSTables.

12. Explain the Read Path in detail.

Coordinator node receives query.

Checks Memtable first.

Uses Bloom filter to check SSTables.

Reads data from SSTables, merges versions.

Applies read repair if needed.

13. What is Compaction and why is it important?

Combines multiple SSTables into a single one to:

Reduce disk space.

Improve read performance.

Delete tombstoned data.

14. What is the difference between Major and Minor Compaction?
Type	Description
Minor Compaction	Merges smaller SSTables automatically.
Major Compaction	Merges all SSTables into one — manual or scheduled.
15. What is a Coordinator Node?

The node that receives a read/write request from a client and coordinates the operation with other nodes (replicas).

16. What is Tunable Consistency?

Cassandra allows developers to tune the consistency level per query — balancing availability and consistency as needed.

17. What are Hinted Handoff and Read Repair part of?

They are part of Cassandra’s eventual consistency model — ensuring replicas become consistent over time, not immediately.

18. What is the CAP Theorem in Cassandra context?

Cassandra prefers Availability (A) and Partition Tolerance (P) over immediate Consistency (C) — though consistency can be tuned.

19. What is Data Compaction Strategy?

Defines how SSTables are merged. Common strategies:

SizeTieredCompactionStrategy (STCS) – default; merges similar-sized SSTables.

LeveledCompactionStrategy (LCS) – good for read-heavy workloads.

TimeWindowCompactionStrategy (TWCS) – good for time-series data.

20. How does Cassandra ensure fault tolerance?

Replication across multiple nodes.

No single point of failure (peer-to-peer).

Hinted handoff, repairs, replication ensure data recovery.

21. What is a Gossip Protocol used for?

Nodes exchange state information about themselves and others — used for cluster membership, failure detection, and metadata propagation.

22. How does Cassandra handle scaling?

Horizontal scaling: Just add new nodes.

Data automatically redistributed using vNodes.

No downtime required.

23. What is the difference between Cassandra and traditional RDBMS?
Feature	Cassandra	RDBMS
Schema	Flexible	Fixed
Joins	Not supported	Supported
Scalability	Horizontal	Vertical
Consistency	Tunable	Strong
Data Model	Wide-column	Relational
24. What is a Secondary Index disadvantage?

Slower on large datasets because Cassandra has to query multiple nodes — not efficient for high-cardinality columns.

25. How do you monitor Cassandra?

Using nodetool commands (e.g., nodetool status, nodetool repair).

Metrics via JMX, Prometheus, Grafana, or DataStax OpsCenter.