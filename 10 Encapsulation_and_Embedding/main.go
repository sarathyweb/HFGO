package main

import (
	"./calender"
	"fmt"
	"log"
)

func main() {
	event := calender.Event{}
	err := event.SetTitle("Partha's Birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(1998)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(06)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

}
