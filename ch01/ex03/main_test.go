package main

import (
	"io/ioutil"
	"testing"
)

func makeArgs(length int) []string {
	args := make([]string, 0, length)
	for i := 0; i < length; i++ {
		args = append(args, "foo")
	}
	return args
}

func BenchmarkEcho1_10args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(ioutil.Discard, makeArgs(10))
	}
}

func BenchmarkEcho1_100args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(ioutil.Discard, makeArgs(100))
	}
}

func BenchmarkEcho3_10args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(ioutil.Discard, makeArgs(10))
	}
}

func BenchmarkEcho3_100args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(ioutil.Discard, makeArgs(100))
	}
}
