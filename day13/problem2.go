package main

import (
	"fmt"

	"./vm"
)

type TileType int

const (
	Empty  TileType = 0
	Wall   TileType = 1
	Block  TileType = 2
	Paddle TileType = 3
	Ball   TileType = 4
)

type Tile struct {
	X    int
	Y    int
	Type TileType
}

type Data struct {
	paddle int
	ball   int
	score  int
}

func ConsumeRelevantData(data <-chan int) <-chan Data {
	out := make(chan Data)
	go func() {
		paddle := 0
		ball := 0
		score := -1
		haveAllThree := 0
		for {
			x := <-data
			y := <-data
			typ, more := <-data
			if TileType(typ) == Paddle {
				paddle = x
				haveAllThree++
			}
			if TileType(typ) == Ball {
				ball = x
				haveAllThree++
			}
			if x == -1 && y == 0 {
				score = typ
				haveAllThree++
			}
			if haveAllThree == 3 {
				out <- Data{paddle, ball, score}
				haveAllThree = 0
			}
			if !more {
				break
			}
		}
	}()
	return out
}

func main() {
	m := vm.CreateVMFromFile("input")
	m.Ram[0] = 2 // play for free

	go m.Run()

	for data := range ConsumeRelevantData(m.Outputs) {
		signal := 0
		if data.paddle < data.ball {
			signal = -1
		} else if data.paddle > data.ball {
			signal = 1
		}
		m.Inputs <- signal
		fmt.Printf("%d, %d, %d === %d\n", data.paddle, data.ball, signal, data.score)
	}

	<-m.Controls
}
