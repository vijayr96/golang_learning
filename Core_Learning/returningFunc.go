package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int { // We are returning an anonymous func of type 'func() int' from func bar()
	return func() int {
		return 999
	}
}
