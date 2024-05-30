package http

import (
	"log"
	"net/http"
	"time"
)

func RunServer() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	log.Println("Started handling request")
	defer log.Println("Finished handling request")

	select {
	case <-time.After(time.Second * 5):
		log.Println(">> Request processed")
		w.Write([]byte("Request processed"))

	case <-requestContext.Done():
		err := requestContext.Err()
		log.Println(">> Request canceled by Client:", err.Error())
		http.Error(w, "Request canceled by Client", http.StatusRequestTimeout)
	}
}
