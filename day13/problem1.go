package main

import (
	"fmt"

	"./vm"
)

type Tile struct {
	X    int
	Y    int
	Type TileType
}

func ReadTiles(data <-chan int) <-chan Tile {
	out := make(chan Tile)
	go func() {
		for {
			x := <-data
			y := <-data
			typ, more := <-data

			out <- Tile{x, y, TileType(typ)}
			if !more {
				close(out)
				break
			}
		}
	}()
	return out
}

func main() {
	m := vm.CreateVMFromFile("input")

	go m.Run()

	blocks := 0
	for tile := range ReadTiles(m.Outputs) {
		if tile.Type == Block {
			blocks++
		}
	}
	fmt.Println(blocks)
}
