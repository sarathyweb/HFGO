package main

import "fmt"

type Litters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Litters

	carFuel = Gallons(10.0)
	busFuel = Litters(240.0)

	fmt.Println(carFuel, busFuel)

	carFuel = Gallons(Litters(40.0) * 0.264)
	busFuel = Litters(Gallons(63.0) * 3.785)

	fmt.Printf("Gallons : %0.1f Litters: %0.1f\n", carFuel, busFuel)
}
