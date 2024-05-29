package services

import (
	"io"
	"net/http"
	"os"
)

func Run13() {
	c := http.Client{}

	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
