package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("sum of all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("sum of even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("sum of odd numbers", s3)
}

func sum(xi ...int) int {

	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f1 func(xi ...int) int, vi ...int) int {
	var zi []int
	for _, v := range vi {
		if v%2 == 1 {
			zi = append(zi, v)
		}
	}
	return f1(zi...)
}

// Here even and odd funcs are using callback to func sum as their parameters f and f1 respectively. Here, even and
// odd are getting unfurled slice and creating even and odd slices respectively. Then those slices are calling back sum() function
// to find the sum total and return it to the calling function in main.
