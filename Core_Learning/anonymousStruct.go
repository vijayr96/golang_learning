package main

import (
	"fmt"
)

func main() {
	p1 := struct { //this is an anonymous struct as we are not defining it with an identifier(name).
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
