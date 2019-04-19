package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Print("false")
	}
	if !false {
		fmt.Println("!false")
	}
	if true {
		fmt.Println("if true")
	} else {
		fmt.Println("else")
	}
	if false {
		fmt.Println("false")
	} else if true {
		fmt.Println("if else true")
	}
	if 12 == 12 {
		fmt.Println("12 == 12")
	}
	if 12 != 12 {
		fmt.Println("12 != 12")
	}
	if 12 > 12 {
		fmt.Println("12 > 12")
	}
	if 12 >= 12 {
		fmt.Println("12 >= 12")
	}
	if 12 == 12 && 6.3 == 3.4 {
		fmt.Println("12 == 12 && 6.3 == 3.4")
	}
	if 12 == 12 || 6.3 == 3.4 {
		fmt.Println("12 == 12 || 6.3 == 3.4")
	}
}
