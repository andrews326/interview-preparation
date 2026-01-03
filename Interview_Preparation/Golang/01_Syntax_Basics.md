🟢 Golang — Basic Q&A (for 2 years experience refresh)
1️⃣ Introduction & Basics

Q1: What is Go (Golang), and why is it used in DevOps/backend systems?
A: Go is a statically typed, compiled language by Google, known for its concurrency support, simplicity, and performance. It’s widely used in DevOps tools like Docker, Kubernetes, and Terraform due to its efficiency and fast compilation.

Q2: What are the key features of Go?
A:

Fast compilation

Built-in concurrency (goroutines & channels)

Garbage collection

Simple syntax

Cross-compilation support

Strong standard library

Q3: What is a package in Go?
A: A package is a collection of Go source files in the same directory that are compiled together. Every Go program starts with a main package and a main() function.

2️⃣ Variables, Constants, and Data Types

Q4: How do you declare variables in Go?
A:

var name string = "Chatie"
age := 25 // short declaration


Q5: What is the difference between var and :=?
A:

var can be used at the package or function level.

:= is shorthand for variable declaration but can only be used inside functions.

Q6: How do you declare constants?
A:

const Pi = 3.14


Q7: What are basic data types in Go?
A:

Numeric: int, float64

Boolean: bool

String: string

Derived: array, slice, map, struct, interface

3️⃣ Control Structures

Q8: How does Go handle loops?
A: Go has only one looping construct: for.

for i := 0; i < 5; i++ {
    fmt.Println(i)
}


Q9: How does Go handle conditional statements?
A:

if x > 10 {
    fmt.Println("big")
} else if x == 10 {
    fmt.Println("equal")
} else {
    fmt.Println("small")
}


Q10: How does a switch statement work in Go?
A:

switch day {
case "Mon":
    fmt.Println("Start")
case "Sun":
    fmt.Println("Rest")
default:
    fmt.Println("Workday")
}

4️⃣ Functions

Q11: How do you define a function in Go?
A:

func add(a int, b int) int {
    return a + b
}


Q12: Can Go return multiple values?
A: Yes.

func divide(a, b int) (int, int) {
    return a / b, a % b
}


Q13: What are named return values?
A: You can name return variables and return without specifying them again.

func rectangle(length, width int) (area int) {
    area = length * width
    return
}

5️⃣ Arrays, Slices, and Maps

Q14: Difference between an array and a slice in Go?
A:

Array: Fixed size.

Slice: Dynamic, built on top of arrays.
Example:

arr := [3]int{1, 2, 3}
slice := []int{1, 2, 3, 4}


Q15: How to create and use a map in Go?
A:

person := map[string]string{
    "name": "Andre",
    "role": "DevOps",
}
fmt.Println(person["name"])


Q16: How to check if a key exists in a map?
A:

val, ok := person["age"]
if ok {
    fmt.Println(val)
} else {
    fmt.Println("key not found")
}

6️⃣ Structs and Methods (Intro)

Q17: What is a struct in Go?
A: A struct is a collection of fields.

type Employee struct {
    Name string
    Age  int
}


Q18: How to attach a method to a struct?
A:

func (e Employee) Display() {
    fmt.Println(e.Name, e.Age)
}