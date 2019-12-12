package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	x int
	y int
}

func readData(reader *bufio.Reader) []Position {
	pos := make([]Position, 0)
	y := 0
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		for x, ch := range data {
			if ch == '#' {
				pos = append(pos, Position{x: x, y: y})
			}
		}
		y++
	}

	return pos
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func computeVisibility(pos Position, posMap []Position) int {
	count := 0
	seen := make(map[Position]bool)
	for _, asteroid := range posMap {
		vx := asteroid.x - pos.x
		vy := asteroid.y - pos.y
		if vx == 0 && vy == 0 {
			continue
		}

		g := gcd(abs(vx), abs(vy))
		vx /= g
		vy /= g

		p := Position{x: vx, y: vy}
		if !seen[p] {
			seen[p] = true
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	posMap := readData(reader)

	vis := 0
	for _, pos := range posMap {
		v := computeVisibility(pos, posMap)
		if v > vis {
			vis = v
		}
	}

	fmt.Println(vis)
}
