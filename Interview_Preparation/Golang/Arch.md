Correct Mental Model (Very Important)

Go Runtime
 └── Go Scheduler
      ├── creates Ps (based on GOMAXPROCS)
      ├── assigns goroutines to P run queues
      ├── balances load (work stealing)
      └── manages Ms (OS threads)


How Goroutines works internally ?

“Goroutines are scheduled by the Go scheduler, which is part of the Go runtime. The scheduler creates logical processors (P) based on GOMAXPROCS and assigns goroutines to their run queues. P itself does not make scheduling decisions; it is a context used by the scheduler to manage execution and limit parallelism.”



First: who decides how many goroutines go to each P?
✅ The Go scheduler (inside Go runtime) decides

❌ Not the OS
❌ Not the programmer
❌ Not P by itself

👉 P does NOT “decide”
👉 P is just a resource owned by the scheduler

Correct Mental Model (Very Important)
Go Runtime
 └── Go Scheduler
      ├── creates Ps (based on GOMAXPROCS)
      ├── assigns goroutines to P run queues
      ├── balances load (work stealing)
      └── manages Ms (OS threads)


So:

The scheduler uses P as a scheduling context.

Now let’s fix this misconception 👇
❌ This looks intentional but is NOT manual:
P0 → G1 G2 G3
P1 → G4 G5 G6


You might think:

“Someone evenly distributed goroutines”

But actually 👇

How goroutines are REALLY assigned to Ps
Rule 1️⃣: Goroutine creation rule

When you do:

go f()


👉 The goroutine is placed into:

The local run queue of the P that is currently running

If no P → global queue

So distribution is opportunistic, not planned.

Example (realistic)

M0 is running with P0

You create 6 goroutines in a loop

for i := 0; i < 6; i++ {
    go work(i)
}


Result (likely):

P0 queue → G1 G2 G3 G4 G5 G6
P1 queue → (empty)


⚠️ No balancing yet

Rule 2️⃣: Load balancing (work stealing)

Later:

P1 becomes idle

Scheduler triggers work stealing

Now:

P0 → G1 G2 G3
P1 → G4 G5 G6


👉 This is done automatically
👉 You did NOT control this
👉 P did NOT decide — scheduler did

So who is the “brain”?
✅ Go Scheduler

The Go scheduler:

Lives inside Go runtime

Implements:

GMP model

Work stealing

Preemption

Fair scheduling

What exactly is Go Runtime?
Go Runtime = hidden system Go adds to your program

It includes:

Memory allocator

Garbage collector

Go scheduler

Goroutine stacks

Syscall handling

Channel implementation

When your program starts:

go run main.go


👉 Go runtime boots before main()

Go Scheduler vs OS Scheduler (important)
Aspect	Go Scheduler	OS Scheduler
Schedules	Goroutines	Threads
Preemption	Cooperative + async	Fully preemptive
Cost	Very cheap	Expensive
Awareness	App-level	System-level

👉 OS does not know goroutines exist

Does P schedule goroutines?
❌ No (this is subtle)

P:

Holds run queues

Tracks execution state

Limits parallelism

But:

Scheduler schedules, P is the scheduling context

Think:

Scheduler = manager

P = desk

M = worker

G = task

Is there a scheduling order?
Short answer: ❌ No strict order

Go scheduler:

Is not FIFO guaranteed

Is not round-robin strictly

Optimizes for:

Throughput

Fairness

Cache locality

Guarantees:

Eventually runnable goroutines will run

No starvation (mostly)

❌ No guarantee:

Execution order

Completion order