package examples

import (
	"fmt"
	"time"
)

func GoRoutines() {
	go task("goroutine 1")
	go task("goroutine 2")
	task("goroutine") // run simultaneously with the two goroutines

	// main thread needs a action for don't finish
	time.Sleep(4 * time.Second)
}

func task(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Task %s: %d\n", name, i)
		time.Sleep(time.Second)
	}
}
