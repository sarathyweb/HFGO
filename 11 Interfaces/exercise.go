package main

import "fmt"

type Car string

func (c Car) Acclerate() {
	fmt.Println("Speeding Up")
}

func (c Car) Brake() {
	fmt.Println("Stopping!")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Acclerate() {
	fmt.Println("Speeding Up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping!")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Acclerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Maruthi 800")
	vehicle.Acclerate()
	vehicle.Steer("right")

	vehicle = Truck("Amazon Truck")
	vehicle.Brake()
	vehicle.Steer("left")
}
