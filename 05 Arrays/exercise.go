package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 43
	numbers[2] = 108
	var letter = [3]string{"a", "b", "c"}
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letter[2])
	fmt.Println(letter[1])
	fmt.Println(letter[0])
}
