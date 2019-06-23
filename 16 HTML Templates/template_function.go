package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func main() {

	templateText := "Before loop: {{.}}\n{{range.}} In Loop: {{.}}\n{{end}}After Loop: {{.}}\n"
	executeTemplate("Dot is :{{.}}\n", "Sarathy")
	executeTemplate("Dot is :{{.}}\n", 24242)
	executeTemplate("Dot is :{{.}}\n", true)
	executeTemplate("Start {{if.}} Dot is true {{end}} finish\n", true)
	executeTemplate("Start {{if.}} Dot is true {{end}} finish\n", false)
	executeTemplate(templateText, []string{"fuck", "you", "seriously"})
	templateText = "Prices:\n{{range.}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{5.43, 7.35, 2.5})

	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cable", Count: 355})

	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: $ {{.Rate}}\n{{end}}"
	subscriber := Subscriber{
		Name:   "Partha Sarathy",
		Rate:   45.33,
		Active: true,
	}

	executeTemplate(templateText, subscriber)
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
