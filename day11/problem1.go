package main

import (
	"fmt"

	"./robot"
	"./vm"
)

func CountPath(path robot.RobotPath) int {
	pathMap := make(map[robot.Point]bool)
	for _, p := range path {
		pathMap[p] = true
	}
	return len(pathMap)
}

func main() {
	m := vm.CreateVMFromFile("input")
	r := robot.Robot{}

	r.Surface = make(robot.PaintedSurface)
	r.Path = robot.RobotPath{robot.Point{0, 0}}

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
	fmt.Println(CountPath(r.Path))
}
