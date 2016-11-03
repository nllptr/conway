package conway

import "testing"

func TestNew(t *testing.T) {
	for _, c := range newTestCases {
		got := New(c.x, c.y)
		if len(got) != c.x {
			t.Fatalf("Number of rows do not match. Got %d, want %d", len(got), c.x)
		}
		if len(got[0]) != c.y {
			t.Fatalf("Number of cols do not match. Got %d, want %d", len(got[0]), c.y)
		}
		for x, row := range got {
			for y, col := range row {
				if col != c.want[x][y] {
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
		for x, row := range got {
			for y, col := range row {
				if col != c.want[x][y] {
					t.Fatalf("Case %d: Next world looks weird.\n\nGot:\n%v\n\nWanted:\n%v", i+1, got, c.want)
				}
			}
		}
	}
}
