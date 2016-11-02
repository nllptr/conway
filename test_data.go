package conway

var newTestCases = []struct {
	x, y int
	want World
}{
	{1, 1, World{{false}}},
	{2, 2, World{{false, false}, {false, false}}},
	{3, 2, World{{false, false}, {false, false}, {false, false}}},
}

var neighborsTestCases = []struct {
	w    World
	x, y int
	want int
}{
	{
		World{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
		1,
		1,
		0,
	},
	{
		World{
			{false, true, false},
			{false, true, false},
			{false, true, false},
		},
		1,
		1,
		2,
	},
	{
		World{
			{true, false, false},
			{false, true, false},
			{true, false, false},
		},
		1,
		0,
		2,
	},
	{
		World{
			{false, false, false},
			{false, true, true},
			{false, true, false},
		},
		2,
		2,
		3,
	},
}
