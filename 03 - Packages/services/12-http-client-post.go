package services

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func Run12() {
	c := http.Client{}

	jsonBody := bytes.NewBuffer([]byte(`{"name": "Jelps", "age": 21}`))

	res, err := c.Post("https://google.com", "application/json", jsonBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
