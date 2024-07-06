package examples

import (
	"fmt"
)

func LoadBalancer() {
	channel := make(chan int)

	startWorkers(channel)
	producer(channel)
}

func startWorkers(channel <-chan int) {
	workersQtd := 10

	for i := 1; i <= workersQtd; i++ {
		go worker(i, channel)
	}
}

func worker(id int, channel <-chan int) {
	for data := range channel {
		fmt.Printf("Worker %d received: %d\n", id, data)
	}
}

func producer(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
}
