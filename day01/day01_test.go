package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01(t *testing.T) {
	assert.Equal(t, 514579, findAndMultiplyTwoMatchingEntries([]int{1721,979,366,299,675,1456}))
	assert.Equal(t, 241861950, findAndMultiplyThreeMatchingEntries([]int{1721,979,366,299,675,1456}))
}
