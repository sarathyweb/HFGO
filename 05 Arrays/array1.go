package main

import "fmt"

func main() {
	var myInt [5]int = [5]int{1, 2, 1, 2, 3}
	var notes [7]string = [7]string{"Sarathy", "Hello", "How", "Are", "You", "Hi", "How"}
	text := [3]string{
		"This is a paragraph",
		"This a seconf paragraphp",
		"This is a third one",
	}

	fmt.Println(notes[6], notes[2], notes[4])
	fmt.Println(myInt[2])
	fmt.Println(text[2])
}
