package conway

import "testing"

func TestNew(t *testing.T) {
	for _, c := range newTestCases {
		got, err := New(c.x, c.y)
		if err != nil && (c.x == 0 || c.y == 0) {
			t.Skip()
		}
		if len(got) != c.y {
			t.Fatalf("Number of rows do not match. Got %d, want %d", len(got), c.y)
		}
		if len(got[0]) != c.x {
			t.Fatalf("Number of cols do not match. Got %d, want %d", len(got[0]), c.x)
		}
		for y, row := range got {
			for x, col := range row {
				if col != c.want[y][x] {
					t.Errorf("New(%d, %d) == \n%v, Wanted:\n%v", c.x, c.y, got, c.want)
				}
			}
		}
	}
}

func TestNeighbors(t *testing.T) {
	for i, c := range neighborsTestCases {
		got := neighbors(c.w, c.x, c.y)
		if got != c.want {
			t.Fatalf("Case %d: Number of neighbors is incorrect. Got %d, wanted %d", i+1, got, c.want)
		}
	}
}

func TestNext(t *testing.T) {
	for i, c := range nextTestCases {
		got := Next(c.world)
		if len(got) != len(c.want) {
			t.Fatalf("Case %d: Number of rows do not match. Got %d, want %d", i+1, len(got), len(c.want))
		}
		if len(got[0]) != len(c.want[0]) {
			t.Fatalf("Case %d: Number of cols do not match. Got %d, want %d", i+1, len(got[0]), len(c.want))
		}
		for y, row := range got {
			for x, col := range row {
				if col != c.want[y][x] {
					t.Fatalf("Case %d: Next world looks weird.\n\nGot:\n%v\n\nWanted:\n%v", i+1, got, c.want)
				}
			}
		}
	}
}

func TestImg(t *testing.T) {
	w1 := World{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	img := Img(w1, 10)
	WriteGif(img, "mytestgif.gif")

	w1 = Next(w1)
	img = Img(w1, 10)
	WriteGif(img, "nexted.gif")
}

func TestAnimatedImg(t *testing.T) {
	w1 := World{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	gif := Anim(w1, 10, 20, 50)
	WriteAnimatedGif(gif, "animated.gif")
}
