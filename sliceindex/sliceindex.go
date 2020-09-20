package main

import (
	"fmt"
)

func main() {

	var countries = []string{"Indonesia", "Malaysia", "Singapore", "Thailand", "Vietnam"}
	fmt.Println("Neighboring countries of Indonesia are", countries[1:3])

}
