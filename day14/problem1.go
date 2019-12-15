package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type component struct {
	from       string
	to         string
	fromWeight int
	toWeight   int
}

type reactionmap map[string][]component

func parseOne(token string) (int, string) {
	tokens := strings.Split(strings.TrimSpace(token), " ")
	if len(tokens) != 2 {
		panic("malformed token")
	}
	count, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	return count, tokens[1]
}

func parseReactions(reader *bufio.Reader) reactionmap {
	m := make(reactionmap)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "=>")
		if len(tokens) != 2 {
			continue
		}
		producers := strings.Split(tokens[0], ",")
		cweight, cname := parseOne(tokens[1])

		m[cname] = make([]component, 0, len(producers))
		for _, prod := range producers {
			pweight, pname := parseOne(prod)
			comp := component{
				from:       pname,
				fromWeight: pweight,
				to:         cname,
				toWeight:   cweight}
			m[cname] = append(m[cname], comp)
		}
	}
	return m
}

func visit(node string, m reactionmap, visited map[string]bool, nodes []string) []string {
	for _, edge := range m[node] {
		neighbor := edge.from
		if !visited[neighbor] {
			nodes = visit(neighbor, m, visited, nodes)
		}
	}
	visited[node] = true
	nodes = append(nodes, node)
	return nodes
}

func getTopologicalOrder(m reactionmap) []string {
	visited := make(map[string]bool)
	nodes := make([]string, 0)

	nodes = visit("FUEL", m, visited, nodes)
	return nodes
}

func ceildiv(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

func calculateAmounts(fuelamount int, nodes []string, m reactionmap) int {
	amounts := make(map[string]int)
	amounts["FUEL"] = fuelamount

	for i := len(nodes) - 1; i >= 0; i-- {
		c := nodes[i]
		for _, edge := range m[c] {
			amount := ceildiv(amounts[edge.to], edge.toWeight) * edge.fromWeight
			amounts[edge.from] += amount
		}

	}
	return amounts["ORE"]
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	m := parseReactions(reader)
	nodes := getTopologicalOrder(m)
	oreAmount := calculateAmounts(1, nodes, m)
	fmt.Println(oreAmount)

	oreInStore := 1000000000000

	fuelLower := oreInStore / oreAmount
	fuelHigher := 3 * fuelLower

	for fuelLower < fuelHigher {
		fmt.Printf("%d, %d\n", fuelLower, fuelHigher)
		fuelMiddle := (fuelLower + fuelHigher) / 2
		oreMiddle := calculateAmounts(fuelMiddle, nodes, m)
		if oreInStore < oreMiddle {
			fuelHigher = fuelMiddle - 1
		} else {
			fuelLower = fuelMiddle
		}
	}
	fmt.Printf("FUEL: %d, ORE NEEDED: %d\n",
		fuelLower, calculateAmounts(fuelLower, nodes, m))
}
