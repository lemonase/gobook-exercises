// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
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
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1 // next color in palette
	yellowIndex = 2
	purpleIndex = 3
	greenIndex  = 4
	blueIndex   = 5
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Requested URL:", r.URL)

			fmt.Printf("\n**********\n")
			fmt.Println("[Headers]")
			for n, v := range r.Header {
				fmt.Fprintf(os.Stdout, "Header [%q] = %q\n", n, v)
			}

			fmt.Println()
			fmt.Println("[Form Parameters]")
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			for n, v := range r.Form {
				fmt.Fprintf(os.Stdout, "Form [%q] = %q\n", n, v)
			}
			fmt.Println("**********")

			start := time.Now()

			lissajous(w, r.Form)
			fmt.Printf("%.2fs to create GIF\n\n", time.Since(start).Seconds())
		}

		http.HandleFunc("/", handler)
		fmt.Printf("%s\n\n", "Listening on http://localhost:8000")
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

		return
	}
	// lissajous(os.Stdout)
}

func lissajous(out io.Writer, formValues map[string][]string) {
	var (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	for k, v := range formValues {
		if k == "cycles" {
			var err error
			cycles, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Printf("%T, %v", cycles, cycles)
			}
		}
		if k == "res" {
			var err error
			res, err = strconv.ParseFloat(v[0], 32)
			if err != nil {
				fmt.Printf("%T, %v", res, res)
			}
		}
		if k == "size" {
			var err error
			size, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Printf("%T, %v", size, size)
			}
		}
		if k == "nframes" {
			var err error
			nframes, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Printf("%T, %v", nframes, nframes)
			}
		}
		if k == "delay" {
			var err error
			delay, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Printf("%T, %v", delay, delay)
			}
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, cpalette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += float64(res) {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				uint8(t)%uint8(len(cpalette)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
