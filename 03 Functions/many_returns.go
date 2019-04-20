package main

import "fmt"

func manyReturns() (int, bool, string) {
	return 1, true, "fuck"
}

func main() {
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
}
