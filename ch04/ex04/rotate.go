package main

func Rotate(s []int, n int) {
	for l := len(s); n > l; n -= l {
	}
	if n == 0 {
		return
	}

	tmp := s[n:]
	tmp = append(tmp, s[:n]...)
	copy(s, tmp)
}
