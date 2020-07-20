package main

import (
	"fmt"
)

type person struct { //syntax for a struct
	name   string //fieldName type
	number int
}
type details struct {
	person   //we can use a struct as an underlying field in this way
	company  string
	position string
	number   int
}

func main() {
	p1 := details{
		person: person{
			name:   "chandan",
			number: 111,
		},
		company:  "verizon",
		position: "dev",
	}
	p2 := details{
		person: person{
			name:   "aniket",
			number: 100,
		},
		company:  "tcs",
		position: "test",
		number:   1352762, //We can have a clash with same field name as long as it is present in another struct.
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.name, p1.person.number, p1.company, p1.position)
	fmt.Println(p2.name, p2.person.number, p2.company, p2.position, p2.number)
	// Here number is same field name present in both person and details. If we call number, then it will be shown from details as it is
	// having more scope. The number field from person can be accessed in the above mentioned way : p2.person.number.

}
