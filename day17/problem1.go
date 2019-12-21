package main

import (
	"fmt"

	"./vm"
)

const (
	PATH    = 35
	GAP     = 46
	NEWLINE = 10
)

type point struct {
	x int
	y int
}

func checkIntersection(p point, s map[point]bool) bool {
	n1 := point{p.x + 1, p.y}
	n2 := point{p.x - 1, p.y}
	n3 := point{p.x, p.y + 1}
	n4 := point{p.x, p.y - 1}
	return s[n1] && s[n2] && s[n3] && s[n4]
}

func main() {
	m := vm.CreateVMFromFile("input")
	s := make(map[point]bool)
	go m.Run()

	x := 0
	y := 0
	for value := range m.Outputs {
		if value == NEWLINE {
			y++
			x = 0
			continue
		}
		if value == PATH {
			s[point{x, y}] = true
		}
		x++
	}

	aln := 0
	for p := range s {
		if checkIntersection(p, s) {
			aln += p.x * p.y
		}
	}

	fmt.Println(aln)
}
