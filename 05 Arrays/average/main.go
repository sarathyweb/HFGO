package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloat("data.txt")
	if err != nil {
		log.Fatal()
	}
	var sum float64 = 0
	var avg float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	avg = sum / sampleCount
	fmt.Println(avg)
}
