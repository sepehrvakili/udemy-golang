package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range numbers {
		var label string
		if num%2 == 0 {
			label = "even"
		} else {
			label = "odd"
		}
		fmt.Println(i, "is "+label)
	}
}
