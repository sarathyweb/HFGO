package main

import (
	"fmt"
)

func Socialize() {
	defer fmt.Println("Hello")
	defer fmt.Println("How Are You?")
	defer fmt.Println("Good Bye!")
}

func main() {
	Socialize()
}
