package http

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func RunClient() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// cancel context if request takes more than 3 seconds
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
