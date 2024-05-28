package services

import (
	"html/template"
	"net/http"
	"os"
)

type course struct {
	Name  string
	Hours int
}

type courses []course

func Run08() {
	ioOutput()
	serverMux()
}

func ioOutput() {
	t := template.Must(template.New("template.html").ParseFiles("./services/public/template.html"))

	coursesGroup := courses{
		{"Go", 20},
		{"Python", 30},
		{"JavaScript", 40},
	}

	err := t.Execute(os.Stdout, coursesGroup)
	if err != nil {
		panic(err)
	}
}

func serverMux() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("./services/public/template.html"))

		coursesGroup := courses{
			{"Go", 20},
			{"Python", 30},
			{"JavaScript", 40},
		}

		err := t.Execute(w, coursesGroup)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
