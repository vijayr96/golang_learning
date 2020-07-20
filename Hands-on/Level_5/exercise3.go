package main

import (
	"fmt"
)

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

	leyland := truck{
		vehicle: vehicle{
			doors: 2,
			color: "maroon",
		},
		fourWheel: true,
	}

	suzuki := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: false,
	}

	fmt.Println(leyland)
	fmt.Println(suzuki)
	fmt.Println(leyland.doors, leyland.color)
	fmt.Println(suzuki.doors, suzuki.color)

}
