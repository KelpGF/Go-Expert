package examples

import (
	"fmt"
	"sync"
	"time"
)

func RangeChannelsWG() {
	channel := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(3)

	go publishWg(channel)
	go readerWg(channel, &wg)

	wg.Wait()
}

func readerWg(channel chan string, wg *sync.WaitGroup) {
	for message := range channel {
		fmt.Println(message)
		wg.Done()
	}
}

func publishWg(channel chan string) {
	for i := 0; i < 3; i++ {
		channel <- fmt.Sprintf("Message %d", i)

		time.Sleep(1 * time.Second)
	}
	close(channel)
}
