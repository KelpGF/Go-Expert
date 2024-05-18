package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 20

	returnNothing()

	fmt.Println("Sum of", a, "and", b, "is", sum(a, b))

	sub, err := subtract(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Subtraction of", a, "and", b, "is", sub)

	fmt.Println("Sum of 1, 2, 3, 4, 5 is", sumVariadic(1, 2, 3, 4, 5))

	closures()

	add := adder()
	fmt.Println(add(10))
	fmt.Println(add(20))
}

func returnNothing() {
	fmt.Println("I don't return anything")
}

func sum(a, b int) int {
	return a + b
}

func subtract(a int, b int) (int, error) {
	result := a - b

	if result < 0 {
		return 0, errors.New("Result is negative")
	}

	return result, nil
}

func sumVariadic(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func closures() {
	total := func() int {
		return sumVariadic(10, 20, 30, 40, 50)
	}()
	divide := func(a, b int) int {
		return a / b
	}
	fmt.Print("Total is ", total)
	fmt.Print(" and divided by 2 is ", divide(total, 2), "\n")
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}
