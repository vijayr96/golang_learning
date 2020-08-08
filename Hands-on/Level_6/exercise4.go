package main

import (
	"fmt"
)

type humanoid struct {
	first string
	last  string
	age   int
}

func (p humanoid) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
	p1 := humanoid{
		first: "James",
		last:  "Hunt",
		age:   24,
	}
	p1.speak()
}
