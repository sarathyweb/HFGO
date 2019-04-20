package main

import "fmt"

func createPonter() *float64 {
	var myFloat = 98.3
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	var myBool bool = true
	var myFloatPointer *float64 = createPonter()
	fmt.Println(*myFloatPointer)
	printPointer(&myBool)
}
