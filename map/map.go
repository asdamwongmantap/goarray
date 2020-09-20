package main

import (
	"fmt"
)

func main() {

	var population = map[string]int{"Indonesia": 267000000, "Malaysia": 31500000, "Singapore": 5639000}

	for key, value := range population {
		fmt.Println(key, "Country", "has", value, "population")
	}

}
