package main

import (
	"flag"
	"fmt"

	"os"

	"github.com/nllptr/conway"
)

func main() {
	xp := flag.Int("width", 100, "The width (X axis) of the World.")
	yp := flag.Int("height", 50, "The height (Y axis) of the world.")
	cp := flag.Int("cellsize", 10, "The cell size in pizels.")
	ap := flag.Bool("animate", false, "Whether or not the output GIF should be animated.")
	fp := flag.Int("frames", 20, "Number of frames in the anmation.")
	dp := flag.Int("delay", 20, "Animation frame delay in 100ths of a second.")
	ip := flag.String("in", "", "The input file.")
	op := flag.String("out", "out.gif", "The name of the output GIF.")

	flag.Parse()

	if *ip == "" {
		fmt.Fprint(os.Stderr, "Flag -in is required.\n\n")
		flag.PrintDefaults()
		return
	}

	w, err := conway.NewWorld(*xp, *yp)
	if err != nil {
		fmt.Println(err)
	}
	r, err := os.Open(*ip)
	if err != nil {
		fmt.Printf("Could not open file '%s'\n", *ip)
	}
	conway.ReadLife106(r, &w)

	config := conway.Config{
		PixelSize: *cp,
		Steps:     *fp,
		Delay:     *dp,
	}
	if !*ap {
		err = conway.WriteGif(w, config, *op)
	} else {
		err = conway.WriteAnimatedGif(w, config, *op)
	}
}
