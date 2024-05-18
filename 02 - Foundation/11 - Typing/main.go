package main

import "fmt"

func main() {
	// AllowAnyType()
	// TypeAssertion()
	// Generics()
	Constraint()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func AllowAnyType() {
	var x interface{} = 10
	var y interface{} = "Hello"

	showType(x)
	showType(y)
}

func showType(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TypeAssertion() {
	var x interface{} = "Hello"
	println(x)          // show something like: 0xc0000b6010
	println(x.(string)) // works because x is a string

	var y interface{} = 10
	// result := y.(string) // Throws an error because y is not a string and we can't check the 'ok' value
	result, ok := y.(string)
	if ok {
		println(result)
	} else {
		println("Error")
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func withoutGenerics() {
	sumInt := func(m map[string]int) int {
		var soma int
		for _, val := range m {
			soma += val
		}

		return soma
	}

	sumFloat := func(m map[string]float64) float64 {
		var soma float64
		for _, val := range m {
			soma += val
		}

		return soma
	}

	mapInt := map[string]int{"a": 10, "b": 20, "c": 30}
	mapFloat := map[string]float64{"a": 10.5, "b": 20.5, "c": 30.5}
	fmt.Println("Sum of mapInt:", sumInt(mapInt))
	fmt.Println("Sum of mapFloat:", sumFloat(mapFloat))
}

func withGenerics() {
	mapInt := map[string]int{"a": 10, "b": 20, "c": 30}
	mapFloat := map[string]float64{"a": 10.5, "b": 20.5, "c": 30.5}

	fmt.Println("[Generics] Sum of mapInt:", Sum(mapInt))
	fmt.Println("[Generics] Sum of mapFloat:", Sum(mapFloat))
}

func Generics() {
	withoutGenerics()
	withGenerics()
}
func Sum[T int | float64](m map[string]T) T {
	var soma T
	for _, val := range m {
		soma += val
	}

	return soma
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type myInt int
type myFloat float64
type NumberRestrict interface {
	int | float64
}
type NumberGeneral interface {
	~int | ~float64
}

func SumRestrict[T NumberRestrict](m map[string]T) T {
	var soma T
	for _, val := range m {
		soma += val
	}

	return soma
}

func SumGeneral[T NumberGeneral](m map[string]T) T {
	var soma T
	for _, val := range m {
		soma += val
	}

	return soma
}

func Constraint() {
	mapInt := map[string]myInt{"a": 10, "b": 20, "c": 30}
	mapFloat := map[string]myFloat{"a": 10.5, "b": 20.5, "c": 30.5}

	// Error because myInt and myFloat are not int or float64
	// fmt.Println("[Constraint] Sum of mapInt:", SumRestrict(mapInt))
	// fmt.Println("[Constraint] Sum of mapFloat:", SumRestrict(mapFloat))

	// with ~int and ~float64, it allows types that are created from int and float64
	fmt.Println("[Constraint] Sum of mapInt:", SumGeneral(mapInt))
	fmt.Println("[Constraint] Sum of mapFloat:", SumGeneral(mapFloat))
}
