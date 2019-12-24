package main

import (
	"log"
	"os"
	"time"

	"./vm"
	"github.com/gdamore/tcell"
)

type StatusCode int

const (
	Wall   StatusCode = 0
	Moved  StatusCode = 1
	Oxygen StatusCode = 2
)

type Direction int

const (
	N Direction = 1
	S Direction = 2
	W Direction = 3
	E Direction = 4
)

type TileType rune

const (
	FreeTile   TileType = '.'
	WallTile   TileType = '#'
	OxygenTile TileType = '$'
	StartTile  TileType = 's'
)

type Point struct {
	x int
	y int
}

func apply(p Point, d Direction) Point {
	x := p.x
	y := p.y
	switch d {
	case N:
		y--
	case S:
		y++
	case E:
		x++
	case W:
		x--
	}
	return Point{x, y}
}

type Robot struct {
	// current position
	pos Point

	// the explored map so far
	mapState map[Point]TileType

	// path from the origin to the current location, used for backtracking
	path []Direction

	m vm.IntegerVM
}

func (r *Robot) init() {
	r.mapState = make(map[Point]TileType)
	r.setStart(Point{0, 0})

	go r.m.Run()
}

func (r *Robot) setPosition(p Point) {
	if _, found := r.mapState[p]; !found {
		r.mapState[p] = FreeTile
	}
	r.pos = p
}

func (r *Robot) setStart(p Point) {
	r.setPosition(p)
	r.mapState[p] = StartTile
}

func (r *Robot) backtrack() bool {
	path := r.path
	if len(path) == 0 {
		return true
	}

	path, d := path[:len(path)-1], path[len(path)-1]

	r.moveOne(reverse(d), true)
	r.path = path

	return false
}

func (r *Robot) moveOne(d Direction, backtrack bool) bool {
	r.m.Inputs <- int(d)
	newpos := apply(r.pos, d)
	status := StatusCode(<-r.m.Outputs)

	if status == Wall {
		r.mapState[newpos] = WallTile
		return false
	}

	if status == Oxygen {
		r.mapState[newpos] = OxygenTile
	}

	// movement happened, update position
	r.setPosition(newpos)
	if !backtrack {
		r.path = append(r.path, d)
	}
	return true
}

func (r *Robot) isUnexplored(d Direction) bool {
	_, found := r.mapState[apply(r.pos, d)]
	return !found
}

func (r *Robot) getUnexplored() (Direction, bool) {
	all := []Direction{N, S, E, W}
	for _, d := range all {
		if r.isUnexplored(d) {
			return d, true
		}
	}
	return N, false
}

func (r *Robot) printMaze(s tcell.Screen) {
	w, h := s.Size()
	ox := w / 2
	oy := h / 2

	// "boundary" of unknown tiles
	for pt, tile := range r.mapState {
		if tile != FreeTile {
			continue
		}
		ch := '?'
		s.SetContent(pt.x-1+ox, pt.y+oy, ch, nil, tcell.StyleDefault)
		s.SetContent(pt.x+1+ox, pt.y+oy, ch, nil, tcell.StyleDefault)
		s.SetContent(pt.x+ox, pt.y-1+oy, ch, nil, tcell.StyleDefault)
		s.SetContent(pt.x+ox, pt.y+1+oy, ch, nil, tcell.StyleDefault)
	}

	// draw explored tiles
	for pt, tile := range r.mapState {
		s.SetContent(pt.x+ox, pt.y+oy, rune(tile), nil, tcell.StyleDefault)
	}

	// robot position
	s.SetContent(r.pos.x+ox, r.pos.y+oy, 'o', nil, tcell.StyleDefault)
}

func reverse(d Direction) Direction {
	var ch Direction
	switch d {
	case N:
		ch = S
	case S:
		ch = N
	case E:
		ch = W
	case W:
		ch = E
	}
	return ch
}

func ExploreMaze(r *Robot) bool {
	done := false
	move, hasMoves := r.getUnexplored()
	if hasMoves {
		r.moveOne(move, false)
	} else {
		done = r.backtrack()
	}
	return done
}

func printStr(s tcell.Screen, x, y int, str string) {
	for i, ch := range str {
		s.SetContent(x+i, y, ch, nil, tcell.StyleDefault)
	}
}

func main() {
	m := vm.CreateVMFromFile("input")
	r := Robot{m: m}
	r.init()

	// set up the screen
	s, e := tcell.NewScreen()
	if e != nil {
		log.Fatal(e)
	}
	if e = s.Init(); e != nil {
		log.Fatal(e)
	}
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.Show()

	defer s.Fini()

	delta := 50 * time.Millisecond
	delay := 1 * delta

	// main loop
	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyPgUp:
					delay -= delta
					if delay <= 0 {
						delay = delta
					}
				case tcell.KeyPgDn:
					delay += delta
				case tcell.KeyUp:
					r.moveOne(N, false)
				case tcell.KeyDown:
					r.moveOne(S, false)
				case tcell.KeyLeft:
					r.moveOne(W, false)
				case tcell.KeyRight:
					r.moveOne(E, false)
				case tcell.KeyRune:
					ch := ev.Rune()
					if ch == 'b' {
						r.backtrack()
					}
				}

				r.printMaze(s)
				s.Sync()
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	go func() {
		for {
			r.printMaze(s)
			time.Sleep(delay)
			if ExploreMaze(&r) {
				break
			}
			s.Sync()
		}
	}()

	<-quit

	f, _ := os.Create("map.txt")
	defer f.Close()

	for y := -40; y < 40; y++ {
		for x := -40; x < 40; x++ {
			p := Point{x, y}
			ch := " "
			if t, found := r.mapState[p]; found {
				ch = string(t)
			}
			f.WriteString(ch)
		}
		f.WriteString("\n")
	}
}
