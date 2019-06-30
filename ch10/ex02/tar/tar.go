package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch10/ex02/registry"
)

func init() {
	// https://www.gnu.org/software/tar/manual/html_node/Standard.html
	registry.RegisterFormat(registry.Format{
		Name:        "tar",
		Magic:       "ustar",
		MagicOffset: 257,
		Read:        Read,
	})
}

func Read(file *os.File, out io.Writer) error {
	tr := tar.NewReader(file)

	for {
		_, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if _, err := io.Copy(out, tr); err != nil {
			return err
		}
	}

	return nil
}
