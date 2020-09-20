package main

import (
	"fmt"
)

func main() {

	var countries = []string{"Indonesia", "Malaysia", "Singapore", "Thailand", "Vietnam"}
	for i := 0; i < len(countries); i++ {
		fmt.Println("Index to", i, countries[i])
	}

}
