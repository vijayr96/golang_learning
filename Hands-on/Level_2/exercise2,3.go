package main

import (
	"fmt"
)

const (
	x     = 49
	y int = 45
)

func main() {

	a := (40 == 44)
	b := (41 <= 43)
	c := (44 >= 43)
	d := (41 != 40)
	e := (40 < 36)
	f := (41 > 39)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(x, y)
}
