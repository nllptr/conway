package conway

import "testing"

func TestImg(t *testing.T) {
	w := World{
		Grid{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
		},
		Srv{2, 3},
		Brth{3},
	}

	c := Config{
		PixelSize: 25,
	}
	WriteGif(w, c, "mytestgif.gif")

	w = Next(w)
	WriteGif(w, c, "nexted.gif")
}

func TestAnim(t *testing.T) {
	w := World{
		Grid{
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		Srv{2, 3},
		Brth{3},
	}

	c := Config{
		Delay: 20,
	}
	WriteAnimatedGif(w, c, "animated.gif")
}
