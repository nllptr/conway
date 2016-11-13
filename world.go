package conway

import "fmt"

// World represents a Game of Life game grid.
type World [][]uint8

// NewWorld creates a new World. The world will have x columns and y rows.
func NewWorld(x, y int) (World, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("x and y must be greater than 0 (x=%d, y=%d)", x, y)
	}
	world := make(World, y)
	for i := range world {
		world[i] = make([]uint8, x)
	}
	return world, nil
}

func neighbors(w *World, x, y int) int {
	loRow := y - 1
	if y == 0 {
		loRow = 0
	}
	hiRow := y + 1
	if y == len(*w)-1 {
		hiRow = y
	}
	loCol := x - 1
	if x == 0 {
		loCol = x
	}
	hiCol := x + 1
	if x == len((*w)[0])-1 {
		hiCol = x
	}
	n := 0
	for i := loRow; i <= hiRow; i++ {
		for j := loCol; j <= hiCol; j++ {
			if !(i == y && j == x) && (*w)[i][j] > 0 {
				n++
			}
		}
	}
	return n
}

// Next returns the next generation of a world.
func Next(old *World, new *World) {
	/*
		next, err := NewWorld(len(w[0]), len(w))
		if err != nil {
			log.Fatalf("Could not create new world with parameters (%d, %d)", len(w), len(w[0]))
		}
	*/
	for y, row := range *old {
		for x, col := range row {
			n := neighbors(old, x, y)
			if col == 0 && n == 3 {
				(*new)[y][x] = 1
			} else {
				switch {
				case n < 2:
					(*new)[y][x] = 0
				case n == 2 || n == 3:
					(*new)[y][x] = col
				case n > 3:
					(*new)[y][x] = 0
				}
			}

		}
	}
	//return next
}
