package main

import "fmt"

func status(grade float64) string {
	if grade >= 60 {
		return ("Passing")
	} else {
		return ("Failing")
	}
}

func main() {
	fmt.Println(status(59.0))
	fmt.Println(status(60))
}
