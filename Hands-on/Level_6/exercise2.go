package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := ava(xi...)
	fmt.Println(x)

	yi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := max(yi)
	fmt.Println(y)
}

func ava(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func max(yi []int) int {
	total := 0
	for _, v := range yi {
		total += v
	}
	return total
}
