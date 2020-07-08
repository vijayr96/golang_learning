package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	foo()

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Printing only even numbers")
		}
	}
}

func foo() {
	fmt.Println("Im now in the function foo and calling it from outside of main")
}
