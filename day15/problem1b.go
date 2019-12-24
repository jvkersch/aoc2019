package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

type Grid map[Point]bool

func (g *Grid) neighbors(p Point) []Point {
	neigh := make([]Point, 0, 4)

	e := Point{p.x + 1, p.y}
	w := Point{p.x - 1, p.y}
	s := Point{p.x, p.y + 1}
	n := Point{p.x, p.y - 1}

	for _, pt := range []Point{e, w, s, n} {
		if (*g)[pt] {
			neigh = append(neigh, pt)
		}
	}
	return neigh
}

func distance(source, dest Point, grid Grid) int {
	queue := make([]Point, 0)
	marked := make(map[Point]int)

	queue = append(queue, source)
	marked[source] = 0

	for len(queue) > 0 {
		source, queue = queue[0], queue[1:]
		if source == dest {
			break
		}
		for _, n := range grid.neighbors(source) {
			if _, found := marked[n]; !found {
				queue = append(queue, n)
				marked[n] = marked[source] + 1
			}
		}
	}

	return marked[dest]
}

func main() {
	f, err := os.Open("map.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	source := Point{}
	dest := Point{}
	grid := make(Grid)

	y := 0
	for scanner.Scan() {
		t := scanner.Text()
		for x, ch := range t {
			p := Point{x, y}
			if ch == '.' {
				grid[p] = true
			}
			if ch == '$' {
				grid[p] = true
				dest = p
			}
			if ch == 's' {
				grid[p] = true
				source = p
			}
		}
		y++
	}

	fmt.Println(source)
	fmt.Println(dest)
	fmt.Println("--")

	fmt.Println(distance(source, dest, grid))
}
