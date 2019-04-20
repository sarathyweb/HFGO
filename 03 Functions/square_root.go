package main

import (
	"fmt"
	"log"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("Can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squareRoot(-9.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.3f", root)
}
