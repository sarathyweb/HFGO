package main

import (
	"./prose"
	"fmt"
)

func main() {
	phrases := []string{
		"my parents",
		"a rodeo clown",
	}
	fmt.Println("A photo with", prose.JoinWWithCommas(phrases))
	phrases = []string{
		"my parents",
		"a rodeo clown",
		"a prize ball",
	}
	fmt.Println("A photo with", prose.JoinWWithCommas(phrases))
}
