package main

import (
	"fmt"
)

func main() {

	x := bar()

	fmt.Printf("%T\n", x)

	fmt.Println(x())

	//	fmt.Println(bar()())	//We can write the previous line in this way as well by substituting x with bar()

}

func bar() func() int { // We are returning an anonymous func of type 'func() int' from func bar()
	return func() int {
		return 999
	}
}
