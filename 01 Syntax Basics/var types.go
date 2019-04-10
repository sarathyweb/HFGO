package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Type of 32 is", reflect.TypeOf(32))
	fmt.Println("Type of 3.2 is", reflect.TypeOf(3.2))
	fmt.Println("Type of True is", reflect.TypeOf(true))
	fmt.Println("Type of string is", reflect.TypeOf("string"))
}
