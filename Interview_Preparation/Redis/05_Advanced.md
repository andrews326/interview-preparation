🚀 Redis Interview Questions and Answers – Advanced Level

46. What are Redis Streams?
Streams are a data structure for storing ordered sequences of messages (log data). They enable consumer groups to process messages in a queue-like fashion.

47. What are HyperLogLogs in Redis?
HyperLogLogs are probabilistic data structures used to count unique items with minimal memory (~12 KB) — used for cardinality estimation.

48. What are Bitmaps and Bitfields?
They let you store and manipulate bits inside string values — useful for tracking flags (e.g., user login days in a month).

49. How does Redis handle large datasets?
By using sharding, compression, and offloading to disk with Redis on Flash (in Enterprise edition). Cluster mode scales horizontally.

50. What is Redis replication lag?
The delay between data written to the master and replicated to the slave — can be monitored with INFO replication.

51. How can you optimize Redis performance?

Use pipelining to reduce round trips.

Enable connection pooling.

Use appropriate eviction policy.

Tune maxmemory.

Avoid large keys and slow Lua scripts.

52. What are Redis modules?
Modules extend Redis functionality (e.g., RedisJSON, RedisSearch, RedisGraph). They can add custom commands and data types.

53. What is RedisJSON, RedisSearch, and RedisGraph?

RedisJSON: Stores and queries JSON data.

RedisSearch: Full-text search and secondary indexing.

RedisGraph: Graph database engine built on Redis.

54. What is the difference between Redis and Memcached?

| Feature     | Redis | Memcached    |
| ----------- | ----- | ------------ |
| Data types  | Many  | Strings only |
| Persistence | Yes   | No           |
| Replication | Yes   | No           |
| Scripting   | Yes   | No           |


55. How can you scale Redis horizontally?
By using Redis Cluster or proxy-based sharding (like Twemproxy or Codis).

56. How does Redis handle consistency in distributed mode?
Redis uses asynchronous replication, so it’s eventually consistent. Strong consistency is achieved with WAIT command.

57. What is a hash slot in Redis Cluster?
Redis Cluster divides the keyspace into 16,384 hash slots. Each key maps to a slot using CRC16 hashing, and each slot belongs to a node.

58. How does Redis Cluster detect node failure?
Nodes ping each other using a gossip protocol. If a node misses pings for a threshold time, it’s marked as “FAIL”.

59. How do you rebalance data in Redis Cluster?
Use the redis-cli --cluster rebalance command to redistribute hash slots evenly across nodes.

60. What happens when the master node fails in a Redis Cluster?
A replica is automatically promoted to master by the cluster election mechanism to maintain availability.

61. How does Redis ensure data durability?
Through AOF, RDB, and replication. You can also configure appendfsync everysec for 1 second durability.

62. What are Lua scripts in Redis?
Redis supports Lua scripting for executing multiple commands atomically on the server side without network overhead.

63. How do you execute a Lua script in Redis?

EVAL "return redis.call('GET', 'key')" 0


64. What is Redis latency, and how do you measure it?
Latency is the delay between sending a command and getting a response. Use redis-cli --latency or INFO latency.

65. What is the difference between Redis Enterprise and open-source Redis?
Redis Enterprise adds:

Active-Active geo replication

Redis on Flash (storage extension)

Enhanced security and monitoring

Automatic sharding and failover

66. How does Redis handle concurrency?
Redis is single-threaded, so commands execute sequentially — no race conditions within a single instance. Concurrency handled via multiple connections.

67. What are common Redis performance issues?

Blocking commands (e.g., KEYS, SAVE)

Large keys or values

Insufficient memory

Network latency

Unoptimized AOF settings

68. What monitoring tools can be used with Redis?

RedisInsight

Prometheus + Grafana

Datadog

ELK Stack

Redis-CLI INFO

69. How do you handle Redis key naming conventions?
Use namespaces with colons:
app:user:1001:name to organize keys and avoid collisions.

70. How do you troubleshoot Redis memory leaks?

Run INFO memory and MEMORY USAGE commands.

Check for keys that never expire.

Use redis-cli --bigkeys to find large keys.

Adjust eviction policy if needed.