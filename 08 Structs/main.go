package main

import (
	"./magazine"
	"fmt"
)

func main() {
	subscriber := magazine.Subscriber{
		Name:   "Sunadr Pichai",
		Rate:   4.99,
		Active: true,
	}

	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
}
