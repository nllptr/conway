package conway

import (
	"errors"
	"log"
)

// World represents a Game of Life game grid.
type World [][]uint8

// New creates a new World.
func New(x, y int) (World, error) {
	if x == 0 || y == 0 {
		return nil, errors.New("x and y must be greater than 0")
	}
	world := make(World, x)
	for i := range world {
		world[i] = make([]uint8, y)
	}
	return world, nil
}

func neighbors(w World, x, y int) int {
	loRow := x - 1
	if x == 0 {
		loRow = 0
	}
	hiRow := x + 1
	if x == len(w)-1 {
		hiRow = x
	}
	loCol := y - 1
	if y == 0 {
		loCol = y
	}
	hiCol := y + 1
	if y == len(w[0])-1 {
		hiCol = y
	}
	n := 0
	for i := loRow; i <= hiRow; i++ {
		for j := loCol; j <= hiCol; j++ {
			if !(i == x && j == y) && w[i][j] > 0 {
				n++
			}
		}
	}
	return n
}

// Next takes a world from one generation to the next.
func Next(w World) World {
	next, err := New(len(w), len(w[0]))
	if err != nil {
		log.Fatalf("Could not create new world")
	}
	for x, row := range w {
		for y, col := range row {
			n := neighbors(w, x, y)
			if col == 0 && n == 3 {
				next[x][y] = 1
			} else {
				switch {
				case n < 2:
					next[x][y] = 0
				case n == 2 || n == 3:
					next[x][y] = col
				case n > 3:
					next[x][y] = 0
				}
			}

		}
	}
	return next
}
