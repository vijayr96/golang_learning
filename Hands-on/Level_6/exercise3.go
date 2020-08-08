package main

import (
	"fmt"
)

func main() {
	defer franck()
	fmt.Println("This will run before the defer function")
}

func franck() {
	defer func() {
		fmt.Println("This is deferred")
	}()
	fmt.Println("func fanck() is running")
}
