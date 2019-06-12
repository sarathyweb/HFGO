package main

import (
	"fmt"
)

func snack() {
	defer fmt.Println("Closing Refrigerator") // 2
	fmt.Println("Opening Refrigerator")       // 1
	panic("Refrigerator is empty")            // 3
}

func main() {
	snack()
}
