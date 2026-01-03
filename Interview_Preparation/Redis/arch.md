🧠 Redis Architecture Overview

Redis follows a client–server architecture built for speed, simplicity, and scalability.

🧩 Core Components

Redis Server

The main daemon process (redis-server) that stores all data in memory.

Handles commands from clients and manages persistence, replication, and eviction.

Redis Client

Any application or service that connects to Redis (via TCP).

Uses the RESP (Redis Serialization Protocol) to communicate with the server.

Example: redis-cli, Python/Go/Java clients.

In-Memory Data Storage

Redis keeps all data in RAM.

Uses efficient encoding for strings, lists, sets, etc.

Optional disk persistence (RDB/AOF) ensures durability.

Persistence Layer

RDB (snapshot) and AOF (append-only file).

RDB gives quick recovery; AOF ensures minimal data loss.

Replication Layer

Provides high availability.

Master node asynchronously replicates data to replicas (slaves).

Slaves can serve read requests and be promoted if master fails.

Sentinel

Monitors Redis instances.

Performs automatic failover when the master node goes down.

Notifies clients and reconfigures connections.

Cluster

Redis Cluster splits the dataset into 16,384 hash slots.

Each master node handles a portion of slots.

Replica nodes mirror masters for redundancy.

Enables horizontal scaling and fault tolerance.

⚙️ Redis Workflow

Let’s walk through the request–response lifecycle:

1️⃣ Client Request

Client sends a command (e.g., SET user:1 "Andre") to Redis server via TCP (default port 6379).

2️⃣ Command Parsing

Redis uses RESP to parse the request.

It validates syntax, identifies the key, and determines which data structure is used.

3️⃣ In-Memory Execution

Redis performs the operation directly in RAM.

Example:

If SET → stores key/value in memory.

If INCR → increments numeric value.

If LPUSH → adds element to list head.

4️⃣ Persistence (Optional)

If persistence is enabled:

RDB: Snapshots the dataset periodically.

AOF: Logs the write command to the append-only file.

This ensures data recovery after restart.

5️⃣ Replication (If Configured)

Master sends the write to replica nodes asynchronously.

Replicas apply the same command to stay in sync.

6️⃣ Response to Client

Redis sends back the result — success (OK), value, or error message.

🗺️ Redis Cluster Architecture

A Redis Cluster typically includes:

+---------------------------+
|         Redis CLI         |
+------------+--------------+
             |
             v
+------------+--------------+
|        Redis Cluster       |
+------------+--------------+
| Master 1 | Master 2 | Master 3 |
|  + Rep1  |  + Rep2  |  + Rep3  |
+-----------+-----------+-----------+
   |          |           |
   v          v           v
 Clients    Clients     Clients

🔹 How it works

Each master manages a subset of hash slots (e.g., 0–5000, 5001–10000, etc.).

Clients can connect to any node.

If the key doesn’t belong to that node, it redirects using MOVED command.

Redis Cluster supports:

Auto-sharding

Auto-failover

Replication per master

🧩 Redis Sentinel Architecture
        +-------------+
        |  Sentinel 1 |
        +-------------+
              |
        +-------------+
        |  Sentinel 2 |
        +-------------+
              |
        +-------------+
        |  Sentinel 3 |
        +-------------+
              |
        +-------------+
        |  Master     |
        +-------------+
              |
      +---------------+
      | Replicas (2+) |
      +---------------+


Sentinels constantly ping master and replicas.

If the master fails → one replica is promoted.

Other sentinels agree via quorum (e.g., 2/3 must agree).

⚡ Redis Memory Management

Redis memory usage depends on:

Key-value data.

Overhead (data structures, metadata).

Expired keys are deleted using:

Lazy deletion: when accessed.

Active deletion: periodic scan.

If memory limit (maxmemory) is reached, Redis uses eviction policies:

allkeys-lru, volatile-lru, etc.

🔒 Redis Security

Disable external access (bind 127.0.0.1).

Enable password auth (requirepass).

Use TLS for encryption.

Avoid running as root.

🧠 Summary
| Component   | Description                           |
| ----------- | ------------------------------------- |
| Server      | Core engine handling data operations  |
| Client      | Applications communicating with Redis |
| Persistence | RDB & AOF ensure durability           |
| Replication | Master-slave redundancy               |
| Sentinel    | Automatic failover & monitoring       |
| Cluster     | Sharding + high availability          |
| Memory Mgmt | In-memory storage & eviction policies |
