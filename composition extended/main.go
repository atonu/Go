package main

import "fmt"

type Vehicle struct {
	numberOfWheels int
	numberOfSeats  int
}

type Car struct {
	Made        string
	model       string
	isHybrid    bool
	vehicleType Vehicle
}

func (c Car) show() {
	fmt.Println("Made", c.Made)
	fmt.Println("model", c.model)
	fmt.Println("isHybrid", c.isHybrid)
	c.vehicleType.show()
}

func (v Vehicle) show() {
	fmt.Println("No. of seats", v.numberOfWheels)
	fmt.Println("No. of wheels", v.numberOfSeats)
}

func main() {
	sedan := Vehicle{
		numberOfWheels: 4,
		numberOfSeats:  4,
	}

	xcorolla := Car{
		Made:        "Toyota",
		model:       "xCorolla",
		isHybrid:    false,
		vehicleType: sedan,
	}

	xcorolla.vehicleType.numberOfSeats = 5

	xcorolla.show()
	println()

	axio := Car{
		Made:        "Toyota",
		model:       "axio",
		isHybrid:    false,
		vehicleType: sedan,
	}

	axio.show()
}
