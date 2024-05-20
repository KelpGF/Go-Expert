package main

import (
	"fmt"
	"go-modules/math"
)

func main() {
	basic()
	structTest()
}

func basic() {
	s := math.Sum(1, 2)

	fmt.Println("Sum of 1 and 2 is", s)
	fmt.Println("A is", math.A)
	// fmt.Println("b is", math.b) // This will not work as B is not exported
	math.Example() // We can't access the sub function, but it can be called from the math package
}

func structTest() {
	p := math.Person{
		Name: "John",
	}
	p.SetAge(25)

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.GetAge())
	// p.getName() // This will not work as setName is not exported
}
