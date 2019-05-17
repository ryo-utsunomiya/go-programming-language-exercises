package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r io.Reader
	n int64
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.n {
		p = p[0:r.n]
	}
	n, err = r.r.Read(p)
	r.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main() {
	limit := int64(5)
	s := "Hello, world!"
	n := int64(len(s))
	r := strings.NewReader(s)
	b := make([]byte, n)
	lr := LimitReader(r, limit)
	fmt.Println(lr.Read(b)) // 5, nil
	fmt.Println(string(b))  // Hello
}
