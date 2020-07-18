package main

import (
	"fmt"
)

func main() {
	//The format for a map is 'map[key]value'. This acts as the type for the map,
	//which in this case is map[string]int
	m := map[string]int{
		"Chandan": 111, //The ',' is compulsory at the end of all lines of key-value pairs.
		"Vijay":   119,
		"Aniket":  101,
		"Karthik": 100,
	}
	fmt.Println("MAP ---")
	fmt.Println(m)
	fmt.Println(m["Chandan"])

	fmt.Println(m["Alex"]) //This key is not present in the map. So it returns 0.
	//But in a large map we might confuse it for having value of 0. So we can check whether the key-value pair is present using
	// the below method.
	fmt.Println(" ")
	fmt.Println("Checking wheteher key-value is in Map ----")
	v, ok := m["Alex"]
	fmt.Println(ok)
	fmt.Println(v)
	if v, ok := m["Vijay"]; ok {
		fmt.Println("this is the value if the key is present in the map", v)
	}
	// Used if statement to print the value of the key only if its present in the map. 'if simpleStatement;expression{ }'
	fmt.Println(" ")
	fmt.Println("Adding a key to Map ---")
	m["Jo"] = 99 //to add a key-value to a map

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(" ")
	fmt.Println("Deleting a key from Map ---")
	delete(m, "Jo") //deleting a key-value from map.
	fmt.Println(m)

	fmt.Println(" ")
	fmt.Println("Checking whether key is present before Deleting ---")
	//To check whether the key is present, before deleting it and to avoid any false deletions, we can use below method.
	if v, ok := m["Vijay"]; ok {
		fmt.Println("value:", v)
		delete(m, "Vijay")
	}
	fmt.Println(m)

}
