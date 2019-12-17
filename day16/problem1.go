package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func applyPattern(data []int, pattern []int, it int) int {
	acc := 0
	for i := 0; i < len(data); i++ {
		idx := ((i + 1) / it) % len(pattern)
		acc += data[i] * pattern[idx]
	}
	return abs(acc) % 10
}

func transform(data []int, pattern []int) []int {
	t := make([]int, len(data))
	for i := range t {
		t[i] = applyPattern(data, pattern, i+1)
	}
	return t
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	data := make([]int, 0)
	for s.Scan() {
		digit, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		data = append(data, digit)
	}

	pattern := []int{0, 1, 0, -1}
	for i := 1; i <= 100; i++ {
		data = transform(data, pattern)
	}

	// print first 8 digits
	for i := 0; i < 8; i++ {
		fmt.Print(data[i])
	}
	fmt.Println()
}
