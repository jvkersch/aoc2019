package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	// remove trailing line feed
	data = data[0 : len(data)-1]

	lowestCount := 1000000
	lowestLayer := 0
	for i := 0; i < len(data); i += 25 * 6 {
		layer := data[i : i+25*6]
		count := 0
		for _, char := range layer {
			if char == '0' {
				count++
			}
		}
		if count < lowestCount {
			lowestCount = count
			lowestLayer = i
		}
	}

	count1 := 0
	count2 := 0
	for _, char := range data[lowestLayer : lowestLayer+25*6] {
		if char == '1' {
			count1++
		} else if char == '2' {
			count2++
		}
	}

	fmt.Println(count1 * count2)
}
