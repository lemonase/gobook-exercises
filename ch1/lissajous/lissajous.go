// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

var cpalette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0x15, 0xFF, 0xDC, 0xFF}, // Cyan
	color.RGBA{0x9B, 0x42, 0xF5, 0xFF}, // Purple
	color.RGBA{0x08, 0x5E, 0xFF, 0xFF}, // Indigo
	color.RGBA{0x42, 0x8A, 0xF5, 0xFF}, // Blue
	color.RGBA{0x2D, 0xDE, 0xD0, 0xFF}, // Green
	color.RGBA{0xF5, 0xE6, 0x42, 0xFF}, // Yellow
	color.RGBA{0xFF, 0x55, 0x00, 0xFF}, // Orange
	color.RGBA{0xE3, 0x00, 0x17, 0xFF}, // Red
}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	yellowIndex = 2
	purpleIndex = 3
	greenIndex = 4
	blueIndex = 5
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 100     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 800   // image canvas covers [-size..+size]
		nframes = 128    // number of animation frames
		delay   = 4     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 5.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, cpalette)

		for x := 0; x < 2*size+1; x++ {
			for y := 0; y < 2*size+1; y++ {
				img.Set(x, y, color.Black)
			}
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(t) % uint8(len(cpalette)))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
