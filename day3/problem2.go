package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type WireLength struct {
	direction byte
	length    int
}

type Wire []WireLength

type Point struct {
	x int
	y int
}
type WireMap map[Point]int

func parseWire(wireStr string) Wire {
	tokens := strings.Split(strings.TrimSpace(wireStr), ",")
	wire := make(Wire, len(tokens))
	for i, token := range tokens {
		token = strings.TrimSpace(token)
		direction := token[0]
		length, err := strconv.Atoi(token[1:])
		if err != nil {
			log.Fatal(err)
		}
		wire[i] = WireLength{direction: direction, length: length}
	}
	return wire
}

func traceWire(wire Wire) WireMap {
	wireMap := make(WireMap)
	x, y := 0, 0
	runlength := 0
	for _, wireLength := range wire {
		for i := 0; i < wireLength.length; i++ {
			runlength++
			switch wireLength.direction {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			default:
				log.Fatalf("Invalid direction: %s", wireLength.direction)
			}
			p := Point{x, y}
			if _, visited := wireMap[p]; !visited {
				wireMap[p] = runlength
			}
		}
	}
	return wireMap
}

func findWireCrossings(wireMap1 WireMap, wireMap2 WireMap) []Point {
	crossings := make([]Point, 0)
	for point1, _ := range wireMap1 {
		if _, present := wireMap2[point1]; present {
			crossings = append(crossings, point1)
		}
	}
	return crossings
}

func ReadOrDie(reader *bufio.Reader) string {
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func getSmallestDelay(wireMap1 WireMap, wireMap2 WireMap, crossings []Point) int {
	dist := wireMap1[crossings[0]] + wireMap2[crossings[0]]
	for _, p := range crossings[1:] {
		d := wireMap1[p] + wireMap2[p]
		if d < dist {
			dist = d
		}
	}
	return dist
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	wireStr1 := ReadOrDie(reader)
	wireStr2 := ReadOrDie(reader)

	wire1 := parseWire(wireStr1)
	wire2 := parseWire(wireStr2)

	wireMap1 := traceWire(wire1)
	wireMap2 := traceWire(wire2)

	fmt.Println(wireMap1)
	fmt.Println(wireMap2)

	crossings := findWireCrossings(wireMap1, wireMap2)
	fmt.Println(crossings)

	dist := getSmallestDelay(wireMap1, wireMap2, crossings)
	fmt.Println(dist)
}
