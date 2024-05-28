package services

import (
	"html/template"
	"os"
)

type person07 struct {
	Name string
	Age  int
}

func Run07() {
	default07()
	must07()
}

func default07() {
	person := person07{Name: "Jelps", Age: 21}

	t := template.New("PersonTemplate")
	t, _ = t.Parse("Name: {{.Name}}\nAge: {{.Age}}")

	err := t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}

func must07() {
	person := person07{Name: "Kelv", Age: 21}

	t := template.Must(template.New("PersonTemplate").Parse("Name2: {{.Name}}\nAge2: {{.Age}}"))
	err := t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}
