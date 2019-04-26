package main

import "fmt"

func main() {
	notes := [7]string{"Venture", "Capitalist", "Startup", "Fund", "Cloud", "Business", "Failures"}
	fmt.Println("Lenth of the array is", len(notes))
	/*	index := 1
		fmt.Println(index, notes[index])
		index = 3
		fmt.Println(index, notes[index])
	*/
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	for index, note := range notes {
		fmt.Println(index, note)
	}
}
