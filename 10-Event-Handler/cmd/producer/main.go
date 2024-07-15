package main

import "github.com/KelpGF/Go-Event-Handling/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "amq.direct", "Hello, World!")
}
