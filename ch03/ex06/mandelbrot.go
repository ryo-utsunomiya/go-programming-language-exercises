package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, superSampling(img))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				R: 0xff - contrast*n*1,
				G: 0xff - contrast*n*2,
				B: 0xff - contrast*n*3,
				A: 0xff,
			}
		}
	}
	return color.Black
}

func superSampling(img *image.RGBA) *image.RGBA {
	rect := img.Bounds()
	width := rect.Max.X - rect.Min.X
	height := rect.Max.Y - rect.Min.Y
	
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			newImg.Set(x, y, subPixel(img, x, y))
		}
	}

	return newImg
}

func subPixel(img *image.RGBA, x, y int) color.RGBA {
	pixs := []color.RGBA{
		img.RGBAAt(x, y),
		img.RGBAAt(x+1, y),
		img.RGBAAt(x, y+1),
		img.RGBAAt(x+1, y+1),
	}
	var r, g, b uint32
	for _, p := range pixs {
		r, g, b = r+uint32(p.R), g+uint32(p.G), b+uint32(p.B)
	}
	return color.RGBA{
		R: uint8(r / 4),
		G: uint8(g / 4),
		B: uint8(b / 4),
		A: 0xff,
	}
}
