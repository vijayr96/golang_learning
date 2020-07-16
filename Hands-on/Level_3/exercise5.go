package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		a := i % 4
		fmt.Printf("The remainder of %d mod 4 is %d \n", i, a)
	}
}
