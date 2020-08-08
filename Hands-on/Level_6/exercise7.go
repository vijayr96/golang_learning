package main

import (
	"fmt"
)

var a int
var b func()

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = f
	b()
	fmt.Printf("%T\n", b)

}
