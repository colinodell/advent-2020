package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = utils.ReadLines("../day24/sample_input.txt")
var realInput = utils.ReadLines("../day24/input.txt")

func TestPart1(t *testing.T) {
	assert.Equal(t, 10, Part1(sampleInput))
	assert.Equal(t, 300, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2208, Part2(sampleInput, 100))
	assert.Equal(t, 3466, Part2(realInput, 100))
}
