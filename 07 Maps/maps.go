package main

import "fmt"

func main() {
	ranks := make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["brnze"] = 3
	fmt.Println(ranks["silver"])
}
