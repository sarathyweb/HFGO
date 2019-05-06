package main

import "fmt"

func main() {
	var notes []string

	notes = make([]string, 7)
	notes[0] = "Sarathy"
	notes[1] = "Rohini"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	fmt.Println(primes[0:3])

	lettrs := []string{"a", "e", "i", "o", "u"}
	for i := 0; i < len(lettrs); i++ {
		fmt.Println(lettrs[i])
	}

	for _, i := range lettrs {
		fmt.Println(i)
	}

	underlyingArray := [5]string{"a", "e", "i", "o", "u"}
	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)

}
