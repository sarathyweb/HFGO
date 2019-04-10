package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 33
	length, width = 12.3, 33.222
	customerName = "Partha Sarathy"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with area of ", width*length, "square meters")
}
