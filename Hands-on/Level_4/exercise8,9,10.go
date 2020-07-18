package main

import (
	"fmt"
)

func main() {
	mainMap := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	mainMap["fleming_ian"] = []string{`steaks`, `cigars`, `novels`}
	delete(mainMap, `no_dr`)

	for key, value := range mainMap {
		fmt.Println("For ", key, "the favourites are : ")
		for sliceIndex, sliceValue := range value {
			fmt.Println(sliceIndex+1, sliceValue)
		}
		fmt.Println(" ")
	}
}
