package main

import (
	"fmt"
)

// Use var keyword when declaring a variable outside function scope. It's scope is package level. So, try to limit it's
// usage as you shouldn't have all variables in global scope.

var y = 50

// GO is a static prog language. A variable is declared to hold a value of a certain TYPE which
// can't be changed unlike a dynamic prog language. Here, its of TYPE int.

var z string // declaring a variable of type string and assigning Zero/nil value to it, so that the variable can be used
// later in different parts of the code. Can be done with float, booleans, strings, pointers, functions, interfaces,
// channels, slices and maps as well.

var a string = ` I said, "both types of quotes are possible for a string" `

// The ' ' quote is a raw string literal

func main() {

	// Declare a variable and assigning value to it directly using a short declaration operator ' := '

	x := 40
	fmt.Println(x)
	fmt.Printf("%T\n", x) // %T gives the TYPE of the variable.

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//z = 30
	// Because z is declared as string, we can't use it as int again unlike in dynamic prog languages. The above line will throw an error.

	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
