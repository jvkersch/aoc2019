package main

import (
	"fmt"

	"./robot"
	"./vm"
)

func main() {
	m := vm.CreateVMFromFile("input")
	r := robot.Robot{}

	startingPoint := robot.Point{0, 0}
	r.Surface = robot.PaintedSurface{startingPoint: true}
	r.Path = robot.RobotPath{startingPoint}

	go m.Run()
	go func() {
		for {
			status := r.LookAtSquare()
			fmt.Printf("Submitting %d\n", status)
			m.Inputs <- status

			toPaint := <-m.Outputs
			toTurn := <-m.Outputs

			fmt.Printf("Paint command: %d\n", toPaint)
			fmt.Printf("Turn command: %d\n", toTurn)

			r.PaintSquare(toPaint)

			if toTurn == 1 {
				r.TurnClockWise()
			} else {
				r.TurnCounterClockWise()
			}
			r.MoveOne()
			r.Path = append(r.Path, r.Pos)
			fmt.Println("===== End of move =====")
		}
	}()

	<-m.Controls

	// display what was painted.
	xmin, xmax := 10000, -10000
	ymin, ymax := 10000, -10000
	for pt := range r.Surface {
		if pt.X < xmin {
			xmin = pt.X
		} else if pt.X > xmax {
			xmax = pt.X
		}
		if pt.Y < ymin {
			ymin = pt.Y
		} else if pt.Y > ymax {
			ymax = pt.Y
		}
	}
	fmt.Printf("Canvas size: (%d, %d) to (%d, %d)\n", xmin, ymin, xmax, ymax)

	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if r.Surface[robot.Point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
