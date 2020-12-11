package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput1 = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
var sampleInput2 = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
var realInput = utils.ReadLinesOfNumbers("../day10/input.txt")

func TestCalculateProductOfDifferences(t *testing.T) {
	assert.Equal(t, 35, CalculateProductOfDifferences(sampleInput1))
	assert.Equal(t, 220, CalculateProductOfDifferences(sampleInput2))
	assert.Equal(t, 2263, CalculateProductOfDifferences(realInput))
}

func TestCalculateDistinctArrangements(t *testing.T) {
	assert.Equal(t, 8, CalculateDistinctArrangements(sampleInput1))
	assert.Equal(t, 19208, CalculateDistinctArrangements(sampleInput2))
	assert.Equal(t, 396857386627072, CalculateDistinctArrangements(realInput))
}
