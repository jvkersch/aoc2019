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

func readData(reader *bufio.Reader) map[int]int {
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	tokens := strings.Split(strings.TrimSpace(data), ",")
	entries := make(map[int]int)
	for i, token := range tokens {
		n, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		entries[i] = n
	}
	return entries
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	ram := readData(reader)

	inputs := make(chan int)
	outputs := make(chan int)

	m := vm.IntegerVM{Ram: ram, Inputs: inputs, Outputs: outputs}
	go m.Run()

	inputs <- 2
	for out := range outputs {
		fmt.Println(out)
	}
}
