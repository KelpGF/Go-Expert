package hotel

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	ctx := context.Background()
	// ctx, cancel = context.WithCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking canceled! Time's up")
		return
	case <-time.After(time.Second * 1):
		fmt.Println("Hotel booked!")
	}
}
