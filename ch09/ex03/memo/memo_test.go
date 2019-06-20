// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch09/ex03/memo"
	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch09/ex03/memo/memotest"
)

var httpGetBody = func(key string, done <-chan struct{}) (interface{}, error) {
	return memotest.HTTPGetBody(key)
}

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
