package robot

import (
	"fmt"
)

type Orientation int

const (
	up    Orientation = 0
	right Orientation = 1
	down  Orientation = 2
	left  Orientation = 3
)

type Point struct {
	X int
	Y int
}
type PaintedSurface map[Point]bool
type RobotPath []Point
type Robot struct {
	Orientation Orientation
	Pos         Point

	Surface PaintedSurface
	Path    RobotPath
}

func (r *Robot) TurnClockWise() {
	r.Orientation = (r.Orientation + 1) % 4
	fmt.Printf("Orientation: %d\n", r.Orientation)
}

func (r *Robot) TurnCounterClockWise() {
	r.Orientation = (r.Orientation + 3) % 4
	fmt.Printf("Orientation: %d\n", r.Orientation)
}

func (r *Robot) MoveOne() {
	switch r.Orientation {
	case up:
		r.Pos.Y--
	case down:
		r.Pos.Y++
	case right:
		r.Pos.X++
	case left:
		r.Pos.X--
	}
	fmt.Printf("Position: %v\n", r.Pos)
}

func (r *Robot) PaintSquare(color int) {
	fmt.Printf("Painting square %v: %d\n", r.Pos, color)
	if color == 1 {
		r.Surface[r.Pos] = true
	} else {
		delete(r.Surface, r.Pos)
	}
}

func (r *Robot) LookAtSquare() int {
	if r.Surface[r.Pos] {
		fmt.Printf("Square %v is painted\n", r.Pos)
		return 1
	} else {
		fmt.Printf("Square %v is not painted\n", r.Pos)
		return 0
	}
}
