package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstName: "chandan",
		lastName:  "kumar",
		favIceCream: []string{
			"Chocolate",
			"pan",
			"rajbhog",
		},
	}
	p2 := person{
		firstName: "karthik",
		lastName:  "narayana",
		favIceCream: []string{
			"vanilla",
			"hazelnut",
			"bubblegum",
		},
	}
	fmt.Println(p1)
	for i, v := range p1.favIceCream {
		fmt.Println(i+1, v)
	}
	fmt.Println(p2)
	for i, v := range p2.favIceCream {
		fmt.Println(i+1, v)
	}
}
