package main

import (
	"fmt"
)

func main() {

	switch true {
	case false:
		fmt.Println("this shouldn't print")
	case (1 == 2):
		fmt.Println("this shouldn't print")
	case (2 == 2):
		fmt.Println("this should print 001")
		fallthrough
		//fallthrough keyword transfers the control to next case and executes it whether it's true or false.
		// Without fallthrough the switch-case ends at the first true statement executed.
	case false:
		fmt.Println("this shouldn't print 002")
	case (2 == 2):
		fmt.Println("this should print")
		fallthrough
	default:
		// default case is executed when there is no case matching the switch expression.
		fmt.Println("this is default")
	}
}
