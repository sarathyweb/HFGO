package main

import "fmt"

func main() {
	var pet struct {
		name string
		age  int
	}

	pet.name = "Python Programmer"
	pet.age = 55

	fmt.Println("Name: ", pet.name)
	fmt.Println("Age: ", pet.age)
}
