package services

import "net/http"

func Run05() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/home", homeHandler)
	mux.Handle("/blog", blog{title: "Jelps"})

	go http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world! MUX 2"))
	})

	http.ListenAndServe(":8081", mux2)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME"))
}

type blog struct {
	title string
}

// blog implements the http.Handler interface
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog " + b.title))
}
