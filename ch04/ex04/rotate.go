package main

func Rotate(s []int, n int) []int {
	if n > len(s) {
		n = n % len(s)
	}
	if n == 0 {
		return s
	}

	s = append(s, s[:n]...)
	return s[n:]
}
