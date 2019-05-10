package main

import "fmt"

func main() {
	numbers := make(map[string]int)

	numbers["I have been assigned"] = 12

	fmt.Printf("%#v\n", numbers["I have been assigned"])
	fmt.Printf("%#v\n", numbers["I have not been assigned"])

	words := make(map[string]string)

	words["I have been assigned"] = "Hi"

	fmt.Printf("%#v\n", words["I have been assigned"])
	fmt.Printf("%#v\n", words["I have not been assigned"])

	counters := make(map[string]int)

	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
}
