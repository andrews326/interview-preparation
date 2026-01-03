1️⃣ Concurrency Patterns

Q1: What are common concurrency patterns in Go?
A:

Fan-out/Fan-in

Worker pools

Pipeline pattern

Select with timeout/cancellation

Q2: Explain the Fan-out/Fan-in pattern.
A:

Fan-out: Launch multiple goroutines to handle input data concurrently.

Fan-in: Combine results from multiple goroutines into a single channel.

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}


Q3: How would you design a worker pool in Go?
A:

Create a channel for jobs and results.

Spawn a fixed number of worker goroutines.

Wait for all workers to complete using WaitGroup.

👉 This pattern prevents resource exhaustion and balances load.

2️⃣ Context Package

Q4: What is the context package used for?
A: To manage deadlines, cancellations, and request-scoped values across API boundaries and goroutines.
Common in web servers, APIs, and distributed systems.

Q5: Example: How do you cancel a goroutine using context?
A:

ctx, cancel := context.WithCancel(context.Background())
go func() {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Goroutine stopped")
            return
        default:
            fmt.Println("Running...")
            time.Sleep(time.Second)
        }
    }
}()
time.Sleep(3 * time.Second)
cancel()


Q6: Difference between WithCancel, WithTimeout, and WithDeadline?
A:

WithCancel → manual cancellation.

WithTimeout → auto-cancels after duration.

WithDeadline → cancels at a specific time.

3️⃣ Memory & Performance

Q7: How does Go manage memory?
A: Go uses garbage collection (GC) to automatically reclaim memory from unused objects. Developers should still minimize allocations and avoid long-lived references.

Q8: What is the escape analysis in Go?
A: Determines whether a variable can be allocated on the stack or must “escape” to the heap.
Use:

go build -gcflags '-m'


to analyze it.

Q9: What’s the difference between stack and heap allocation in Go?
A:

Stack: Fast, automatically freed when function returns.

Heap: Managed by GC, slower, used for long-lived data.

Q10: How do you reduce memory allocations in Go?
A:

Use object pooling (sync.Pool).

Reuse slices instead of reallocating.

Avoid unnecessary conversions (like string ↔ byte).

4️⃣ Advanced Interfaces & Reflection

Q11: What is reflection in Go?
A: Reflection allows you to inspect and modify objects at runtime using the reflect package.

Example:

v := reflect.ValueOf(10)
fmt.Println(v.Kind()) // int


Q12: When should you avoid reflection?
A:

It’s slower and less type-safe.

Use it only when writing frameworks or generic utilities (like serialization, logging, testing).

Q13: What are generics in Go (since Go 1.18)?
A: Generics allow functions and types to operate on different data types using type parameters.

func Sum[T int | float64](a, b T) T {
    return a + b
}

5️⃣ Real-World & DevOps Scenarios

Q14: How is Go used in DevOps?
A:

Building CLIs (Terraform, kubectl, Docker)

Writing microservices for CI/CD systems

Automating AWS/K8s operations

Creating monitoring agents or log processors

Q15: Example: How to handle concurrent API calls efficiently?
A: Use a worker pool or rate limiter pattern with goroutines + channels + context.

Q16: What is the best way to limit concurrency?
A: Using a buffered channel or semaphore-like control.

sem := make(chan struct{}, 5)
for _, task := range tasks {
    sem <- struct{}{}
    go func(t Task) {
        defer func() { <-sem }()
        process(t)
    }(task)
}


Q17: How do you handle timeouts in external calls (e.g., API requests)?
A: Use context.WithTimeout with HTTP client:

ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)


Q18: What tools do you use to test or benchmark Go programs?
A:

go test for unit tests.

go test -bench . for benchmarks.

pprof for CPU/memory profiling.

Q19: How do you find data races or concurrency bugs?
A:

go run -race main.go


This helps detect unsafe concurrent access.

Q20: How do you structure a Go project for large applications?
A:

/cmd          → main apps
/pkg          → reusable code
/internal     → private packages
/api          → HTTP handlers
/config       → configs/env