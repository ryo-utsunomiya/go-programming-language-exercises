package zip

import (
	"archive/zip"
	"io"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch10/ex02/registry"
)

func init() {
	// https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
	registry.RegisterFormat(registry.Format{
		Name:  "zip",
		Magic: "PK",
		Read:  Read,
	})
}

func Read(file *os.File, out io.Writer) error {
	stat, err := file.Stat()
	if err != nil {
		return err
	}

	r, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return err
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(out, rc)
		if err != nil {
			return err
		}
		if err := rc.Close(); err != nil {
			return err
		}
	}

	return nil
}
