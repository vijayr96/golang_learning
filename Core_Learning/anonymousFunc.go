package main

import (
	"fmt"
)

func main() {
	foo()
	func() {
		fmt.Println("anonymous func ran")
	}()
	func(x int) {
		fmt.Println("also an anonymous func", x)
	}(100)
}

func foo() {
	fmt.Println("func foo ran")
}
