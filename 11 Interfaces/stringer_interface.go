package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + "Coffee Pot"
}

func main() {
	coffeePot := CoffeePot("Brew")
	fmt.Println(coffeePot.String())
}
