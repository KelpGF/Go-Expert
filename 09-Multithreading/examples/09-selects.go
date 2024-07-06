package examples

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID      int64
	Content string
}

func Selects() {
	var id int64 = 0

	rabbitMqChannel := make(chan Message)
	kafkaChannel := make(chan Message)

	go func() {
		for {
			time.Sleep(3 * time.Second)

			msgId := atomic.AddInt64(&id, 1)
			msg := Message{ID: msgId, Content: "Message from RabbitMQ"}
			rabbitMqChannel <- msg
		}
	}()

	go func() {
		for {
			msgId := atomic.AddInt64(&id, 1)
			msg := Message{ID: msgId, Content: "Message from Kafka"}
			kafkaChannel <- msg

			time.Sleep(2 * time.Second)
		}
	}()

	// for i := 0; i < 4; i++ {
	for {
		// select waits for the first channel to be ready
		select {
		case msg1 := <-rabbitMqChannel:
			fmt.Printf("RabbitMQ received %d with content: %s\n", msg1.ID, msg1.Content)
		case msg2 := <-kafkaChannel:
			fmt.Printf("Kafka received %d with content: %s\n", msg2.ID, msg2.Content)
		case <-time.After(1 * time.Second):
			println("Timeout. No message received")
			// default:
			// 	println("no one was ready")
		}

		time.Sleep(500 * time.Millisecond)
	}
}
