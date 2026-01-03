🧠 Redis Interview Questions and Answers – Basic Level

1. What is Redis?
Redis stands for REmote DIctionary Server and is an open-source, in-memory data structure store. It is primarily used as a database, cache, and message broker. It stores data in RAM for high-speed read/write operations.

2. What are the main features of Redis?

Extremely low-latency performance (In-memory storage (very fast))

Supports complex data types

Persistence options (RDB, AOF)

Pub/Sub messaging

High availability with replication & clustering

Lua scripting

Atomic operations

3. What data structures are supported by Redis?
Redis supports:

Strings

Lists

Sets

Sorted Sets (ZSets)

Hashes

Bitmaps

HyperLogLogs

Streams

4. How does Redis store data in memory?
Redis stores all data in RAM as key-value pairs. The values are stored as encoded binary structures, optimized for speed. Optionally, it can write data to disk for persistence.

5. What is the difference between Redis and a traditional database?

| Feature  | Redis            | Traditional DB    |
| -------- | ---------------- | ----------------- |
| Storage  | In-memory        | Disk-based        |
| Speed    | Very fast        | Slower            |
| Schema   | Schema-less      | Schema-based      |
| Use case | Cache, real-time | Long-term storage |


6. What are the use cases of Redis?

Caching frequently accessed data

Session management

Leaderboards and rankings

Real-time analytics

Pub/Sub notifications

Message queues

Rate limiting

7. How do you install Redis on Linux?

sudo apt update
sudo apt install redis-server
sudo systemctl enable redis-server
sudo systemctl start redis-server


8. How to start and stop the Redis server?
Start:

redis-server


Stop:

redis-cli shutdown


9. What is the default port for Redis?
Redis runs on port 6379 by default.

10. What is a Redis key?
A Redis key is the unique identifier used to store and retrieve a value.
Example:

SET user:1 "Andre"
GET user:1


11. How do you set and get a key in Redis?

SET name "Chatie"
GET name


12. What is the maximum size of a Redis key or value?

Key: up to 512 MB

Value: up to 512 MB

13. What is the difference between SET and MSET commands?
SET sets a single key, MSET sets multiple keys.

SET name "Andre"
MSET name "Andre" age "25"


14. How do you delete a key in Redis?

DEL name


15. What is the difference between DEL and UNLINK?

DEL: Deletes key immediately (blocking).

UNLINK: Deletes key asynchronously (non-blocking, better for performance).

16. What are Redis strings, lists, sets, hashes, and sorted sets?

String – simple text or binary data.

List – ordered collection of strings.

Set – unique unordered strings.

Hash – field-value pairs (like JSON).

Sorted Set – unique strings ordered by a score.

17. How does Redis handle expiration of keys?
Redis allows setting a TTL (time-to-live) for any key.

SET session "abc" EX 60


This key expires after 60 seconds.

18. How do you check the TTL (time-to-live) of a key?

TTL session


19. What happens when a key expires?
Redis automatically removes the key from memory to free up space.

20. How do you persist data in Redis?
Redis supports two main persistence methods:

RDB (snapshotting) – saves data at intervals.

AOF (append-only file) – logs every operation.