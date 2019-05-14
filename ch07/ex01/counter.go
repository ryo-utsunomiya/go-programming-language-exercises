package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

// bufio.ScanWords
func (c *WordCounter) Write(p []byte) (int, error) {
	advance, _, err := bufio.ScanWords(p, false)
	if err != nil {
		return 0, err
	}
	return advance, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
