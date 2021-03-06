package main

import (
	"fmt"
	"math"
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
	return 0.1 * (math.Cos(x) + math.Cos(y))
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

func main() {
	fmt.Printf(`
<svg xmlns="http://www.w3.org/2000/svg" style="stroke:grey;fill:white;stroke-width:0.7" width="%d" height="%d">
`, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if allIsFinite([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Printf(`
<polygon points="%g,%g %g,%g %g,%g %g,%g"/>
`, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Println("</svg>")
}
