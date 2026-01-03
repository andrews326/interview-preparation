

⚙️ Redis Interview Questions and Answers – Intermediate Level (Q21–45)

21. What are RDB and AOF in Redis persistence?

RDB (Redis Database File): Creates point-in-time snapshots of your dataset at specific intervals.

AOF (Append Only File): Logs every write operation received by the server for full recovery.

22. What is the difference between RDB and AOF?

| Feature       | RDB                                         | AOF                             |
| ------------- | ------------------------------------------- | ------------------------------- |
| Speed         | Faster snapshot                             | Slower (writes every operation) |
| Durability    | Less (data loss possible between snapshots) | More (replays all operations)   |
| File size     | Smaller                                     | Larger                          |
| Recovery time | Faster                                      | Slower                          |


23. Can Redis use both RDB and AOF at the same time?
✅ Yes, Redis can use both. RDB ensures fast restarts, and AOF ensures minimal data loss.

24. What happens if Redis crashes?
If Redis crashes:

If RDB is enabled → restores last snapshot.

If AOF is enabled → replays operations from the log to restore data.

25. How can you recover Redis data after a crash?
Copy your latest RDB or AOF file into the Redis data directory and restart Redis. It will load the data automatically.

26. What is a Redis pipeline?
A pipeline allows sending multiple commands to Redis in a single network round trip — improving performance by reducing latency.

Example:

MULTI
SET user:1 "Andre"
SET user:2 "Amose"
EXEC


27. What is the difference between pipelining and transactions in Redis?

| Feature   | Replication    | Clustering                   |
| --------- | -------------- | ---------------------------- |
| Data copy | Entire dataset | Partitioned dataset          |
| Purpose   | Redundancy     | Scalability                  |
| Node type | Master/Slave   | Master/Replica (with shards) |


28. What is a Redis transaction?
A transaction is a sequence of commands executed as a single atomic operation using MULTI, EXEC, and DISCARD.

29. How do you use MULTI, EXEC, and DISCARD commands?

MULTI
SET name "Andre"
INCR count
EXEC


If you change your mind:

DISCARD


30. What is optimistic locking in Redis?
Redis supports optimistic locking using the WATCH command — it monitors keys for changes before executing a transaction.

31. What are Redis Pub/Sub channels?
Redis Publish/Subscribe allows real-time message broadcasting between clients.

Example:

SUBSCRIBE news
PUBLISH news "Breaking: Redis update released!"


32. How does Redis Pub/Sub mechanism work?
Publishers send messages to channels; subscribers listening to those channels receive them instantly — no data persistence.

33. What is Redis Cluster?
Redis Cluster is a distributed setup that splits data across multiple nodes using sharding and provides high availability.

34. What is sharding in Redis?
Sharding means dividing data across multiple Redis nodes so each node stores a portion of the dataset, improving scalability.

35. How does Redis Cluster ensure fault tolerance?
Each master node has one or more replica nodes. If a master fails, a replica is promoted to master automatically.

36. What is master-slave replication in Redis?
Redis supports asynchronous replication — the master node handles writes and replicates data to one or more slave nodes.

37. What is the difference between replication and clustering?

| Feature   | Replication    | Clustering                   |
| --------- | -------------- | ---------------------------- |
| Data copy | Entire dataset | Partitioned dataset          |
| Purpose   | Redundancy     | Scalability                  |
| Node type | Master/Slave   | Master/Replica (with shards) |


38. What are sentinel nodes in Redis?
Redis Sentinel is a monitoring and failover system for Redis replication setups. It monitors masters and automatically promotes a replica if the master fails.

39. What is the role of Redis Sentinel?

Monitors master/slave nodes

Performs automatic failover

Notifies admins and clients of changes

40. How do you monitor a Redis instance?
Use:

redis-cli monitor


or commands like:

INFO
SLOWLOG GET


41. How do you handle memory management in Redis?
Redis allows you to set memory limits and policies via maxmemory and maxmemory-policy configurations.

42. What is eviction policy in Redis?
When Redis runs out of memory, it evicts keys based on the configured policy.

43. What are the types of eviction policies available?

noeviction (default)

allkeys-lru

volatile-lru

allkeys-random

volatile-random

volatile-ttl

44. How can you secure Redis?

Use requirepass for password auth

Bind Redis to 127.0.0.1

Use firewalls or TLS

Avoid running as root

45. How do you take a backup of Redis data?
Copy the dump.rdb or appendonly.aof file from the Redis data directory.
Example:

cp /var/lib/redis/dump.rdb /backup/