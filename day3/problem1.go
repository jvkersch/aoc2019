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

func traceWire(wire Wire, wireMap WireMap) WireMap {
	x, y := 0, 0
	visited := make(map[Point]bool)
	for _, wireLength := range wire {
		for i := 0; i < wireLength.length; i++ {
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
			if !visited[p] {
				visited[p] = true
				wireMap[p]++
			}
		}
	}
	return wireMap
}

func findWireCrossings(wireMap WireMap) []Point {
	crossings := make([]Point, 0)
	for point, multiplicity := range wireMap {
		if multiplicity > 1 {
			crossings = append(crossings, point)
		}
	}
	return crossings
}

func getClosestManhattanDistance(points []Point) int {
	dist := Manhattan(points[0])
	for _, p := range points[1:] {
		d := Manhattan(p)
		if d < dist {
			dist = d
		}
	}
	return dist
}

func Manhattan(p Point) int {
	return Abs(p.x) + Abs(p.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func ReadOrDie(reader *bufio.Reader) string {
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	wireStr1 := ReadOrDie(reader)
	wireStr2 := ReadOrDie(reader)

	wire1 := parseWire(wireStr1)
	wire2 := parseWire(wireStr2)

	wireMap := make(WireMap)
	wireMap = traceWire(wire1, wireMap)
	wireMap = traceWire(wire2, wireMap)

	crossings := findWireCrossings(wireMap)
	fmt.Println(crossings)

	dist := getClosestManhattanDistance(crossings)
	fmt.Println(dist)
}
