package main

import "fmt"

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}
func main() {
	paintNeeded(2.3, 4.5)
	paintNeeded(3.5, 7.44)
	paintNeeded(6.44, 67.3)

}
