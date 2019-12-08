package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"./vm"
)

func readData(reader *bufio.Reader) []int {
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	tokens := strings.Split(strings.TrimSpace(data), ",")
	entries := make([]int, len(tokens))
	for i, token := range tokens {
		n, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		entries[i] = n
	}
	return entries
}

func permutations(arr []int) [][]int {
	// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func computeOutput(ram []int, phases []int) int {
	inputs := make([](chan int), len(phases)+1)

	for i := 0; i < len(phases)+1; i++ {
		inputs[i] = make(chan int, 1)
	}

	for i := 0; i < len(phases); i++ {
		copyram := make([]int, len(ram))
		copy(copyram, ram)

		m := vm.IntegerVM{
			Id:      i,
			Ram:     copyram,
			Inputs:  inputs[i],
			Outputs: inputs[i+1]}
		go m.Run()
		inputs[i] <- phases[i]
	}

	inputs[0] <- 0

	var lastOutput int
	for output := range inputs[len(phases)] {
		inputs[0] <- output
		lastOutput = output
	}

	return lastOutput
}

func maximizeOutput(ram []int) int {
	output := -100000
	perms := permutations([]int{9, 8, 7, 6, 5})
	for _, p := range perms {
		out := computeOutput(ram, p)
		if out > output {
			output = out
		}
	}
	return output
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	ram := readData(reader)

	output := maximizeOutput(ram)
	fmt.Println(output)
}
