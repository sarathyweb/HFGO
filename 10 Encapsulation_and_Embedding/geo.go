package main

import (
	"./geo"
	"fmt"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-1222.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longititude())
}
