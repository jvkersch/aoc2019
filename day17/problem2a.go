package main

import (
	"fmt"

	"./vm"
	// "github.com/gdamore/tcell"
)

type Point struct {
	x int
	y int
}

type Orientation int

const (
	N Orientation = 0
	E Orientation = 1
	S Orientation = 2
	W Orientation = 3
)

type Turn rune

const (
	Straight Turn = 'S'
	Right    Turn = 'R'
	Left     Turn = 'L'
	Back     Turn = 'B'
)

type PathPrimitive struct {
	turn   Turn
	length int
}

type Robot struct {
	// Coordinates of the robot
	pos         Point
	orientation Orientation

	// Current motion primitive
	turn  Turn
	steps int

	// Previous motion primitives
	path []PathPrimitive
	pts  []Point
}

func (r *Robot) turnRight() {
	// reset motion primitive
	r.turn = Right
	// update heading
	r.orientation = Orientation((int(r.orientation) + 1) % 4)
}

func (r *Robot) turnLeft() {
	// reset motion primitive
	r.turn = Left
	// update heading
	r.orientation = Orientation((int(r.orientation) + 3) % 4)
}

func (r *Robot) proposeOne() Point {
	x := r.pos.x
	y := r.pos.y
	switch r.orientation {
	case N:
		y--
	case S:
		y++
	case W:
		x--
	case E:
		x++
	}
	return Point{x, y}
}

func (r *Robot) moveOne() {
	r.pos = r.proposeOne()
	r.steps++

	r.pts = append(r.pts, r.pos)
}

func (r *Robot) commit() {
	r.path = append(r.path, PathPrimitive{r.turn, r.steps})
	r.steps = 0
}

func (r *Robot) canGoForward(m map[Point]bool) bool {
	pos := r.proposeOne()
	return m[pos]
}

func greedyWalk(r *Robot, m map[Point]bool) {
	// try to walk long, straight stretches
	for {
		for r.canGoForward(m) {
			r.moveOne()
		}
		r.commit()

		r.turnLeft()
		if r.canGoForward(m) {
			continue
		}
		r.turnRight()

		r.turnRight()
		if r.canGoForward(m) {
			continue
		}

		break
	}
}

func main() {
	m := vm.CreateVMFromFile("input")
	maze := make(map[Point]bool)
	go m.Run()

	start := Point{}
	x := 0
	y := 0
	for value := range m.Outputs {
		fmt.Printf("%c", value)

		if value == '\n' {
			y++
			x = 0
			continue
		}

		if value == '#' {
			p := Point{x, y}
			maze[p] = true
		}
		if value == '^' {
			start = Point{x, y}
			maze[start] = true
		}

		x++
	}

	r := Robot{pos: start}
	greedyWalk(&r, maze)
	for _, primitive := range r.path[1:] {
		fmt.Printf("%c,%d,", primitive.turn, primitive.length)
	}
	fmt.Println()
}
