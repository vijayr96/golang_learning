package main

import (
	"fmt"
)

// Use var keyword when declaring a variable outside function scope. It's scope is package level. So, try to limit it's
// usage as you shouldn't have all variables in global scope.

var y = 50

func main() {

	// Declare a variable and assigning value to it directly using a short declaration operator ' := '

	x := 40
	fmt.Println(x)

	fmt.Println(y)
}
