package math

var A int = 1
var b int = 2

type myInt interface{ ~int }       // can't be accessed from outside the package
type MyFloat interface{ ~float64 } // can be accessed from outside the package

func Sum[T myInt | float64](a, b T) T {
	return a + b
}

func sub[T myInt | float64](a, b T) T {
	return a - b
}

func Example() {
	print(sub(A, b))
}

type Person struct {
	Name string
	age  int // can'n be accessed from outside the struct
}

func (p Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	p.age = age
}

// can'n be accessed from outside the struct
func (p *Person) getName() string {
	return p.Name
}
