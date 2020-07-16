package main

import (
	"fmt"
)

func main() {

	x := 100
	if x == 90 {
		fmt.Println("value is 90")
	} else if x == 100 {
		fmt.Println("value is 100")
	} else {
		fmt.Println(x)
	}
}
