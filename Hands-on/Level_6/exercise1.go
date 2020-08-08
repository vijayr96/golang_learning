package main

import (
	"fmt"
)

func main() {
	a := foo()
	fmt.Println(a)
	x, y := bar()
	fmt.Println(x, y)

}

func foo() int {
	return 119
}

func bar() (int, string) {
	return 111, "Chandan"
}
