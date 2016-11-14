package conway

var newTestCases = []struct {
	x, y int
	want World
}{
	{
		1,
		1,
		World{
			Grid{{0}},
			Srv{2, 3},
			Brth{3},
		},
	},
	{
		2,
		2,
		World{
			Grid{{0, 0}, {0, 0}},
			Srv{2, 3},
			Brth{3},
		},
	},
	{
		3,
		2,
		World{
			Grid{{0, 0, 0}, {0, 0, 0}},
			Srv{2, 3},
			Brth{3},
		},
	},
	{
		0,
		0,
		World{
			Grid{{0}},
			Srv{2, 3},
			Brth{3},
		},
	},
}

var neighborsTestCases = []struct {
	w    World
	x, y int
	want int
}{
	{
		World{
			Grid{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		1,
		1,
		0,
	},
	{
		World{
			Grid{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		1,
		1,
		2,
	},
	{
		World{
			Grid{
				{1, 0, 0},
				{0, 0, 0},
				{1, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		0,
		1,
		2,
	},
	{
		World{
			Grid{
				{0, 0, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		2,
		2,
		3,
	},
	{
		World{
			Grid{
				{0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		2,
		2,
		4,
	},
}

var nextTestCases = []struct {
	world, want World
}{
	{
		World{
			Grid{{0}},
			Srv{2, 3},
			Brth{3},
		},
		World{
			Grid{{0}},
			Srv{2, 3},
			Brth{3},
		},
	},
	// Block
	{
		World{
			Grid{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		World{
			Grid{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
	},
	// Beehive
	{
		World{
			Grid{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 0, 0},
				{0, 1, 0, 0, 1, 0},
				{0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		World{
			Grid{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 0, 0},
				{0, 1, 0, 0, 1, 0},
				{0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
	},
	// Blinker
	{
		World{
			Grid{
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		World{
			Grid{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
	},
	// Toad
	{
		World{
			Grid{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 0},
				{0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
		World{
			Grid{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 0},
				{0, 1, 0, 0, 1, 0},
				{0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
	},
}

var readLife106TestCases = []struct {
	input string
	x, y  int
	want  World
}{
	{
		"#Life 1.05\n0 0\n0 1\n1 0\n1 2\n2 0",
		5, 5,
		World{
			Grid{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
	},
	{
		"#Life 1.06\n0 0\n0 1\n1 0\n1 2\n2 0",
		5, 5,
		World{
			Grid{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 0},
			},
			Srv{2, 3},
			Brth{3},
		},
	},
}

var centerOffsetTestCases = []struct {
	wX, wY int
	oX, oY int
}{
	{3, 3, 1, 1},
	{10, 10, 4, 4},
	{13, 7, 6, 3},
}
