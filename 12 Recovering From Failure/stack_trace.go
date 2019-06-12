package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("Deffered in one()")
	two()
}

func two() {
	defer fmt.Println("Deffered in two()")
	three()
}

func three() {
	panic("I'm panicked")
}
