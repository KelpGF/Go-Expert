package services

import (
	"io"
	"net/http"
	"time"
)

func Run11() {
	c := http.Client{
		Timeout: time.Millisecond, // Get "https://google.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	}

	res, err := c.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
