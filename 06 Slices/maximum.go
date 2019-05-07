package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(maximum(1, 2, 3, 4, 5, 6, 7, 8, 5, 3, 5, 2, 5, 6, 11, 35, 52, 5, 2, 98, 200))
	fmt.Println(inRange(1, 5, 3, 4, 2, 4, 21, 42, 5, 1))
}
