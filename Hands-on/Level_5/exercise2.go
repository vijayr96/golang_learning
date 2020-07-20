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

	mapped := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	for k, v := range mapped {
		fmt.Println(v.firstName)
		fmt.Println(k)
		for i, icecream := range v.favIceCream {
			fmt.Println(i+1, icecream)
		}
		fmt.Println("xxxxxxxxxxxxxx")
	}
}
