package examples

import "fmt"

func Directions() {
	channel := make(chan string)

	go channelSendOnly("content", channel)
	channelReceiveOnly(channel)
}

func channelSendOnly(content string, channel chan<- string) {
	channel <- content
}

func channelReceiveOnly(channel <-chan string) {
	fmt.Println(<-channel)
}
