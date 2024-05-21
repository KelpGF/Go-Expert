package services

import (
	"io"
	"net/http"
)

func Run02() {
	httpRequest()
}

func httpRequest() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// defer is used to do something after the function ends
	defer req.Body.Close()

	// req is a stream. We need to read all of it.
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
}
