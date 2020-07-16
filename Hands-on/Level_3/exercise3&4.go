package main

import (
	"fmt"
)

func main() {
	bday := 1996
	age := 0
	for bday <= 2020 {
		fmt.Println(bday)
		if bday == 2020 {
			fmt.Println("Age is :", age)
		}
		bday++
		age++
	}

	// for {						//another way of for, without expression.
	//	if bday > 2020 {
	//		break
	//	}
	//	fmt.Println(bday)
	//	bday++
	// }
}
