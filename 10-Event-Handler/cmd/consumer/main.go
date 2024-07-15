package main

import (
	"fmt"

	"github.com/KelpGF/Go-Event-Handling/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	forever := make(chan bool)

	out := make(chan amqp.Delivery)
	go startWorkers(out)
	go rabbitmq.Consumer(ch, "orders", out)

	<-forever
}

func startWorkers(channel <-chan amqp.Delivery) {
	workersQtd := 10

	for i := 1; i <= workersQtd; i++ {
		go worker(i, channel)
	}
}

func worker(id int, channel <-chan amqp.Delivery) {
	for message := range channel {
		fmt.Printf("Worker %d received: %s\n", id, string(message.Body))
		message.Ack(false)
	}
}
