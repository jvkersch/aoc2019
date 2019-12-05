package main

import (
	"fmt"
)

func checkConditions(i int) bool {
	hasSame := false
	previous := -1
	for i > 0 {
		// peel off a digit
		digit := i % 10
		i /= 10
		// compare with previous
		if previous != -1 {
			if digit == previous {
				hasSame = true
			}
			if digit > previous {
				// not descending
				return false
			}
		}
		previous = digit
	}
	return hasSame
}

func main() {
	start, stop := 284639, 748759
	for i := start; i <= stop; i++ {
		if checkConditions(i) {
			fmt.Println(i)
		}
	}
}
