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

type human interface {
	speak() // As speak() method is attached to types secretAgent and personnel, both these types are also type human.
}

func main() {
	sa1 := secretAgent{
		personnel: personnel{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := personnel{
		first: "Dr",
		last:  "yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(p1)
	p1.speak()
	bar(sa1)
	bar(p1)
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "-- speak() attached to secretAgent")
}

func (s personnel) speak() {
	fmt.Println("I am", s.first, s.last, "-- speak() attached to personnel")
}

func bar(h human) {
	switch h.(type) {
	case personnel:
		fmt.Println("passed into bar and using type personnel under the interface of type human --", h.(personnel).first)
		// Using h.(personnel).first means that we are asserting that it is of type personnel
	case secretAgent:
		fmt.Println("passed into bar and using type secretAgent under the interface of type human --", h.(secretAgent).first)

	}
	fmt.Println("human interface is passed into bar", h)
}

// If there is an empty interface with no method, then every value can implement that type . All types can implement empty interface
// and anything can be passed into that.
