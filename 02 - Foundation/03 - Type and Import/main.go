package main

import "fmt"

type ID int

func main() {
	var y ID = 42
	a := "Hello, Go!"

	fmt.Printf("Type of y (%v): %T\n", y, y)
	fmt.Printf("Type of a (%v): %T\n", a, a)
}
