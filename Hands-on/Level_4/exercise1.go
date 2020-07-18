package main

import (
	"fmt"
)

func main() {
	x := [5]int{36, 34, 35, 48, 71}
	for i, v := range x {
		fmt.Println("At index", i, "the value is", v)
	}
	fmt.Printf("%T\n", x)

}
