package main

import (
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	print(year)
	print("\n")
}

// Methods are functions that are associated with the values of particular type
