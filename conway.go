package conway

// World represents a Game of Life game grid.
type World [][]uint8

// New creates a new World.
func New(x, y int) World {
	world := make(World, x)
	for i := range world {
		world[i] = make([]uint8, y)
	}
	return world
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
	return New(1, 1)
}
