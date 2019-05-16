package main

import "testing"

func TestWordCounter_Write(t *testing.T) {
	var c WordCounter
	c.Write([]byte("hello world"))
	if c != 2 {
		t.Errorf("want 2, got %d", c)
	}
}

func TestLineCounter_Write(t *testing.T) {
	var c LineCounter
	c.Write([]byte("foo\nbar\r\nbaz"))
	if c != 3 {
		t.Errorf("want 3, got %d", c)
	}
}
