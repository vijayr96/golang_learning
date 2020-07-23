package main

import (
	"fmt"
)

func main() {
	x := sum(2, 3, 4, 5, 6, 7)
	fmt.Println("The total is", x)
}

// Below is the function definition.
func sum(a ...int) int { // Here, multiple integers can be passed as ... denotes a variadic parameter.
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	total := 0
	for i, v := range a {
		total += v
		fmt.Println("At index", i, "we are adding", v, "to the sum which is", total)
	}
	return total

}

// func (r receiver) identifier(parameter(s)) (return(s)) {code}
