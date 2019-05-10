package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

var div int

func init() {
	div = 32 << (^uint(0) >> 63)
	fmt.Println(div)
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/div, uint(x%div)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/div, uint(x%div)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	result := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result++
			}
		}
	}
	return result
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: s.words}
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) Elems() []uint {
	result := make([]uint, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, uint(64*i+j))
			}
		}
	}
	return result
}

func main() {
	var x, y IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	//fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	//fmt.Println(y.String())

	fmt.Println(x.Elems(), y.Elems())
}
