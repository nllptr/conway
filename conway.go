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

// New creates a new World.
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

// Next takes a world from one generation to the next.
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

// Img generates a gray scale image from the supplied world and pixel size.
func Img(w World, pix int) *image.Paletted {
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

// WriteGif writes an image to disk.
func WriteGif(img *image.Paletted, fileName string) error {
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

// Anim returns a new animated GIF
func Anim(w World, pix int, steps int, delay int) *gif.GIF {
	var images []*image.Paletted
	var delays []int
	for step := 0; step < steps; step++ {
		img := Img(w, pix)
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
func WriteAnimatedGif(img *gif.GIF, fileName string) error {
	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Could not open or create file '%v'", fileName)
	}
	err = gif.EncodeAll(out, img)
	if err != nil {
		return fmt.Errorf("Could not encode animated GIF")
	}
	return nil
}
