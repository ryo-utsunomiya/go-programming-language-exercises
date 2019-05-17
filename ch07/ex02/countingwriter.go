package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriter struct {
	writer io.Writer
	n      int64
}

func (c *countingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.n += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w, 0}
	return cw, &cw.n
}

func main() {
	b := bytes.Buffer{}
	cw, n := CountingWriter(&b)

	io.WriteString(cw, "hello")
	fmt.Println(*n) // 5

	io.WriteString(cw, ", world")
	fmt.Println(*n) // 12
}
