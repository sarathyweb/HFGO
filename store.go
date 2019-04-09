package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "$")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is ", tax, "$")

	var total = float64(price) + tax
	fmt.Println("Total cost is", total, "$")

	var availableFunds int = 120
	fmt.Println(availableFunds, "$ available")
	fmt.Println("Within Budget?", total <= float64(availableFunds))
}
