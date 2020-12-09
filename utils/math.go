package utils

import (
	"math/bits"
)

// From https://stackoverflow.com/a/54421330/158766
var MinInt = (1 << bits.UintSize) / -2
var MaxInt = (1 << bits.UintSize) / 2 - 1

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MinIntSlice(s []int) int {
	if len(s) == 0 {
		panic("empty slice")
	}

	min := MaxInt

	for _, n := range s {
		if n < min {
			min = n
		}
	}

	return min
}

func MaxIntSlice(s []int) int {
	if len(s) == 0 {
		panic("empty slice")
	}

	max := MinInt

	for _, n := range s {
		if n > max {
			max = n
		}
	}

	return max
}

func SumSlice(seq []int) int {
	sum := 0
	for _, n := range seq {
		sum += n
	}

	return sum
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
