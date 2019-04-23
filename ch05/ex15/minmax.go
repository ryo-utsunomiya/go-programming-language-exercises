package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(min0(1, 2, 3))
	fmt.Println(max0(1, 2, 3))

	fmt.Println(min0())
	fmt.Println(max0())

	fmt.Println(min1(1, 2, 3))
	fmt.Println(max1(1, 2, 3))
}

func min0(nums ...int) int {
	result := math.MaxInt32

	for _, n := range nums {
		if n < result {
			result = n
		}
	}

	return result
}

func min1(n int, nums ...int) int {
	result := n

	for _, n = range nums {
		if n < result {
			result = n
		}
	}

	return result
}

func max0(nums ...int) int {
	result := math.MinInt32

	for _, n := range nums {
		if n > result {
			result = n
		}
	}

	return result
}

func max1(n int, nums ...int) int {
	result := n

	for _, n := range nums {
		if n > result {
			result = n
		}
	}

	return result
}
