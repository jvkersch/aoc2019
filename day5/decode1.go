package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func decodeModes(instr int) (bool, bool, bool) {
	m1 := (instr/10000)%10 == 1
	m2 := (instr/1000)%10 == 1
	m3 := (instr/100)%10 == 1
	return m3, m2, m1
}

func retrieveValue(ram []int, locOrValue int, mode bool) int {
	if !mode {
		locOrValue = ram[locOrValue]
	}
	return ram[locOrValue]
}

func setValue(ram []int, locOrValue int, mode bool, value int) {
	if !mode {
		locOrValue = ram[locOrValue]
	}
	ram[locOrValue] = value
}

func doAdd(ram []int, pc int) int {
	instr := ram[pc]
	m1, m2, m3 := decodeModes(instr)

	value1 := retrieveValue(ram, pc+1, m1)
	value2 := retrieveValue(ram, pc+2, m2)

	setValue(ram, pc+3, m3, value1+value2)
	return pc + 4
}

func doMul(ram []int, pc int) int {
	instr := ram[pc]
	m1, m2, m3 := decodeModes(instr)

	value1 := retrieveValue(ram, pc+1, m1)
	value2 := retrieveValue(ram, pc+2, m2)

	setValue(ram, pc+3, m3, value1*value2)
	return pc + 4
}

func doSave(ram []int, pc int) int {
	// read a value from stdin and save it to RAM
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	value, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		log.Fatal(err)
	}

	instr := ram[pc]
	m1, _, _ := decodeModes(instr)
	setValue(ram, pc+1, m1, value)
	return pc + 2
}

func doLoad(ram []int, pc int) int {
	instr := ram[pc]
	m1, _, _ := decodeModes(instr)
	value := retrieveValue(ram, pc+1, m1)
	fmt.Printf("*** %d ***\n", value)
	return pc + 2
}

func doJumpIfTrue(ram []int, pc int) int {
	instr := ram[pc]
	m1, m2, _ := decodeModes(instr)
	value1 := retrieveValue(ram, pc+1, m1)
	value2 := retrieveValue(ram, pc+2, m2)
	if value1 > 0 {
		return value2
	} else {
		return pc + 3
	}
}

func doJumpIfFalse(ram []int, pc int) int {
	instr := ram[pc]
	m1, m2, _ := decodeModes(instr)
	value1 := retrieveValue(ram, pc+1, m1)
	value2 := retrieveValue(ram, pc+2, m2)
	if value1 == 0 {
		return value2
	} else {
		return pc + 3
	}
}

func doLessThan(ram []int, pc int) int {
	instr := ram[pc]
	m1, m2, m3 := decodeModes(instr)

	value1 := retrieveValue(ram, pc+1, m1)
	value2 := retrieveValue(ram, pc+2, m2)

	var value3 int
	if value1 < value2 {
		value3 = 1
	} else {
		value3 = 0
	}

	setValue(ram, pc+3, m3, value3)
	return pc + 4
}

func doEquals(ram []int, pc int) int {
	instr := ram[pc]
	m1, m2, m3 := decodeModes(instr)

	value1 := retrieveValue(ram, pc+1, m1)
	value2 := retrieveValue(ram, pc+2, m2)

	var value3 int
	if value1 == value2 {
		value3 = 1
	} else {
		value3 = 0
	}

	setValue(ram, pc+3, m3, value3)
	return pc + 4
}

func decodeInstruction(instr int) int {
	return instr % 100
}

func run(ram []int) int {
	pc := 0
	for {
		switch decodeInstruction(ram[pc]) {
		case 1:
			pc = doAdd(ram, pc)
		case 2:
			pc = doMul(ram, pc)
		case 3:
			pc = doSave(ram, pc)
		case 4:
			pc = doLoad(ram, pc)
		case 5:
			pc = doJumpIfTrue(ram, pc)
		case 6:
			pc = doJumpIfFalse(ram, pc)
		case 7:
			pc = doLessThan(ram, pc)
		case 8:
			pc = doEquals(ram, pc)
		case 99:
			// halt
			return ram[0]
		default:
			log.Fatalf("Invalid opcode %d", ram[pc])
		}
	}
	return ram[0]
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	ram := readData(reader)
	run(ram)
}
