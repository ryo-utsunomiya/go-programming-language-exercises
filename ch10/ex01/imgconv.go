package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	var format = flag.String("f", "jpg", "format")
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Errorf("failed to read: %v", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)

	switch *format {
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95})
	case "png":
		err = png.Encode(os.Stdout, img)
	case "gif":
		err = gif.Encode(os.Stdout, img, &gif.Options{})
	default:
		fmt.Fprintln(os.Stderr, "unknown format")
		os.Exit(1)
	}

	if err != nil {
		fmt.Errorf("failed to convert: %v", err)
		os.Exit(1)
	}
}
