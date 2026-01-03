🟡 Golang — Intermediate Q&A (for 2 years experience)
1️⃣ Goroutines

# What is a goroutine?
A: A goroutine is a lightweight thread managed by the Go runtime. It’s created using the go keyword and runs concurrently with other goroutines.

go fmt.Println("Hello, Goroutine!")


# How are goroutines different from OS threads?
A:

Goroutines are managed by Go runtime, not the OS.

They use less memory (a few KB) compared to OS threads (MBs).

Thousands of goroutines can run concurrently in a single process.

# How do you synchronize goroutines?
A: Using channels, WaitGroups, or mutexes (from the sync package).

2️⃣ Channels

# What is a channel in Go?
A: A channel is used to communicate between goroutines.

ch := make(chan int)
go func() {
    ch <- 10 // send
}()
fmt.Println(<-ch) // receive


# Difference between buffered and unbuffered channels?
A:

Unbuffered: Send blocks until another goroutine receives.

Buffered: Has capacity; send blocks only when full.

ch := make(chan int, 2)
ch <- 1
ch <- 2


# What happens if you close a channel?
A:

Receivers can still read remaining values.

Further sends cause a panic.

close(ch)


# How do you check if a channel is closed?
A:

val, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}


# What is a select statement?
A: It’s used to wait on multiple channel operations.

select {
case msg := <-ch1:
    fmt.Println("Received:", msg)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout")
}

3️⃣ Interfaces

# What is an interface in Go?
A: An interface defines a set of method signatures. A type implements an interface automatically if it has those methods.

type Speaker interface {
    Speak() string
}

type Human struct{}

func (h Human) Speak() string {
    return "Hello!"
}


# Can a type implement multiple interfaces?
A: Yes, a struct can implement multiple interfaces simply by defining the required methods.

Q11: What is an empty interface (interface{})?
A: It can hold values of any type — it’s Go’s equivalent of any or object in other languages.

var x interface{}
x = 10
x = "Chatie"


# What is type assertion in Go?
A: Used to extract the underlying value from an interface.

val, ok := x.(string)
if ok {
    fmt.Println("String:", val)
}

4️⃣ Error Handling

# How does Go handle errors?
A: Go uses explicit error values instead of exceptions.

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}


Q14: What is the difference between panic and error?
A:

error: Used for expected issues that can be handled.

panic: Used for serious, unexpected errors that should stop execution.

Q15: What is recover() in Go?
A: It’s used to handle a panic and continue execution.

func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    fmt.Println(a / b)
}

5️⃣ Sync Package (Concurrency Control)

Q16: What is a WaitGroup?
A: A WaitGroup waits for multiple goroutines to finish.

var wg sync.WaitGroup
wg.Add(2)
go func() { defer wg.Done(); fmt.Println("Task 1") }()
go func() { defer wg.Done(); fmt.Println("Task 2") }()
wg.Wait()


Q17: What is a Mutex and when is it used?
A: A Mutex provides mutual exclusion for shared data.

var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()


Q18: What are race conditions, and how do you detect them?
A: Race conditions happen when multiple goroutines access shared data without synchronization.
Detect using:

go run -race main.go