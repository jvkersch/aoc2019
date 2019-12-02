package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// type RAM []int

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

func doAdd(ram []int, pc int) {
	op1, op2, op3 := ram[pc+1], ram[pc+2], ram[pc+3]
	ram[op3] = ram[op1] + ram[op2]
}

func doMul(ram []int, pc int) {
	op1, op2, op3 := ram[pc+1], ram[pc+2], ram[pc+3]
	ram[op3] = ram[op1] * ram[op2]
}

func run(ram []int) int {
	pc := 0
	for {
		switch ram[pc] {
		case 1:
			doAdd(ram, pc)
			pc += 4
		case 2:
			doMul(ram, pc)
			pc += 4
		case 99:
			// halt
			return ram[0]
		default:
			log.Fatalf("Invalid opcode %d", ram[pc])
		}
		fmt.Println(ram)
	}
	return ram[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ram := readData(reader)

	ram[1] = 12
	ram[2] = 2
	fmt.Println(run(ram))
}
