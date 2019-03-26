package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 230
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func allIsFinite(fs []float64) bool {
	for _, f := range fs {
		if math.IsInf(f, 0) {
			return false
		}
	}
	return true
}

func svg(w io.Writer) {
	fmt.Fprintf(w, `
<svg xmlns="http://www.w3.org/2000/svg" style="stroke:grey;fill:white;stroke-width:0.7" width="%d" height="%d">
`, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if allIsFinite([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(w, `
<polygon points="%g %g %g %g %g %g %g %g"/>
`, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
	})
	addr := "localhost:4190"
	fmt.Printf("Starting server at: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
