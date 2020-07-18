package main

import (
	"fmt"
)

func main() {
	x := []int{36, 34, 35, 48, 71, 88, 37, 43, 74, 53}
	for i, v := range x {
		fmt.Println("At index", i, "the value is", v)
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
