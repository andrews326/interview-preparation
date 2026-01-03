# What is a goroutine?
A: A goroutine is a lightweight thread managed by the Go runtime, created using the go keyword.

# How are goroutines different from OS threads?
A: Goroutines are lighter and multiplexed by the Go scheduler; you can run thousands in a single process, unlike heavier OS threads.

# How do you create a goroutine?
A: By prefixing a function call with go, like go functionName().

# Can goroutines run concurrently or in parallel?
A: Both. Concurrency means tasks overlap in execution, parallelism means tasks run on multiple cores simultaneously.

# What happens if a goroutine panics?
A: That goroutine stops, but others continue. You can recover using the recover() function.

# What is a channel?
A: A channel is a communication mechanism to safely pass data between goroutines.

# How do you send and receive values in a channel?
A: Send with ch <- val, receive with val := <-ch.

# What is the difference between buffered and unbuffered channels?
A: Buffered channels have capacity and don’t block until full; unbuffered channels block until a receiver is ready.

# How do you close a channel?
A: Use close(ch). Sending to a closed channel will cause a panic.

# How do you detect a closed channel?
A: Use val, ok := <-ch; ok is false if the channel is closed.

# What is the select statement?
A: select waits on multiple channel operations and can include timeout or default cases.

# What is a directional channel?
A: A channel restricted to send-only (chan<- int) or receive-only (<-chan int).

# What is a WaitGroup?
A: WaitGroup waits for a collection of goroutines to finish before continuing.

# What is a Mutex?
A: A Mutex provides mutual exclusion to safely access shared data.

# What are race conditions?
A: Race conditions occur when multiple goroutines access shared data unsafely.

# How do you detect race conditions?
A: Run the program with the race detector: go run -race main.go.

# Difference between Mutex and Semaphore?
A: Mutex allows only one goroutine to access a resource; a Semaphore limits access to a set number of concurrent goroutines.

# How do you implement a semaphore in Go?
A: Using a buffered channel with capacity equal to the number of permits:

sem := make(chan struct{}, maxConcurrent)
sem <- struct{}{} // acquire
<-sem            // release


# Difference between sync.Mutex and sync.RWMutex?
A: RWMutex allows multiple readers simultaneously but only one writer at a time.

# What is Fan-out / Fan-in?
A: Fan-out = multiple goroutines handling tasks; Fan-in = combining results into one channel.

# What is a Worker Pool?
A: A pool of goroutines that limits the number of concurrent tasks being processed.

# What is the Pipeline pattern?
A: Data flows through multiple stages connected by channels.

# How do you limit concurrent goroutines?
A: Use a semaphore (buffered channel) or a worker pool.

# How to cancel a goroutine using context?
A: Use ctx, cancel := context.WithCancel(...) and check <-ctx.Done() inside the goroutine.

# What are the types of context in Go?
A: WithCancel = manual cancel, WithTimeout = auto-cancel after a duration, WithDeadline = cancel at a specific time.

# How to safely pass data between goroutines?
A: Use channels; avoid sharing memory without proper synchronization.

# How to handle goroutine leaks?
A: Ensure goroutines exit, cancel context when done, avoid blocking channels.

# How to make concurrent API calls efficiently?
A: Use a worker pool, semaphore, rate-limiting, and context for timeout handling.

# What is the impact of too many goroutines?
A: High memory usage, scheduler overhead, and potential slowdowns.

# How does Go scheduler handle goroutines?
A: It uses an M:N scheduler to multiplex many goroutines onto fewer OS threads.

# How is memory managed for goroutines?
A: Each goroutine starts with a small (~2KB) stack, which grows dynamically; garbage collector cleans unused objects.

# How can you optimize high-concurrency programs?
A: Use worker pools, minimize allocations, and use channels efficiently.

# Parallelism:

Definition: Parallelism is when multiple tasks are executed simultaneously on different CPU cores (or processors).

Key Point: True parallelism can only happen when you have multiple CPU cores or processors. Each core can run one task at the same time (in parallel) without needing to wait for the other cores.


# Concurrency:

Definition: Concurrency is when multiple tasks are handled at the same time, but not necessarily simultaneously. It’s about making progress on many tasks, even if only one task is actually running at any given moment.

Key Point: In concurrent systems, tasks can be executed on a single core, but the OS switches between them rapidly, creating the illusion that they're running at the same time. This is time-sharing or context switching.





# FLASHCARD VERSION

1. What is a goroutine? – Lightweight Go thread
2. Goroutine vs OS thread? – Lighter, multiplexed vs heavier
3. How to create? – go functionName()
4. Concurrent or parallel? – Both: overlap vs multiple cores
5. Panic in goroutine? – Stops goroutine, use recover()
6. Channel? – Goroutine communication
7. Send/receive? – ch <- val / <-ch
8. Buffered vs unbuffered? – Buffered = capacity, unbuffered blocks
9. Close channel? – close(ch)
10. Detect closed? – val, ok := <-ch
11. select? – Multiple channel ops
12. Directional channel? – Send-only / receive-only
13. WaitGroup? – Wait for goroutines
14. Mutex? – Mutual exclusion
15. Race condition? – Unsafe concurrent access
16. Detect race? – go run -race
17. Mutex vs Semaphore? – 1 vs limited goroutines
18. Implement semaphore? – Buffered channel
19. Mutex vs RWMutex? – Readers/writer distinction
20. Fan-out / Fan-in? – Split/merge tasks
21. Worker Pool? – Limit concurrent tasks
22. Pipeline? – Multi-stage data flow
23. Limit goroutines? – Semaphore / worker pool
24. Cancel with context? – ctx.Done()
25. Context types? – Cancel, Timeout, Deadline
26. Pass data safely? – Channels, not shared memory
27. Handle goroutine leaks? – Exit, cancel, avoid block
28. Efficient concurrent API calls? – Worker pool, semaphore
29. Too many goroutines? – Memory & scheduler overhead
30. Go scheduler & memory? – M:N scheduler, dynamic stack


# How Does Concurrency Work on One Core?

On a system with one core, only one task can be executed at a time, but the operating system (or in Go's case, the Go runtime scheduler) switches between tasks very quickly. This is called time-sharing or context switching, and it gives the illusion that multiple tasks are running at the same time.

# When you run 5 jobs concurrently on a single core, the OS (or Go runtime) will:

Start executing one job (task).

Pause it after a short time (called a time slice).

Switch to the next job and start executing it.

Repeat the process, going back to previously paused tasks, so that each job gets a slice of CPU time.

Time-Slicing and Context Switching

The time slice is the amount of time the OS gives each task to run before switching to another task. The time slice is usually very small (e.g., a few milliseconds).

This switching happens so quickly that it appears like all tasks are running concurrently, even though only one task is actually executing at any given time.

Example of 5 Jobs Running on One Core

Let’s say you have 5 jobs (Job A, Job B, Job C, Job D, and Job E) that need to run concurrently on a system with one CPU core. Here’s how it might work in practice:

Job A starts running, and the OS scheduler gives it a time slice (e.g., 10 milliseconds).

After 10 milliseconds, the scheduler pauses Job A and starts Job B.

Job B runs for its time slice (10 milliseconds).

After Job B’s time slice, it gets paused, and Job C starts.

This continues, with the OS switching between Job A, Job B, Job C, Job D, and Job E in quick succession.