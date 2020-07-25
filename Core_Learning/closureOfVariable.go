package main

import (
	"fmt"
)

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("")
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println("")
	fmt.Println(a())
	fmt.Println(a())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// This code shows that variables have certain enclosures. Here a is incremented 4 times and then once b is used the
// scope is changed. b is incremented 3 times and then a is called again. Now a returns with the previous scope where a was 4 and
// now incremented 2 times.
