package vm

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

func CreateVMFromFile(fname string) IntegerVM {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	ram := readData(reader)

	inputs := make(chan int)
	outputs := make(chan int)
	controls := make(ControlChannel)

	return IntegerVM{Ram: ram, Inputs: inputs, Outputs: outputs, Controls: controls}
}
