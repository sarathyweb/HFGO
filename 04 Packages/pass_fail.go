package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Printf("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
