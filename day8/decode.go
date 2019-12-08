package main

import (
	"fmt"
	"io/ioutil"
)

const (
	width  = 25
	height = 6

	TRANSPARENT = '2'
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	// remove trailing line feed
	data = data[0 : len(data)-1]

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			offset := width*i + j
			for data[offset] == TRANSPARENT {
				offset += width * height
			}
			if data[offset] == '1' {
				fmt.Printf("x")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
