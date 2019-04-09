package main

import "fmt"

func main() {
	var length float64 = 98
	var width int = 33
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width:", length > float64(width))
}
