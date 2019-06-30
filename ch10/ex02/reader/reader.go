package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch10/ex02/registry"
	_ "github.com/ryo-utsunomiya/go-programming-language-exercises/ch10/ex02/tar"
	_ "github.com/ryo-utsunomiya/go-programming-language-exercises/ch10/ex02/zip"
)

func Read(name string, out io.Writer) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}

	r := bufio.NewReader(f)

	var found *registry.Format
	for _, f := range registry.GetAllFormats() {
		p, err := r.Peek(f.MagicOffset + len(f.Magic))
		if err != nil {
			continue
		}
		if string(p[f.MagicOffset:]) == f.Magic {
			found = &f
			break
		}
	}
	if found == nil {
		return fmt.Errorf("unknown format")
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	if err := found.Read(f, out); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
