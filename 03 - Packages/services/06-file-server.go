package services

import (
	"log"
	"net/http"
)

func Run06() {
	// path starts from the calling file
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HOME"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
