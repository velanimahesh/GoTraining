// Embedding struct

package main

import "fmt"

type vehicle struct {
	color string
	door  string
}

type Car struct {
	vehicle
	fourWheel bool
}

type Bike struct {
	vehicle
	fourWheel bool
}

func main() {

	// BIKE
	v1 := Bike{
		vehicle: vehicle{
			color: "Black",
			door:  "NO_DOOR",
		},
		fourWheel: false,
	}

	// CAR
	v2 := Car{
		vehicle: vehicle{
			color: "RED",
			door:  "AUTOMATIC",
		},
		fourWheel: true,
	}

	fmt.Println("Bike Conf = ", v1)
	fmt.Println("CAR conf = ", v2)

	fmt.Println("Bike Color = ", v1.color)
	fmt.Println("CAR Color = ", v2.color)

}
