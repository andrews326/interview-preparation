💠 Cassandra Architecture – Complete Explanation (A to Z)

Cassandra is a peer-to-peer, distributed NoSQL database designed for high availability, scalability, and fault tolerance — with no single point of failure.

Let’s break it down step by step 👇

🧩 1. Core Design Philosophy

Decentralized (peer-to-peer): Every node in Cassandra is equal; there’s no master or slave.

Distributed: Data is spread across multiple nodes.

Fault-tolerant: Even if nodes fail, data remains accessible.

Scalable: Add or remove nodes anytime — no downtime.

Tunable consistency: You can control how many replicas must respond for a read/write to succeed.

🧠 2. Main Components
Component                	Description
Node	                 A single Cassandra instance; stores data and handles read/write requests.
Cluster	                 A collection of nodes that work together.
Keyspace        The outermost namespace, like a “database” in RDBMS; defines replication strategy.
Table        (Column Family)	Stores rows of data, grouped by a primary key.
Row          	A set of columns identified by a unique primary key.
Partition      	Group of rows sharing the same partition key; defines which node stores them.

🌐 3. Token Ring & Data Distribution

Cassandra uses consistent hashing to distribute data.

Each node is assigned a token range (e.g., 0–100).

The partition key is hashed to determine which node stores that data.

Together, all nodes form a ring — called the Token Ring.

🌀 With Virtual Nodes (vNodes):
Each node manages multiple small token ranges → improves data balance and makes node recovery faster.

🧩 4. Replication

Cassandra stores multiple copies (replicas) of each piece of data for durability.

Replication Factor (RF): Number of copies.

Example: RF=3 means 3 nodes will store the same data.

Replication Strategies:

SimpleStrategy → single data center.

NetworkTopologyStrategy → multiple data centers (recommended).

⚙️ 5. Gossip Protocol

Nodes exchange information (like heartbeat, status, schema changes) using the Gossip protocol.

Gossip ensures all nodes know about each other and detect failures automatically.

🧭 6. Coordinator Node

When a client sends a request:

The node that receives it becomes the Coordinator Node.

It finds which nodes own the data (based on token ring).

Then coordinates read/write with those nodes.

💾 7. Write Path (How Data is Written)

Client sends write request to any node (Coordinator).

Data is first written to the Commit Log (for durability).

Then written to Memtable (in-memory cache).

Once Memtable fills up → flushed to disk as SSTable (Sorted String Table).

Later, compaction merges SSTables and removes tombstones (deleted data markers).

🧠 Durability rule:
Even if Cassandra crashes, data is safe because it’s in the Commit Log.

🔍 8. Read Path (How Data is Read)

Client sends read request → goes to Coordinator.

Coordinator contacts replica nodes.

Each replica checks:

Memtable (most recent data)

Bloom Filter + SSTables (on disk)

The newest version is merged and returned.

Optional: Read Repair updates outdated replicas.

🔄 9. Consistency & Availability

Cassandra provides Tunable Consistency:

You can choose between strong and eventual consistency.

Example:

ONE → fast, but may return stale data.

QUORUM → balance of speed & accuracy.

ALL → slow but always consistent.

✅ Formula for strong consistency:
R + W > RF

🧰 10. Background Processes
Process	Purpose
Compaction	Merges SSTables and removes old/tombstoned data.
Read Repair	Fixes inconsistencies found during reads.
Hinted Handoff	Saves missed writes for temporarily down nodes.
Anti-Entropy Repair	Periodic full sync between replicas.
💪 11. Fault Tolerance

If a node fails:

Writes are saved as hints → replayed when node returns.

Reads still succeed (if other replicas are alive).

No master node → cluster keeps running even with multiple failures.

⚡ 12. Scaling Cassandra

Horizontal scaling — just add more nodes.

Data automatically gets redistributed via vNodes.

No downtime or manual sharding needed.

📊 13. Cassandra in CAP Theorem

Cassandra = AP (Available + Partition-tolerant)

You can trade off consistency using tunable consistency levels.

🧩 14. Example Cluster Setup
   +-----------------------------+
   |         Cassandra Cluster   |
   +-----------------------------+
     |       |         |        |
   Node1   Node2     Node3    Node4
   | | |   | | |     | | |    | | |
   | | |   | | |     | | |    | | |
  SSTables CommitLog MemTable Gossip


All nodes are equal, each holds part of the data, and replicas ensure redundancy.

🧠 15. Key Advantages

Linearly scalable — add nodes, get more capacity.

High availability (no single failure point).

Tunable consistency.

Efficient for time-series, IoT, logs, and large-scale systems.

✅ Summary

Flow to remember easily:

Client → Coordinator → CommitLog → Memtable → SSTable
       ↕
Gossip + Replication + Read Repair + Compaction

