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
