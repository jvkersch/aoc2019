package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input1")
	if err != nil {
		log.Fatal(err)
	}

	fuel := 0
	totalFuel := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		weight, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel += weight/3 - 2
		totalFuel += computeFuel(weight)
	}

	f.Close()
	fmt.Printf("%d\n", fuel)
	fmt.Printf("%d\n", totalFuel)
}

func computeFuel(weight int) int {
	fuel := weight/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel + computeFuel(fuel)
}
