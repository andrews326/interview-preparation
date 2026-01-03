🟢 Golang Concurrency — Complete Q&A for Interviews
🌱 Basic Level

1. What is concurrency in Go?
Concurrency allows multiple tasks to make progress independently, potentially at the same time.

2. What is a goroutine?
A lightweight thread managed by Go runtime. Created using:

go myFunction()


3. How do goroutines communicate?
Primarily using channels for safe communication.

4. What is a channel?
A typed conduit for sending and receiving data between goroutines.

5. What is a buffered vs unbuffered channel?

Unbuffered: send blocks until received.

Buffered: send blocks only if buffer is full.

6. How do you wait for multiple goroutines?
Use sync.WaitGroup.

7. What is a race condition?
Occurs when multiple goroutines access shared data simultaneously, and at least one modifies it.

8. How do you prevent race conditions?

Use channels for communication

Or sync.Mutex / sync.RWMutex for shared data

9. What is a semaphore (basic concept)?
A concurrency control mechanism that limits the number of goroutines accessing a resource simultaneously. Implemented using buffered channels in Go.

10. What is select in Go?
Allows waiting on multiple channel operations, proceeding with the first ready channel.

⚙️ Intermediate Level

1. How do you implement a worker pool?

Create a task channel.

Start N worker goroutines reading from the channel.

Send tasks to the channel; workers process concurrently.

2. How do you implement a timeout in goroutines?
Use select with time.After to enforce a timeout.

3. What is context.Context?
Used to cancel goroutines, set deadlines, and pass values across API boundaries.

4. How do you use a semaphore in Go?
Use a buffered channel as a counting semaphore:

sem := make(chan struct{}, 3) // max 3 concurrent goroutines
sem <- struct{}{}             // acquire
<-sem                         // release


5. How do you avoid goroutine leaks?

Use context cancellation

Close channels

Avoid indefinite blocking

6. How do you monitor the number of goroutines?
Use runtime.NumGoroutine().

7. How do you handle panics in goroutines?
Use defer recover() to catch panics and prevent program crash.

8. How do you implement rate limiting?
Use a channel-based semaphore or time.Ticker to limit goroutine execution.

9. What are mutex vs RWMutex?

Mutex: exclusive lock for read/write

RWMutex: multiple readers allowed, only one writer

10. How do you ensure exactly-once processing in worker pools?

Use channels for task assignment

Remove tasks from queue after successful processing

Handle retries carefully with custom logic

🚀 Advanced Level

1. What is the Go scheduler?

A user-space scheduler that multiplexes goroutines onto OS threads.

Manages goroutine states: running, runnable, waiting, or sleeping.

2. How do goroutines differ from threads?

Lighter, smaller stack

Managed by Go runtime, not OS

Thousands can coexist easily

3. How do you prevent deadlocks?

Avoid circular waits

Acquire locks in a consistent order

Use channels carefully and avoid blocking indefinitely

4. How do you implement fan-out/fan-in concurrency?

Fan-out: multiple goroutines consume tasks from a single channel

Fan-in: multiple goroutines send results to a single channel

5. How do you implement semaphores in high concurrency?

Use buffered channels with capacity = max concurrent goroutines

Acquire by sending into the channel, release by reading from the channel

6. How do you combine context with worker pools?

Pass context.Context to each worker

Workers stop processing if context is canceled

7. How do you optimize performance in concurrent Go programs?

Avoid creating unbounded goroutines

Use buffered channels and worker pools

Minimize shared state and locks

Monitor goroutine count and memory usage

8. How do you handle cancellation across multiple goroutines?

Use a single context.WithCancel or context.WithTimeout

Goroutines listen to ctx.Done() channel to exit gracefully

9. How do you implement backpressure?

Limit the size of buffered channels

Use rate limiting or semaphores to prevent resource exhaustion

10. What are best practices for concurrency in Go?

Prefer channels over shared memory

Avoid unnecessary goroutines

Use WaitGroup to wait for completion

Use context for cancellation

Handle errors and panics gracefully

Avoid deadlocks and race conditions

🔹 Enhanced Rapid-Fire Questions

How do you detect race conditions in Go?

go run -race main.go


Detects unsafe access to shared memory.

How do you stop a goroutine gracefully?
Use context.Context and check ctx.Done().

What is the default stack size of a goroutine?
Starts at 2 KB, grows dynamically as needed.

Can goroutines communicate without channels?
Yes, using shared memory with sync.Mutex / sync.RWMutex.

Difference between unbuffered and buffered channels?

Unbuffered: synchronous (blocks until received)

Buffered: asynchronous up to buffer capacity

How do you cancel multiple goroutines?
Share the same context.Context and cancel it.

What is a deadlock?
Two or more goroutines wait for each other indefinitely.

How do you implement a simple semaphore?
Use a buffered channel as a counting semaphore (see above).

Difference between sync.Mutex and sync.RWMutex?
Mutex = exclusive lock; RWMutex = multiple readers, single writer.

How do you wait for multiple goroutines?
Use sync.WaitGroup and wg.Wait().