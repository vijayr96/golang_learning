package main

import (
	"fmt"
)

var a int

// Below line shows why  GO is all about TYPE.
type sandwich int // creating our own TYPE. Here sandwich is a user TYPE whose underlying TYPE is int.

var b sandwich // Here b is a variable of the type sandwich

func main() {
	a = 40
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 38
	fmt.Println(b)
	fmt.Printf("%T\n", b) // The output for this will be main.sandwich as 38 is of type sandwich.

	// a = b				This won't work as a is int and b is of type sandwich
	// fmt.Println(a)

	a = int(b) //Here we are explicitly converting 'b' from the type sandwich to type int before assigning it to 'a' which is of type int.
	// It wont work without conversion. This is also known as casting in other languages.

	fmt.Println(a)

}
