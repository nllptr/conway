package conway

var newTestCases = []struct {
	x, y int
	want World
}{
	{1, 1, World{{0}}},
	{2, 2, World{{0, 0}, {0, 0}}},
	{3, 2, World{{0, 0}, {0, 0}, {0, 0}}},
}

var neighborsTestCases = []struct {
	w    World
	x, y int
	want int
}{
	{
		World{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		1,
		1,
		0,
	},
	{
		World{
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
		},
		1,
		1,
		2,
	},
	{
		World{
			{1, 0, 0},
			{0, 0, 0},
			{1, 0, 0},
		},
		1,
		0,
		2,
	},
	{
		World{
			{0, 0, 0},
			{0, 1, 1},
			{0, 1, 0},
		},
		2,
		2,
		3,
	},
	{
		World{
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		},
		2,
		2,
		4,
	},
}
