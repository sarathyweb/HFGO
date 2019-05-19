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
		Address: magazine.Address{
			Street:     "Pallam Road",
			City:       "Nagercoil",
			State:      "Tamil Nadu",
			PostalCode: "629180",
		},
	}

	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
	fmt.Println("Address", subscriber.Address)

	employee := magazine.Employee{
		Name:   "Jeff Bezos",
		Salary: 100000,
	}
	fmt.Println("Name:", employee.Name)
	fmt.Println("Salary:", employee.Salary)

}
