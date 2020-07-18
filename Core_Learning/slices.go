package main

import (
	"fmt"
)

func main() {
	// x:= type{values}   It is a composite literal.
	x := []int{3, 5, 7, 11, 13} //This is a slice. It is built on top of an array.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(" ")

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(" ")

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
	fmt.Println(" ")

	y := []int{17, 19, 23, 29}
	x = append(x, y...)
	// The 3 dots ... are needed at the end as it is the format for appending a slice to another
	// slice. It denotes that any number of variadic parameters of the same type can be appended at the end.

	fmt.Println(x)
	fmt.Println(" ")

	x = append(x[:3], x[4:]...)
	fmt.Println(x)
	fmt.Println(" ")

	z := make([]int, 10, 11)
	// make is an inbuilt function used to create a slice when a higher capacity one is needed which
	// can be filled later. This reduces the processing power needed to increase the size of the slice.
	// make(type, length, capacity) is the format. Here, length is the length of slice to be created.
	// Capacity is the total capacity of the underlying array.

	z[3] = 37
	z[8] = 31
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	fmt.Println(" ")

	z = append(z, 333)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	fmt.Println(" ")

	z = append(z, 391)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	fmt.Println(" ")
	// Here, as the capacity of underlying array was full, the runtime created another array of doubled size and
	// deleted the current underlying array in order to fit the appended element.

	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Money", "Penny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	fmt.Println(" ")

	jbxmp := [][]string{jb, mp} // A slice of slice (multi-dimensional slice) is written this way. A composite slice holds smaller slices.
	fmt.Println(jbxmp)

}
