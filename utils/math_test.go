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
