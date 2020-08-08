package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	sq := square{
		length: 10.5,
	}

	circ := circle{
		radius: 5,
	}
	info(sq)
	info(circ)
}
