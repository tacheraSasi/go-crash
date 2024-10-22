Let's dive deep into some fundamental concepts of Go (Golang). I'll explain each concept, followed by related practice questions with solutions and detailed explanations.

---

## 1. **Structs**

### **What are Structs?**
Structs in Go are used to group together variables of different types under a single name. A struct represents a collection of fields, which can be of different data types. You can think of it like a blueprint for an object in object-oriented programming. However, Go is not an object-oriented language, so it uses structs to group related data.

### **Syntax:**
```go
type StructName struct {
    Field1 DataType1
    Field2 DataType2
}
```

You can create an instance of the struct and initialize it with values:
```go
person := StructName{Field1: value1, Field2: value2}
```

---

### **Practice Question 1: Basic Employee Info**

**Question:**
Create a struct called `Employee` with fields `Name`, `Age`, and `Position`. Write a function that takes an employee and prints their details.

**Solution:**
```go
package main

import "fmt"

type Employee struct {
    Name     string
    Age      int
    Position string
}

func printEmployeeDetails(emp Employee) {
    fmt.Printf("Name: %s, Age: %d, Position: %s\n", emp.Name, emp.Age, emp.Position)
}

func main() {
    emp := Employee{Name: "Alice", Age: 30, Position: "Engineer"}
    printEmployeeDetails(emp)
}
```

**Explanation:**
- We define the `Employee` struct with three fields: `Name`, `Age`, and `Position`.
- The `printEmployeeDetails` function accepts an `Employee` struct and prints the values of the fields.
- We create an `Employee` instance in `main()` and pass it to the function.

---

## 2. **Pointers**

### **What are Pointers?**
Pointers in Go hold the memory address of a variable rather than the actual value. You can think of them as "pointers" to where the data is stored in memory. With pointers, you can directly manipulate the memory, which allows you to change the value of a variable even outside the function where it's being modified.

- `&` gives you the address of a variable (pointer to that variable).
- `*` allows you to access the value stored at a given pointer.

---

### **Practice Question 2: Modify Value Using Pointers**

**Question:**
Create a function that takes a pointer to an integer and modifies its value.

**Solution:**
```go
package main

import "fmt"

func modifyValue(num *int) {
    *num = 50 // Modifying the value of the integer at the pointer address
}

func main() {
    value := 10
    fmt.Println("Before:", value) // 10
    modifyValue(&value)
    fmt.Println("After:", value) // 50
}
```

**Explanation:**
- The function `modifyValue` accepts a pointer to an integer (`*int`), allowing it to modify the value at that address.
- The `&value` operator is used to pass the memory address of the variable `value` to the function.
- Inside `modifyValue`, `*num = 50` changes the value at the memory location to 50.

---

## 3. **Closures**

### **What are Closures?**
Closures in Go are functions that "close over" or "capture" variables from their surrounding lexical scope. This means that the inner function can access and modify variables defined in the outer function even after the outer function has finished executing.

Closures are useful when you want to keep track of some state, like a counter, across multiple function calls.

---

### **Practice Question 3: Counter Example**

**Question:**
Write a closure in Go that acts as a counter. Each time you call it, it should return the next number.

**Solution:**
```go
package main

import "fmt"

func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := createCounter()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

**Explanation:**
- `createCounter` is an outer function that defines the variable `count`.
- The inner function (returned by `createCounter`) increments and returns `count` every time it's called.
- This closure "remembers" the value of `count`, even after `createCounter` has finished executing.

---

## 4. **Interfaces**

### **What are Interfaces?**
Interfaces in Go provide a way to define a set of methods that a type must implement. It’s a way to specify behavior without specifying how that behavior is achieved. A type is said to implement an interface if it defines all the methods declared by the interface.

An interface can hold any type that implements its methods. This allows for polymorphism, where different types can be treated as the same interface type.

---

### **Practice Question 4: Shape Interface with Area Calculation**

**Question:**
Create an interface `Shape` that has a method `Area()`. Then implement two structs: `Circle` and `Rectangle` that calculate their area.

**Solution:**
```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 3, Height: 4}
    
    printArea(c) // 78.54
    printArea(r) // 12.00
}
```

**Explanation:**
- We define an interface `Shape` with one method `Area`.
- `Circle` and `Rectangle` are structs that implement the `Area` method.
- The `printArea` function accepts any type that implements the `Shape` interface and prints the area.

---

## 5. **Slices**

### **What are Slices?**
Slices are a flexible, dynamic-sized array in Go. Unlike arrays, slices can grow and shrink. They are built on top of arrays but provide more flexibility since their size isn't fixed. Slices store a reference to an underlying array, and changing the elements of a slice will modify the array.

Slices can be created in various ways:
- Using a literal: `nums := []int{1, 2, 3}`
- Using `make`: `nums := make([]int, 5)`

---

### **Practice Question 5: Manipulate Slice Elements**

**Question:**
Create a slice of integers and write a function to double each element in the slice.

**Solution:**
```go
package main

import "fmt"

func doubleSlice(nums []int) {
    for i := range nums {
        nums[i] *= 2
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Before:", numbers) // [1 2 3 4 5]
    doubleSlice(numbers)
    fmt.Println("After:", numbers) // [2 4 6 8 10]
}
```

**Explanation:**
- We define a slice `numbers` and pass it to the `doubleSlice` function.
- Since slices in Go are passed by reference, any changes made inside the function directly affect the original slice.
  
---

## 6. **Maps**

### **What are Maps?**
Maps in Go are used to store key-value pairs. They are similar to dictionaries in Python or hash maps in other languages. Keys must be unique, and each key is associated with exactly one value.

A map can be created using:
- A map literal: `m := map[string]int{"Alice": 25}`
- `make`: `m := make(map[string]int)`

---

### **Practice Question 6: Create and Access a Map**

**Question:**
Create a map that stores the names and ages of people. Write a function to print each person's name and age.

**Solution:**
```go
package main

import "fmt"

func printAges(people map[string]int) {
    for name, age := range people {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}

func main() {
    people := map[string]int{
        "Alice": 30,
        "Bob":   25,
        "Eve":   35,
    }
    
    printAges(people)
}
```

**Explanation:**
- We define a map `people` where the keys are strings (names) and the values are integers (ages).
- The `printAges` function iterates over the map using `range` and prints each name and age.

---

By working through these fundamental concepts and practice problems, you’ll strengthen your understanding of Go's core features and how they are used in real-world programming!
