package conway

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

// Config is used to pass information into the write functions.
// PixelSize is used to scale up the actual pixels to produce a more visible Game of Life.
// Steps is the number of generations of the world that will be generated in an animated GIF.
// Delay describes how long each frame will be displayed in an animated gif.
type Config struct {
	PixelSize int // Defaults to 10
	Steps     int // Defaults to 20
	Delay     int // 100ths of a second. Defaults to 50
}

func newImage(w World, c Config) *image.Paletted {
	pix := c.PixelSize
	if pix == 0 {
		pix = 10
	}
	r := image.Rect(0, 0, len(w.g[0])*pix, len(w.g)*pix)
	palette := []color.Color{
		color.White,
		color.Black,
	}
	img := image.NewPaletted(r, palette)

	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y, row := range w.g {
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
