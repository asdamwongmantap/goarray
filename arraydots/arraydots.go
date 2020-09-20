package main

import (
	"fmt"
)

func main() {

	var city = [...]string{"Cilegon", "Kuala Lumpur", "Queenstown", "Bangkok"}
	for i := 0; i < len(city); i++ {
		fmt.Println("Index to", i, city[i])
	}

}
