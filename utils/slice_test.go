package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceContains(t *testing.T) {
	assert.Equal(t, true, SliceContains([]int{1,2,3}, 1))
	assert.Equal(t, true, SliceContains([]int{1,2,3}, 2))
	assert.Equal(t, true, SliceContains([]int{1,2,3}, 3))
	assert.Equal(t, false, SliceContains([]int{1,2,3}, 4))
}
