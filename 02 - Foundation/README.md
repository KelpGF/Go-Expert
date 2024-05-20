# Foundation

## Package

Every go file belong a package. The package declaration should be the first line of the file and you project must have a main package to be executable.

```go
package main

func main() {
    println("Hello, World!")
}

```

You can't use more than one package iin the same folder. The package name is the same as the folder name.

Also, files with the same package have access to each other.

## Declaration and Attribution

When declare a variable, you can use the type inference to declare the type of the variable or you can declare the type explicitly. If you don't attribute a value for it, Go apply a default value for the type.

```go
var b bool // type explicit, default value is false
var b2 = false // type inference, default value is false

var (
    c int // type explicit, default value is 0
    d: int = 1 // type explicit, default value is 1
    e string // type explicit, default value is ""
    f float64 // type explicit, default value is 0.0
)
```

In Go, you have:

- **const**: constant values, can't be changed.
- **var**: variables that can be changed.

### Short Declaration

You can use the short declaration to declare and attribute a value to a variable.

```go
a := 1 // ":" declares the variable and "=" attributes the value

a = 2 // the next attribution needn't the ":"
```

### Scopes

Go has three scopes: package, block and file.

- Package: variables declared at the package level are visible to all files in the package.
- Block: variables declared inside a block are visible only inside the block.
- File: variables declared at the file level are visible to all blocks in the file.

## Import

To import a package, you can use the `import` keyword.

```go
import "fmt"

// or

import (
    "fmt"
    "math"
)
```

## Creating Type

You can create your own type using the `type` keyword.

```go
type Celsius float64
```

And using FMT to print the type.

```go
package main

import "fmt"

type ID int

func main() {
    var y ID = 42
    a := "Hello, Go!"

    fmt.Printf("Type of y (%v): %T\n", y, y)
    fmt.Printf("Type of a (%v): %T\n", a, a)
}
```

Output:

``` text
Type of y (42): main.ID
Type of a (Hello, Go!): string
```

## Arrays

Arrays are fixed-size sequences of elements of the same type.

```go
// Declaring an array
var myArray [3]int

// Adding values to the array
myArray[0] = 1
myArray[1] = 2
myArray[2] = 3

// Getting the array size
arraySize := len(myArray)

// Accessing the array
fmt.Printf(myArray[0])
fmt.Printf(myArray[arraySize-1])

// Iterating over the array
for index, value := range myArray {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Slices

Slices works like arrays, but they are dynamic and have more features.

You can create a slices...

```go
// Using the make function
slice := make([]int, 3) // []int{0, 0, 0}

// Using the slice literal
slice := []int{1, 2, 3}
```

Slices have capacity and length. The length is the number of elements in the slice and the capacity is the number of elements in the underlying array, counting from the first element in the slice.

```go
func main() {
    slice := []int{1, 2, 3}

    fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
```

```text
// Output
len=3 cap=3 [1 2 3]
```

We can make a new slice from an existing slice.

```go
func main() {
    slice := []int{1, 2, 3}

    newSlice := slice[0:2] // [starts:ends]

    fmt.Printf("len=%d cap=%d %v\n", len(newSlice), cap(newSlice), newSlice)
}
```

```text
// Output
len=2 cap=3 [1 2]
```

If we change the start index of the slice, the capacity will change.

```go
func main() {
    slice := []int{1, 2, 3}

    newSlice := slice[1:3] // [starts:ends]

    fmt.Printf("len=%d cap=%d %v\n", len(newSlice), cap(newSlice), newSlice)
}
```

```text
// Output
len=2 cap=2 [2 3]
```

And we can append elements to the slice.

```go
func main() {
    slice := []int{1, 2, 3}

    newSlice := append(slice, 4)

    fmt.Printf("len=%d cap=%d %v\n", len(newSlice), cap(newSlice), newSlice)
}
```

```text
// Output
len=4 cap=6 [1 2 3 4]
```

As we can see, the capacity of the slice is doubled when we append a new element. It's because slices are references to the underlying array. So, when the capacity is greater than length, the slice create a new array with the same length.

Try create a slices with closest size to the expected size to avoid unnecessary allocations.

Finally, to iterate over a slice, we do like a array.

## Maps (Hash Tables)

Maps are key-value pairs.

```go
// Declaring a empty map
var myMap map[string]int
myMap = make(map[string]int)
myMap = map[string]int{}

// Declaring a map with values
numbersMap := map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
} // map[keyType]valueType

// Adding a new key-value pair
numbersMap["four"] = 4

// Getting a value from a key
value := numbersMap["one"]

// Deleting a key-value pair
delete(numbersMap, "one")

// Iterating over a map
for key, value := range numbersMap {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```

## Functions

Functions have a name, a list of parameters and a list of return values.

```go
// simple
func add(a int, b int) int {
    return a + b
}

// multiple parameters of the same type
func sub(a, b int) int { // a and b are int
    return a - b
}

// multiple return values
func multi(a, b int) (int, error) {
    if a == 0 || b == 0 {
        return 0, errors.New("Can't multiply by zero")
    }

    return a * b, nil
}
result, err := multi(2, 3)
```

### Variadic Functions

Variadic functions can receive a variable number of arguments.

```go
func sum(nums ...int) int {
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
```

### Blank Identifier

The blank identifier is used to discard values.

```go
_, err := multi(2, 3) // discard the first return value
```

### Closures (Anonymous Functions)

Closures are functions that can access variables from outside its body.

```go
func adder() func(int) int {
    sum := 0

    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    add := adder()

    fmt.Println(add(1)) // 1
    fmt.Println(add(2)) // 3
}
```

and Anonymous Functions are functions without a name. They can be assigned to a variable.

```go
func main() {
    func() {
        fmt.Println("Hello, Go!")
    }()
    print := func() {
        fmt.Println("Hello, Go!")
    }
    print()
}
```

## Structs

Structs are collections of fields. It have properties public and private.

```go
type Person struct {
    Name string
    Age int
    active bool // private property because it starts with lowercase
}

func main() {
    p := Person{
        Name: "John",
        Age: 30,
        active: true,
    }

    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

    p.Name = "Doe"
    p.Age = 31
    p.active = false // can't access this property

    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
```

and we can attach methods to structs informing the struct as the first parameter.

```go

type Person struct {
    Name string
    Age int
    active bool 
}
func (p Person) GetActive() bool {
 return p.active
}
func (p Person) Inactivate() {
 p.active = false
}

func main() {
    p := Person{
        Name: "John",
        Age: 30,
        active: true,
    }

    fmt.Printf("Name: %s, Age: %d, Active: %t\n", p.Name, p.Age, p.GetActive())

    p.Name = "Doe"
    p.Age = 31
    p.Inactivate() // active isn't changed. We will see how to change it in Pointers

    fmt.Printf("Name: %s, Age: %d, Active: %t\n", p.Name, p.Age, p.GetActive())
}
```

## Interfaces

Interfaces are collections of method signatures.

```go
type Shape interface {
    Area() float64
}
```

and you needn't to implement the interface explicitly. If the struct has the methods with the same signature, it implements the interface.
But the method signature must be the same as the interface.

```go
type Circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width float64
    Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Triangle struct {
    Base float64
    Height float64
}
func (t Triangle) Area(a int) float64 {
    return 0.5 * t.Base * t.Height
}

func showArea(s Shape) {
    fmt.Printf("Area: %f\n", s.Area())
}

func main() {
    c := Circle{Radius: 2}
    r := Rectangle{Width: 2, Height: 3}
    t := Triangle{Base: 2, Height: 3}

    showArea(c)
    showArea(r)
    showArea(t) // error: Triangle doesn't implement the Area method
}
```

## Pointers

Pointers are variables that its value is the memory address of another variable.

```go
func main() {
    a := 1 // a is a variable value, &a is the memory address of a
    var b *int // b is a pointer to an int
    b = &a // b receives the memory address of a

    fmt.Printf("a: %d, b: %p\n", a, b) // b will be shown the memory address of a
    fmt.Printf("a: %d, b: %d\n", a, *b) // *b is the value of a

    *b = 2 // change the value of a, it's the same as a = 2
    fmt.Printf("a: %d, b: %d\n", a, *b)
}
```

### Pointer as a Function Parameter

You can use pointers as function parameters to change the value of the variable.

```go
func changeValue(param *int) {
    *param = 2
}

func main() {
    a := 1

    changeValue(&a)

    fmt.Printf("a: %d\n", a)
}
```

```text
// Output
a: 2
```

We use pointer when we need to turn the variable mutable.

### Pointer to Struct

You can use pointers to structs to change the struct properties.

```go
    type BankAccount struct {
        Balance int
    }

    // This method have a copy of the struct, so it will not change the original value
    func (b BankAccount) simulateLoan(value int) int {
        b.Balance += value

        return b.Balance
    }

    // This method have a pointer to the struct, so it will change the original value
    func (b *BankAccount) Deposit(amount int) {
        b.Balance += amount
    }

    func main() {
        account := BankAccount{Balance: 1000}
        fmt.Println("Initial balance:", account.Balance)

        loan := account.simulateLoan(500)
        fmt.Println("Loan balance:", loan)
        fmt.Println("Original balance:", account.Balance)

        account.Deposit(100)
        fmt.Println("Deposit balance:", account.Balance)
    }
```

```text
// Output
Initial balance: 1000
Loan balance: 1500
Original balance: 1000
Deposit balance: 1100
```

Now, we solved the problem in Structs section.

## Typing

### Empty Interface

Empty interface is an interface with no methods. It can hold values of any type.

```go
func main() {
    var x interface{} = 10
    var y interface{} = "Hello"

    showType(x)
    showType(y)
}

func showType(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}
```

```text
// Output
Type: int, Value: 10
Type: string, Value: Hello
```

We must be careful when using empty interfaces because we lose the type safety.

### Type Assertion

Type assertion is used to convert an interface to another type.

```go
func main() {
  var x interface{} = 10

  value, ok := x.(int)

  if ok {
    fmt.Printf("Value: %d\n", value)
  } else {
    fmt.Println("Value is not an int")
  }

  var y interface{} = 10
  value2, ok2 := y.(string)
  if ok2 {
    fmt.Printf("Value2: %s\n", value2)
  } else {
    fmt.Println("Value2 is not an int")
  }

  value2 = y.(string) // throws a panic because y is not a string and we don't catch the ok value
}
```

```text
// Output
Value: 10
Value2 is not an int
panic: interface conversion: interface {} is int, not string

goroutine 1 [running]:
main.main()
        /home/kelp/dev/estudos/Go-Expert/02 - Foundation/11 - Typing/main.go:24 +0xbf
exit status 2
```

### Generics Types

Sometimes we need to create a function that works with different types. In Go, we do this:

```go
func Sum[T int | float64](m map[string]T) T {
  var soma T
  for _, val := range m {
    soma += val
  }

  return soma
}

func main() {
    mapInt := map[string]int{"a": 10, "b": 20, "c": 30}
    mapFloat := map[string]float64{"a": 10.5, "b": 20.5, "c": 30.5}

    fmt.Println("[Generics] Sum of mapInt:", Sum(mapInt))
    fmt.Println("[Generics] Sum of mapFloat:", Sum(mapFloat))
}
```

```text
// Output
[Generics] Sum of mapInt: 60
[Generics] Sum of mapFloat: 61.5
```

### Constraints

We can use constraints to limit the types that can be used in generics.

```go
type Number interface {
  int | float64
}
func Sum[T Number](m map[string]T) T {
  var soma T
  for _, val := range m {
    soma += val
  }

  return soma
}

func main() {
    mapInt := map[string]int{"a": 10, "b": 20, "c": 30}
    mapFloat := map[string]float64{"a": 10.5, "b": 20.5, "c": 30.5}

    fmt.Println("[Generics] Sum of mapInt:", Sum(mapInt))
    fmt.Println("[Generics] Sum of mapFloat:", Sum(mapFloat))
}
```

```text
// Output
[Generics] Sum of mapInt: 60
[Generics] Sum of mapFloat: 61.5
```

but constraints with same type doesn't work.

```go
type Number interface {
  int | float64
}

type myInt int

func ShowNumber[T Number](n T) {
  fmt.Println(n)
}

func main() {
    var a myInt = 10
    var b int = 20

    ShowNumber(a) // error: myInt and int are not the same type
    ShowNumber(b)
}
```

For it works, we can use `~`. Now, all types that are `int` or `float64` will be accepted.

```go
type Number interface {
  ~int | ~float64
}

type myInt int

func ShowNumber[T Number](n T) {
  fmt.Println(n)
}

func main() {
    var a myInt = 10
    var b int = 20

    ShowNumber(a) // now it works
    ShowNumber(b)
}
```

## Modules

We can create sub-packages to access on the main module. However, we need to create a `go.mod` file in the main module.

```bash
go mod init github.com/KelpGF/my-module-name
```

### Why use modules?

Because it's a way to manage dependencies in Go. If you don't use modules, Go will use the `GOPATH` to manage dependencies and you can't access your local modules.

```go
package main

import (
    "fmt"
    "github.com/KelpGF/my-module-name/sub_package"
)

func main() {
    fmt.Println(sub_package.Hello())
}
```

```go
// sub_package/sub_package.go
package sub_package

func Hello() string {
    return "Hello, Go!"
}
```

### Access Modifiers

When we create a package, we can use the access modifiers to define the visibility of the package.

- **Public**: the name starts with an uppercase letter.
- **Private**: the name starts with a lowercase letter.

```go
package sub_package

func hello() string {
    return "Hello, Go!"
}
func Hello() string {
    return "Hello, Go!"
}
```

```go
package main

import (
    "fmt"
    "github.com/KelpGF/my-module-name/sub_package"
)

func main() {
    fmt.Println(sub_package.Hello()) // works
    fmt.Println(sub_package.hello()) // error: can't access the private function
}
```

It's the same way to exports functions, variables, types and structs.

About structs, you can create private and public properties/methods.

```go
package sub_package

type Person struct {
    Name string
    age int // private property
}

// private method
func (p *Person) getName() string {
    return p.Name
}

func (p *Person) GetAge() int {
    return p.age
}

func (p *Person) SetAge(age int) {
    p.age = age
}
```

```go
package main

import (
    "fmt"
    "github.com/KelpGF/my-module-name/sub_package"
)

func main() {
    p := sub_package.Person{Name: "John", age: 30}
    p.SetAge(31) // works

    fmt.Println(p.Name) // works
    fmt.Println(p.age) // error: can't access the private property
    fmt.Println(p.GetAge()) // works
}
```

### Packages

Using modules, we can install packages that are not in the Go standard library.

```bash
go get github.com/google/uuid
```

If we have some uninstalled dependencies in the project, we can use the `go mod tidy` command to download the those or remove the unnecessary dependencies.

```bash
go mod tidy
```

## Loop

The `for` loop is the only loop in Go.

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

```go
collection := []int{1, 2, 3, 4, 5}
// or
collection := map[string]int{"a": 1, "b": 2, "c": 3}

for index, value := range collection {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

but we can create `while` and `do while` loops.

```go
func main() {
    i := 0

    // while: for with a condition
    for i < 10 {
        fmt.Println(i)
        i++
    }

    // do while: for with a break
    for {
        if i < 10 {
            continue // skip the next lines and go to the next iteration
        }

        fmt.Println(i)
        i++

        if i == 20 {
            break // exit the loop
        }
    }
}
```

and we have infinite loops.

```go
func main() {
    for {
        fmt.Println("Hello, Go!")
    }
}
```

## Conditional

Go has the `if` statement to create conditions.

```go
a := 10
b := 20

if a > b { // the condition must be a boolean statement
    fmt.Println("a is greater than b")
} else if a < b {
    fmt.Println("a is less than b")
} else {
    fmt.Println("a is equal to b")
}
```

and we can use the `switch` statement.

```go
func main() {
    a := 10

    switch a {
    case 1:
        fmt.Println("a is 1")
    case 2:
        fmt.Println("a is 2")
    case 3:
        fmt.Println("a is 3")
    default:
        fmt.Println("a is not 1, 2 or 3")
    }
}
```

## Build

To build a Go project, we use the `go build` command.

```bash
# without a module
go build [file]
```

```bash
# with a module
go build
```

We also can use pass what O.S. and architecture we want to build.

```bash
GOOS=linux GOARCH=amd64 go build [file]
```

[Here](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04) is a tutorial to build Go executables for multiple platforms.
