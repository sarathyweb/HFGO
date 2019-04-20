package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("Height can't be negative")
	fmt.Println(err)
	fmt.Printf("Type of err is %T\n", err)
	err1 := fmt.Errorf("Fuck you %.2f\n", 2.4432432422)
	fmt.Println(err1)
	fmt.Println(err1.Error())
	log.Fatal(err)
}
