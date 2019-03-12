package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		lissajous(w, &LissajousParams{
			cycles:  parseOrDefault(q.Get("cycles"), 5),
			size:    parseOrDefault(q.Get("size"), 100),
			nframes: parseOrDefault(q.Get("nframes"), 64),
			delay:   parseOrDefault(q.Get("delay"), 8),
		})
	})
	log.Fatal(http.ListenAndServe("localhost:8798", nil))
}

func parseOrDefault(value string, defaultValue int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
}

const (
	blackIndex = 0
	greenIndex = 1
)

type LissajousParams struct {
	cycles  int
	size    int
	nframes int
	delay   int
}

func lissajous(out io.Writer, params *LissajousParams) {
	const (
		res = 0.001
	)
	var (
		fcycles = float64(params.cycles)
		size    = params.size
		fsize   = float64(params.size)
		nframes = params.nframes
		delay   = params.delay
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < fcycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*fsize+0.5), size+int(y*fsize+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}
}
