package main

import (
	"fmt"
)

type personnel struct {
	first string
	last  string
}

type secretAgent struct {
	personnel
	ltk bool
}

func main() {
	sa1 := secretAgent{
		personnel: personnel{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}

// func (r receiver) identifier(parameter(s)) (return(s)) {code}

// When you have a receiver, (in this case its secretAgent) it is going to
// attach the function to that type. So any value of that type will have access to this function.
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
