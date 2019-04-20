package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Printf("Enter a Temperature in Fahrenheit:")

	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degree celsius\n", celsius)
}
