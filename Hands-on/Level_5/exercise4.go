package main

import (
	"fmt"
)

func main() {
	anonStruct := struct {
		name         string
		favSports    []string
		sportsRating map[string]int
	}{
		name: "Chandan",
		favSports: []string{
			"cricket",
			"marathon",
			"football",
		},
		sportsRating: map[string]int{
			"cricket":  8,
			"marathon": 10,
			"football": 4,
		},
	}

	fmt.Println(anonStruct)
	fmt.Println("xxxxxxxxxxxxxxxxx")
	fmt.Println(anonStruct.name)
	for i, v := range anonStruct.favSports {
		fmt.Println(i+1, v)
	}
	for k, v := range anonStruct.sportsRating {
		fmt.Println("For", k, "the rating is", v)
	}

}
