package examples

import (
	"fmt"
	"time"
)

func RangeChannels() {
	channel := make(chan string)

	go publish(channel)
	reader(channel)
}

func reader(channel chan string) {
	for message := range channel {
		fmt.Println(message)
	}
}

func publish(channel chan string) {
	for i := 0; i < 3; i++ {
		channel <- fmt.Sprintf("Message %d", i)

		time.Sleep(1 * time.Second)
	}
	close(channel)
}
