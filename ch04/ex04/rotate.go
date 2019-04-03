package main

func Rotate(s []int, n int) {
	for n > len(s) {
		n = n % len(s)
	}
	if n == 0 {
		return
	}

	tmp := s[n:]
	tmp = append(tmp, s[:n]...)
	copy(s, tmp)
}
