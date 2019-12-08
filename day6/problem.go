package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	name  string
	depth int
}

func readData(filename string) map[string][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[string][]string)
	s := bufio.NewScanner(f)
	for s.Scan() {
		parts := strings.Split(strings.TrimSpace(s.Text()), ")")
		m[parts[0]] = append(m[parts[0]], parts[1])
	}
	return m
}

func traverse(orbitMap map[string][]string) map[string]int {
	queue := []string{"COM"}
	distance := map[string]int{"COM": 0}
	bfs(queue, distance, orbitMap)
	return distance
}

func bfs(queue []string,
	distance map[string]int,
	orbitMap map[string][]string) map[string]int {

	if len(queue) == 0 {
		return distance
	}

	for _, n := range orbitMap[queue[0]] {
		queue = append(queue, n)
		distance[n] = distance[queue[0]] + 1
	}
	return bfs(queue[1:], distance, orbitMap)
}

func computeTotalDistance(distances map[string]int) int {
	d := 0
	for _, val := range distances {
		d += val
	}
	return d
}

func invertOrbitMap(orbitMap map[string][]string) map[string]string {
	parentMap := make(map[string]string)
	for parent, children := range orbitMap {
		for _, child := range children {
			parentMap[child] = parent
		}
	}
	return parentMap
}

func findAncestors(parentMap map[string]string, node string) []string {
	ancestors := []string{}
	for node != "COM" {
		node = parentMap[node]
		ancestors = append(ancestors, node)
	}
	return ancestors
}

func findCommonAncestor(parentMap map[string]string, node1 string, node2 string) int {
	ancestors1 := findAncestors(parentMap, node1)
	ancestors2 := findAncestors(parentMap, node2)
	fmt.Println(ancestors1)
	fmt.Println(ancestors2)

	i := len(ancestors1) - 1
	j := len(ancestors2) - 1
	for {
		if ancestors1[i] != ancestors2[j] {
			break
		}
		i--
		j--
	}

	// fmt.Printf("%s, %s\n", ancestors1[i], ancestors2[i])

	return i + j + 2
}

func main() {
	orbitMap := readData("input")

	dists := traverse(orbitMap)
	fmt.Println(computeTotalDistance(dists))

	parentMap := invertOrbitMap(orbitMap)
	d := findCommonAncestor(parentMap, "YOU", "SAN")
	fmt.Println(d)
}
