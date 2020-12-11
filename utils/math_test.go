package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 42, Max(3, 42))
	assert.Equal(t, 42, Max(42, 3))
}

func TestMinIntSlice(t *testing.T) {
	assert.Equal(t, -55, MinIntSlice([]int{3, 42, 7, -55}))
	assert.Equal(t, -3, MinIntSlice([]int{-1, -2, -3}))
}

func TestMaxIntSlice(t *testing.T) {
	assert.Equal(t, 42, MaxIntSlice([]int{3, 42, 7, -55}))
	assert.Equal(t, -1, MaxIntSlice([]int{-1, -2, -3}))
}

func TestTribonacci(t *testing.T) {
	assert.Equal(t, 0, Tribonacci(0))
	assert.Equal(t, 0, Tribonacci(1))
	assert.Equal(t, 1, Tribonacci(2))
	assert.Equal(t, 1, Tribonacci(3))
	assert.Equal(t, 2, Tribonacci(4))
	assert.Equal(t, 4, Tribonacci(5))
	assert.Equal(t, 7, Tribonacci(6))
	assert.Equal(t, 13, Tribonacci(7))
	assert.Equal(t, 24, Tribonacci(8))
	assert.Equal(t, 44, Tribonacci(9))
	assert.Equal(t, 81, Tribonacci(10))
}
