package main

import "fmt"

type Number int

func (n *Number) Display() {
	fmt.Println(*n)
}

func (n *Number) Double() {
	*n *= 2
}

func main() {
	number := Number(4)
	number.Double()
	number.Display()
}
