package main

import "fmt"

func FuckYou() {
	fmt.Println("Fuck You")
	FuckYou()
}

func main() {
	FuckYou()
}
