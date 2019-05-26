package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(512)
	fmt.Println("Slwwpy Creek County Population", population)
	fmt.Println("Congratulation, Kava in and Anna! iIt's a girl")
	population += 1
	fmt.Println("Slwwpy Creek County Population", population)
}
