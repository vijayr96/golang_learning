package main

import (
	"fmt"
)

func main() {
	x := sum("james", 1, 2, 3)
	fmt.Println("The total is", x)
}

func sum(s string, x ...int) int { //variadic parameter has to be the final parameter
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
