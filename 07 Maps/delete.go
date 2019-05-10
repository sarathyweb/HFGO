package main

import "fmt"

func main() {
	var k bool
	ranks := make(map[string]int)

	ranks["bronze"] = 3
	rank, k := ranks["bronze"]
	fmt.Println(rank, k)
	delete(ranks, "bronze")
	rank, k = ranks["bronze"]
	fmt.Println(rank, k)
}
