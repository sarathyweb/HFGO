package main

import "fmt"

func main() {
	notes := [3]string{"sarathy", "abc", "xyz"}
	primes := [3]int{2, 3, 5}
	fmt.Println(notes)
	fmt.Println(primes)
	fmt.Printf("%v\n", primes)
	fmt.Printf("%v\n", primes)
	fmt.Printf("%#v\n", primes)
}
