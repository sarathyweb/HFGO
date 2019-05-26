package main

import "fmt"

type Title string
type Litters float64
type Gallons float64

func main() {
	fmt.Println(Litters(1.2) + Litters(Gallons(9.8)))
	fmt.Println(Gallons(Litters(78.0)) - Gallons(23.4))
	fmt.Println(Gallons(10.11) * Gallons(9.3))
	fmt.Println(Litters(10000) / Litters(Gallons(332)))
	fmt.Println(Gallons(99.33) == Gallons(99.33))
	fmt.Println(Litters(9.3) < Litters(8.3))
	fmt.Println(Litters(9.3) > Litters(8.3))

	fmt.Println(Title("Hi") + Title("How"))
	//fmt.Println(Title("Hi") - Title("How")) - Invalid operation
	fmt.Println(Title("Hi") == Title("How"))
	fmt.Println(Title("Hi") < Title("How"))

	fmt.Println(Gallons(100.33) + 00.4)
	fmt.Println(Litters(32.4) - 9.4)
	fmt.Println(Title("Fuck ") + "You!")
}
