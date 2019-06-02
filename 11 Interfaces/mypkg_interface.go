package main

import (
	"./mypkg"
	"fmt"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWitParameter(12.44)
	fmt.Println(value.MethodWithReturnValue())
}
