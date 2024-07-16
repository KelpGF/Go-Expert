package main

import (
	"fmt"

	"github.com/KelpGF/fc-utils/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
