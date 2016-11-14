package conway

import (
	"fmt"
	"log"
)

// World represents a Game of Life game world with grid and rules.
type World struct {
	g Grid
	s Srv
	b Brth
}

// Grid is the 2 dimensional world the games takes place in.
type Grid [][]uint8

// Srv contains the different numbers of neighbors a cell requires to survive.
type Srv []int

// Brth contains the different numbers of neighbors a cell requires to become alive.
type Brth []int

// NewWorld creates a new World. The world will have x columns and y rows.
// s describes the neighbors required for a cell to survive, and b the number of
// neighbors required for birth. If they are not supplied, standard Conway rules
// will be used (survival at 2 and 3 neighbors, birth at 3 neighbors).
func NewWorld(x, y int, s, b []int) (World, error) {
	if x == 0 || y == 0 {
		return World{nil, nil, nil}, fmt.Errorf("x and y must be greater than 0 (x=%d, y=%d)", x, y)
	}
	if s == nil {
		s = []int{2, 3}
	}
	if b == nil {
		b = []int{3}
	}
	g := make([][]uint8, y)
	for i := range g {
		g[i] = make([]uint8, x)
	}
	world := World{
		g, s, b,
	}
	return world, nil
}

func neighbors(w World, x, y int) int {
	loRow := y - 1
	if y == 0 {
		loRow = 0
	}
	hiRow := y + 1
	if y == len(w.g)-1 {
		hiRow = y
	}
	loCol := x - 1
	if x == 0 {
		loCol = x
	}
	hiCol := x + 1
	if x == len(w.g[0])-1 {
		hiCol = x
	}
	n := 0
	for i := loRow; i <= hiRow; i++ {
		for j := loCol; j <= hiCol; j++ {
			if !(i == y && j == x) && w.g[i][j] > 0 {
				n++
			}
		}
	}
	return n
}

// Next returns the next generation of a world.
func Next(w World) World {
	next, err := NewWorld(len(w.g[0]), len(w.g), w.s, w.b)
	if err != nil {
		log.Fatalf("Could not create new world with parameters (%d, %d)", len(w.g), len(w.g[0]))
	}
	for y, row := range w.g {
		for x, col := range row {
			n := neighbors(w, x, y)
			if col == 0 && n == 3 {
				next.g[y][x] = 1
			} else {
				switch {
				case n < 2:
					next.g[y][x] = 0
				case n == 2 || n == 3:
					next.g[y][x] = col
				case n > 3:
					next.g[y][x] = 0
				}
			}

		}
	}
	return next
}
