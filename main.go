package main

import (
	"fmt"
	"time"
)

type CalcArea interface {
	Area() float64
}

type CalcPerimetro interface {
	Perimetro() float64
}

type Forma interface {
	CalcArea
	CalcPerimetro
}

type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}
func (q Quadrado) Perimetro() float64 {
	return q.Lado * 4
}

func printForma(f Forma) {
	fmt.Println(f.Area(), f.Perimetro())
}

func printArea(f CalcArea) {
	fmt.Println(f.Area())
}

func printPerimetro(f CalcPerimetro) {
	fmt.Println(f.Perimetro())
}

func main() {
	var forma Forma
	forma = Quadrado{Lado: 2}

	printForma(forma)
	printArea(forma)
	printPerimetro(forma)
}

func main2() {
	// for i := 0; i < 10; i++ {
	// 	println(i)
	// }

	// numberArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for index, value := range numberArr {
	// 	println(index, value)
	// }

	// for {
	// 	println("Hello, World!")
	// 	time.Sleep(1 * time.Second)
	// }

	// condition := true
	// for condition {
	// 	println("Hello, World!")
	// 	time.Sleep(1 * time.Second)
	// 	condition = false
	// }

	condition := true

	for {
		println("Hello, World! kelvin")
		time.Sleep(1 * time.Second)

		condition = false
		if !condition {
			break
		}
	}
}
