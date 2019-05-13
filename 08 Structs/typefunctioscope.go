package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name: ", porsche.name)
	fmt.Println("Top Speed: ", porsche.topSpeed)

	var bolts part
	bolts.count = 24
	bolts.description = "Hex bolts"
	fmt.Println("Description: ", bolts.description)
	fmt.Println("Count: ", bolts.count)

}
