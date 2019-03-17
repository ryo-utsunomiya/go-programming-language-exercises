package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount1(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))],
	)
}

func PopCount2(x uint64) int {
	var result byte
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>(uint(i)*8))]
	}
	return int(result)
}

func PopCount3(x uint64) int {
	var cnt int
	for i := 0; i < 64; i++ {
		cnt += int((x >> uint64(i)) & 1)
	}
	return cnt
}

func PopCount4(x uint64) int {
	var cnt int
	for x > 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}
