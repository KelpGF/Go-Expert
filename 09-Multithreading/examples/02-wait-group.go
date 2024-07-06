package examples

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2) // set the number of "credits" to wait

	go taskWG("goroutine 1", &waitGroup)
	go taskWG("goroutine 2", &waitGroup)

	waitGroup.Wait()
}

func taskWG(name string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Task %s: %d\n", name, i)
		time.Sleep(time.Second)

		wg.Done() // decrement the number of "credits"
	}
}
