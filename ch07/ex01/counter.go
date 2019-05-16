package main

import (
	"bufio"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	for tmp := p; ; {
		advance, token, err := bufio.ScanWords(tmp, false)
		if err != nil {
			return 0, err
		}

		*c++

		if len(token) == 0 {
			break
		}

		tmp = tmp[advance:]
	}

	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	for tmp := p; ; {
		advance, token, err := bufio.ScanLines(tmp, false)
		if err != nil {
			return 0, err
		}

		*c++

		if len(token) == 0 {
			break
		}

		tmp = tmp[advance:]
	}

	return len(p), nil
}
