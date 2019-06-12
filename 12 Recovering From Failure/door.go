package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You won a Cursie!")
	} else if doorNumber == 2 {
		fmt.Println("You won a Car!")
	} else if doorNumber == 3 {
		fmt.Println("You won a goat!")
	} else {
		panic("Invalid Number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
