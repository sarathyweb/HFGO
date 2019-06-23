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
	var text = "Here's my template\n"
	tmpl, err := template.New("text").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}
