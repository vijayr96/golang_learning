package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	sliceOfSlice := [][]string{x, y}
	fmt.Println(sliceOfSlice)

	for i, record := range sliceOfSlice {
		fmt.Println("record :", i)
		for j, value := range record {
			fmt.Println("At index: ", j, "the value is: ", value)
		}
	}
}
