package examples

func Buffers() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
