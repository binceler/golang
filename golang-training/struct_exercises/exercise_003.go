package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	fmt.Println(t, t.vehicle, t.vehicle.color, t.vehicle.doors, t.fourWheel)

	s := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(s, s.vehicle, s.vehicle.color, s.vehicle.doors, s.luxury)

}
