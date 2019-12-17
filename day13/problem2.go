package main

import (
	"fmt"
	"time"

	"./vm"
)

const (
	WIDTH  = 41
	HEIGHT = 25
)

type TileType uint8

const (
	Empty  TileType = 0
	Wall   TileType = 1
	Block  TileType = 2
	Paddle TileType = 3
	Ball   TileType = 4
)

type Point struct {
	X int
	Y int
}

type Arcade struct {
	field  [][]TileType // column-major
	ball   Point
	paddle Point
	score  int
}

func (r *Arcade) InitializeArcade() {
	r.field = make([][]TileType, WIDTH)
	for i := 0; i < WIDTH; i++ {
		r.field[i] = make([]TileType, HEIGHT)
	}
}

func (r *Arcade) UpdateFromTiles(data <-chan int) {
	go func() {
		for {
			x := <-data
			y := <-data
			t, more := <-data

			if x == -1 && y == 0 {
				r.score = t
			} else {
				r.field[x][y] = TileType(t)
			}

			// update ball
			if TileType(t) == Ball {
				r.ball.X = x
				r.ball.Y = y
			}
			// update paddle
			if TileType(t) == Paddle {
				r.paddle.X = x
				r.paddle.Y = y
			}

			if !more {
				break
			}
		}
	}()
}

func (r *Arcade) PrintArcade() {
	// for j := 0; j < HEIGHT; j++ {
	// 	for i := 0; i < WIDTH; i++ {
	// 		switch r.field[i][j] {
	// 		case Empty:
	// 			fmt.Print(" ")
	// 		case Wall:
	// 			fmt.Print("W")
	// 		case Block:
	// 			fmt.Print("#")
	// 		case Paddle:
	// 			fmt.Print("=")
	// 		case Ball:
	// 			fmt.Print("o")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	fmt.Printf("Score: %d\n", r.score)
}

func readFromStdin() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		var i int
		for {
			_, err := fmt.Scanf("%d", &i)
			if err != nil {
				break
			}
			out <- i
		}
	}()
	return out
}

func play(r *Arcade, output chan int) {
	for {
		signal := 0
		if r.ball.X > r.paddle.X {
			signal = +1
		} else if r.ball.X < r.paddle.X {
			signal = -1
		}

		output <- signal
		time.Sleep(1 * time.Millisecond)
		r.PrintArcade()
	}
}

func main() {
	m := vm.CreateVMFromFile("input")
	m.Ram[0] = 2 // play for free

	arcade := Arcade{}
	arcade.InitializeArcade()
	arcade.UpdateFromTiles(m.Outputs)

	go m.Run()

	// for input := range readFromStdin() {
	// 	m.Inputs <- input
	// 	time.Sleep(time.Second)
	// 	arcade.PrintArcade()
	// }
	go play(&arcade, m.Inputs)

	<-m.Controls
	arcade.PrintArcade()
}
