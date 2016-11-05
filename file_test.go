package conway

import "testing"
import "strings"

func TestRead106(t *testing.T) {
	for i, c := range read106TestCases {
		w, err := NewWorld(c.x, c.y)
		if err != nil {
			t.Fatalf("Case %d: Failed while creating world.", i+1)
		}
		reader := strings.NewReader(c.input)
		err = Read106(reader, &w)
		if err != nil {
			t.Logf("Case %d: Read106 returned an error:\n%v", i, err)
		}
		for y, row := range w {
			for x, col := range row {
				if col != c.want[y][x] {
					t.Fatalf("Case %d: Failed!\nGot:\n%v\n\nWanted:\n%v", i, w, c.want)
				}
			}
		}
	}
}

func TestCenterOffset(t *testing.T) {
	for i, c := range centerOffsetTestCases {
		w, err := NewWorld(c.wX, c.wY)
		if err != nil {
			t.Fatalf("Case %d: Failed while creating world.", i+1)
		}
		gotX, gotY := centerOffset(&w)
		if gotX != c.oX || gotY != c.oY {
			t.Fatalf("Case %d: Wanted (%d,%d), got (%d,%d)", i, c.oX, c.oY, gotX, gotY)
		}
	}
}
