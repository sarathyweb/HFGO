package main

import (
	"fmt"
)

func main() {
	myInt := 4
	var myIntPonter *int
	myIntPonter = &myInt
	fmt.Println(myIntPonter)
	fmt.Println(*myIntPonter)

	myFloat := 9.45
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	var myBoolPointer *bool
	myBoolPointer = &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
	*myBoolPointer = false
	fmt.Println(*myBoolPointer)

}
