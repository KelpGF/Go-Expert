package main

var b bool

var (
	c int
	d = 1
	e string
	f float64
)

func main() {
	y := 42
	println(y)

	y = 43
	println(y)
}

func test() {
	// println(y) // Error: undefined: y
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
