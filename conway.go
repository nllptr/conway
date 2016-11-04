package conway

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"os"
)

// World represents a Game of Life game grid.
type World [][]uint8

// Config is used to pass information into the write functions.
// PixelSize is used to scale up the actual pixels to produce a more visible Game of Life.
// Steps is the number of generations of the world that will be generated in an animated GIF.
// Delay describes how long each frame will be displayed in an animated gif.
type Config struct {
	PixelSize int // Defaults to 10
	Steps     int // Defaults to 20
	Delay     int // 100ths of a second. Defaults to 50
}

// New creates a new World. The world will have x columns and y rows.
func New(x, y int) (World, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("x and y must be greater than 0 (x=%d, y=%d)", x, y)
	}
	world := make(World, y)
	for i := range world {
		world[i] = make([]uint8, x)
	}
	return world, nil
}

func neighbors(w World, x, y int) int {
	loRow := y - 1
	if y == 0 {
		loRow = 0
	}
	hiRow := y + 1
	if y == len(w)-1 {
		hiRow = y
	}
	loCol := x - 1
	if x == 0 {
		loCol = x
	}
	hiCol := x + 1
	if x == len(w[0])-1 {
		hiCol = x
	}
	n := 0
	for i := loRow; i <= hiRow; i++ {
		for j := loCol; j <= hiCol; j++ {
			if !(i == y && j == x) && w[i][j] > 0 {
				n++
			}
		}
	}
	return n
}

// Next returns the next generation of a world.
func Next(w World) World {
	next, err := New(len(w[0]), len(w))
	if err != nil {
		log.Fatalf("Could not create new world with parameters (%d, %d)", len(w), len(w[0]))
	}
	for y, row := range w {
		for x, col := range row {
			n := neighbors(w, x, y)
			if col == 0 && n == 3 {
				next[y][x] = 1
			} else {
				switch {
				case n < 2:
					next[y][x] = 0
				case n == 2 || n == 3:
					next[y][x] = col
				case n > 3:
					next[y][x] = 0
				}
			}

		}
	}
	return next
}

func newImage(w World, c Config) *image.Paletted {
	pix := c.PixelSize
	if pix == 0 {
		pix = 10
	}
	r := image.Rect(0, 0, len(w[0])*pix, len(w)*pix)
	palette := []color.Color{
		color.White,
		color.Black,
	}
	img := image.NewPaletted(r, palette)

	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y, row := range w {
		for x, col := range row {
			if col > 0 {
				pixel := image.Rect(x*pix, y*pix, x*pix+pix, y*pix+pix)
				draw.Draw(img, pixel, &image.Uniform{color.Black}, image.ZP, draw.Src)
			}
		}
	}
	return img
}

// WriteGif writes an image to disk as a GIF.
func WriteGif(w World, c Config, fileName string) error {
	img := newImage(w, c)
	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Could not open or create file '%v'", fileName)
	}
	var opts gif.Options
	opts.NumColors = 256
	err = gif.Encode(out, img, &opts)
	if err != nil {
		return fmt.Errorf("Could not encode GIF")
	}
	return nil
}

func newAnim(w World, c Config) *gif.GIF {
	var images []*image.Paletted
	var delays []int
	steps := c.Steps
	if steps == 0 {
		steps = 20
	}
	delay := c.Delay
	if delay == 0 {
		delay = 50
	}
	for step := 0; step < steps; step++ {
		img := newImage(w, c)
		images = append(images, img)
		delays = append(delays, delay)

		w = Next(w)
	}
	return &gif.GIF{
		Image: images,
		Delay: delays,
	}
}

// WriteAnimatedGif writes an animated GIF to disk.
func WriteAnimatedGif(w World, c Config, fileName string) error {
	anim := newAnim(w, c)
	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Could not open or create file '%v'", fileName)
	}
	err = gif.EncodeAll(out, anim)
	if err != nil {
		return fmt.Errorf("Could not encode animated GIF")
	}
	return nil
}
