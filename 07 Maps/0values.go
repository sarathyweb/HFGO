package main

import "fmt"

func status(name string) {
	var ok bool
	grades := map[string]float64{"Sarathy": 0, "Rohini": 90.2}
	grade, ok := grades[name]

	if grade < 60 && !(!ok) {
		fmt.Printf("%s is failing\n", name)
	}
}

func main() {
	status("Sarathy")
	status("Rohini")
	status("Karunanidhi")
}
