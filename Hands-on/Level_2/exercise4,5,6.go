package main

import (
	"fmt"
)

const (
	m = 2020 + iota
	n = 2020 + iota
	o = 2020 + iota
	p = 2020 + iota
)

func main() {

	a := 10
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
	fmt.Println(m, n, o, p)

}
