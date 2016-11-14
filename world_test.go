package conway

import (
	"testing"
)

func TestNew(t *testing.T) {
	for _, c := range newTestCases {
		got, err := NewWorld(c.x, c.y, nil, nil)
		if err != nil && (c.x == 0 || c.y == 0) {
			t.Skip()
		}
		if len(got.g) != c.y {
			t.Fatalf("Number of rows do not match. Got %d, want %d", len(got.g), c.y)
		}
		if len(got.g[0]) != c.x {
			t.Fatalf("Number of cols do not match. Got %d, want %d", len(got.g[0]), c.x)
		}
		for y, row := range got.g {
			for x, col := range row {
				if col != c.want.g[y][x] {
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
		if len(got.g) != len(c.want.g) {
			t.Fatalf("Case %d: Number of rows do not match. Got %d, want %d", i+1, len(got.g), len(c.want.g))
		}
		if len(got.g[0]) != len(c.want.g[0]) {
			t.Fatalf("Case %d: Number of cols do not match. Got %d, want %d", i+1, len(got.g[0]), len(c.want.g))
		}
		for y, row := range got.g {
			for x, col := range row {
				if col != c.want.g[y][x] {
					t.Fatalf("Case %d: Next world looks weird.\n\nGot:\n%v\n\nWanted:\n%v", i+1, got, c.want)
				}
			}
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewWorld(1000, 1000, nil, nil)
	}
}

func BenchmarkNext(b *testing.B) {
	w, err := NewWorld(1000, 1000, nil, nil)
	b.ResetTimer()
	if err != nil {
		b.Fatalf("BenchmarkNext failed while creating new world.")
	}
	for i := 0; i < b.N; i++ {
		w = Next(w)
	}
}
