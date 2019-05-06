package main

import "fmt"

func main() {
	array1 := [5]string{"a", "e", "i", "o", "u"}
	slice1 := array1[0:3]
	array1[1] = "x"
	fmt.Println(array1)
	fmt.Println(slice1)
}
