package conway

// World represents a Game of Life game grid.
type World [][]bool

// New creates a new World.
func New(x, y int) World {
	world := make(World, x)
	for i := range world {
		world[i] = make([]bool, y)
	}
	return world
}

func neighbors(w World, x, y int) int {
	return 0
}
