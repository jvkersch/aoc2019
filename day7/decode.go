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

func computeOutput(ram []int, phases []int) int {
	previous := 0
	for i := 0; i < len(phases); i++ {
		inputs := make(chan int)
		outputs := make(chan int)

		m := vm.IntegerVM{Ram: ram, Inputs: inputs, Outputs: outputs}
		go m.Run()

		inputs <- phases[i]
		inputs <- previous
		previous = <-outputs
	}
	return previous
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

func maximizeOutput(ram []int) int {
	output := -100000
	perms := permutations([]int{0, 1, 2, 3, 4})
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
