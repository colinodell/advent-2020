package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := []int{1721,979,366,299,675,1456}

	assert.Equal(t, 514579, findMatchingEntryProduct(sampleInput, 2, 2020))
	assert.Equal(t, 241861950, findMatchingEntryProduct(sampleInput, 3, 2020))

	realInput := utils.ReadLinesOfNumbers("../day01/input.txt")

	assert.Equal(t, 1015476, findMatchingEntryProduct(realInput, 2, 2020))
	assert.Equal(t, 200878544, findMatchingEntryProduct(realInput, 3, 2020))
}
