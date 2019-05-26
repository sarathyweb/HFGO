package main

import "fmt"

type MyType string

func (m MyType) SayHi() {
	fmt.Println("Hi from", m)
}

func main() {
	value := MyType("A MyType value")
	value.SayHi()
	anathorvalue := MyType("Anathor value")
	anathorvalue.SayHi()
}
