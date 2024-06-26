package services

import (
	"html/template"
	"net/http"
)

type TemplateContent struct {
	Courses courses
	Title   string
}

func Run09() {
	templatesList := []string{
		"./public/header.html",
		"./public/content.html",
		"./public/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templatesList...))

		coursesGroup := courses{
			{"Go", 20},
			{"Python", 30},
			{"JavaScript", 40},
		}

		templateContent := TemplateContent{
			Courses: coursesGroup,
			Title:   "Jelps",
		}

		err := t.Execute(w, templateContent)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
