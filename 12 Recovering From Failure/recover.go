package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
}

func freakOut() {
	defer calmDown()
	panic("oh, no")
	fmt.Println("I wont run")
}

func main() {
	freakOut()
	fmt.Println("Exiting Normally!")
}
