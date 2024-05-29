package services

import (
	"html/template"
	"net/http"
	"strings"
)

func Run10() {
	templatesList := []string{
		"./public/header.html",
		"./public/content-upper.html",
		"./public/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content-upper.html")
		t.Funcs(template.FuncMap{
			"ToUpper": strings.ToUpper,
		})

		t = template.Must(t.ParseFiles(templatesList...))

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
