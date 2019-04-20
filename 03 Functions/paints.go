package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("A width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("A height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}
func main() {
	var amount, total float64
	var err error
	amount, err = paintNeeded(2.3, 4.5)
	fmt.Printf("%.2f liters paint needed\n", amount)
	total += amount
	amount, err = paintNeeded(3.5, 7.44)
	fmt.Printf("%.2f liters paint needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f liters paint needed\n", total)
	amount, err = paintNeeded(6.7, -9.5)
	fmt.Println(err)
	fmt.Printf("%.2f liters paint needed\n", amount)
}
