package main

import (
	"fmt"

	"github.com/KelpGF/Go-Expert/06-Packaging/01-Intro/math"
)

func main() {
	mathPublic := math.MathPublic{A: 1, B: 2}
	fmt.Println(mathPublic.A, mathPublic.B, mathPublic.Add())

	mathPrivate := math.NewMathPrivate(1, 2)
	fmt.Println(mathPrivate.Add()) // a and b are not accessible

	mathMixed := math.MathMixed{A: 1}         // b is not accessible
	mathMixed.SetB(2)                         // b is set
	fmt.Println(mathMixed.A, mathMixed.Add()) // getB() is not accessible
}
