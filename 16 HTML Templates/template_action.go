package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Template Start\nAction: {{.}}\nTemplate End\n"
	tmpl, err := template.New("text").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, "Hello")
}
