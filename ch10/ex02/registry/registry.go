package registry

import (
	"io"
	"os"
)

type Format struct {
	Name        string
	Magic       string
	MagicOffset int
	Read        func(file *os.File, out io.Writer) error
}

var formats []Format

func RegisterFormat(f Format) {
	formats = append(formats, f)
}

func GetAllFormats() []Format {
	return formats
}
