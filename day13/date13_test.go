package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = []string{"939", "7,13,x,x,59,x,31,19"}
var realInput = utils.ReadLines("../day13/input.txt")

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 295, SolvePart1(sampleInput))
	assert.Equal(t, 2238, SolvePart1(realInput))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3417, SolvePart2("17,x,13,19"))
	assert.Equal(t, 754018, SolvePart2("67,7,59,61"))
	assert.Equal(t, 779210, SolvePart2("67,x,7,59,61"))
	assert.Equal(t, 1261476, SolvePart2("67,7,x,59,61"))
	assert.Equal(t, 1202161486, SolvePart2("1789,37,47,1889"))
	assert.Equal(t, 560214575859998, SolvePart2(realInput[1]))
}
