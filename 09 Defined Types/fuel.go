package main

import "fmt"

type Litters float64
type Milliliters float64
type Gallons float64

func (l Litters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLitters() Litters {
	return Litters(g * 3.785)
}

func (g Gallons) ToMilliLitters() Litters {
	return Litters(g * 3785.41)
}

func main() {
	soda := Litters(2)
	fmt.Printf("%0.3f Litters equals %0.3f Gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f Milliliters equals %0.3f Gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f Gallons equals %0.3f Litters\n", milk, milk.ToLitters())
	fmt.Printf("%0.3f Gallons equals %0.3f MilliLitters\n", milk, milk.ToMilliLitters())
}
