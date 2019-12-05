package main

import (
	"fmt"
)

func checkConditions(i int) bool {
	hasSame := false
	previous := -1
	runcount := 1
	for i > 0 {
		// peel off a digit
		digit := i % 10
		i /= 10
		// compare with previous
		if previous != -1 {
			if digit == previous {
				// same digits, update current run
				runcount++
			} else if digit > previous {
				// not descending
				return false
			} else {
				// strictly descending. close current run and start a new one
				if runcount == 2 {
					hasSame = true
				}
				runcount = 1
			}
		}
		previous = digit
	}
	// we may have an open run
	if runcount == 2 {
		hasSame = true
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
	// fmt.Println(checkConditions(112233))
	// fmt.Println(checkConditions(123444))
	// fmt.Println(checkConditions(111122))
	// fmt.Println(checkConditions(669999))
}
