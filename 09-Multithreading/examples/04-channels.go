package examples

import "time"

func Channels() {
	channel := make(chan string)

	go func() {
		channel <- "Hello"
		channel <- "Hello 2"

		time.Sleep(2 * time.Second)
		channel <- "Hello 3"
		channel <- "Hello 4" // This will not be printed
	}()

	message := <-channel
	println(message) // Hello

	message = <-channel
	println(message) // Hello 2

	message = <-channel
	println(message) // waiting for message
}

// Used to keep the main thread alive when the goroutine is still running
func ChannelsForever() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			println(i)
			time.Sleep(1 * time.Second)
		}

		forever <- true
	}()

	<-forever
}
