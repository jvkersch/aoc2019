package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

const (
	PX = 11
	PY = 13
)

type Position struct {
	x     int
	y     int
	theta float64
	dist  int
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

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	posMap := readData(reader)

	// we're at 11,13

	angleMap := make(map[float64][]Position)
	for _, p := range posMap {
		vx := p.x - PX
		vy := p.y - PY
		if vx == 0 && vy == 0 {
			continue
		}

		g := gcd(abs(vx), abs(vy))

		theta := math.Atan2(float64(vx/g), -float64(vy/g))
		if theta < 0 {
			theta += 2.0 * math.Pi
		}

		p.theta = theta
		p.dist = vx*vx + vy*vy

		angleMap[theta] = append(angleMap[theta], p)
	}

	angles := make([]float64, 0, len(angleMap))
	for angle := range angleMap {
		angles = append(angles, angle)
	}
	sort.Slice(angles, func(i, j int) bool {
		return angles[i] < angles[j]
	})

	for _, theta := range angles {
		asteroids := angleMap[theta]

		// sort by distance to the base station
		sort.Slice(asteroids, func(i, j int) bool {
			return asteroids[i].dist < asteroids[j].dist
		})

		// update angles
		for i, _ := range asteroids {
			asteroids[i].theta += 2.0 * float64(i) * math.Pi
		}
		angleMap[theta] = asteroids
	}

	// concatenate everything
	asteroids := make([]Position, 0, len(posMap)-1)
	for _, ast := range angleMap {
		for _, p := range ast {
			asteroids = append(asteroids, p)
		}
	}
	// one last sort
	sort.Slice(asteroids, func(i, j int) bool {
		return asteroids[i].theta < asteroids[j].theta
	})

	fmt.Println(asteroids[199].x*100 + asteroids[199].y)

}
