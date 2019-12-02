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

type ProgramState int

const (
	running   ProgramState = 0
	outOfTime ProgramState = 1
	halted    ProgramState = 2
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

func doAdd(ram []int, pc int) {
	op1, op2, op3 := ram[pc+1], ram[pc+2], ram[pc+3]
	ram[op3] = ram[op1] + ram[op2]
}

func doMul(ram []int, pc int) {
	op1, op2, op3 := ram[pc+1], ram[pc+2], ram[pc+3]
	ram[op3] = ram[op1] * ram[op2]
}

func run(ram []int, maxcount int) ProgramState {
	pc := 0
	for n := 0; n < maxcount; n++ {
		switch ram[pc] {
		case 1:
			doAdd(ram, pc)
			pc += 4
		case 2:
			doMul(ram, pc)
			pc += 4
		case 99:
			return halted
		default:
			log.Fatalf("Invalid opcode %d", ram[pc])
		}
	}
	return outOfTime
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ram := readData(reader)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			ramCopy := make([]int, len(ram))
			copy(ramCopy, ram)
			ramCopy[1] = i
			ramCopy[2] = j

			state := run(ramCopy, 1000)
			if state == outOfTime {
				fmt.Printf("%d %d: out of time\n", i, j)
			} else { // halted
				if ramCopy[0] == 19690720 {
					fmt.Printf("%d %d: halted correct\n", i, j)
					fmt.Println(100*i + j)
				}
			}
		}
	}
}
