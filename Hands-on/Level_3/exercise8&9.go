package main

import (
	"fmt"
)

func main() {
	favSport := "football"
	switch favSport {
	case "cricket":
		fmt.Println("find a cricket ground")
	case "badminton":
		fmt.Println("get a racket")
	case "tennis":
		fmt.Println("find a tennis court")
	case "football":
		fmt.Println("get a playing team of 11")
	case "swimming":
		fmt.Println("find a pool")

	default:
		fmt.Println("sit at home and chill")
	}
}
