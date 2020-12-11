package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
var realInput = utils.ReadFile("../day11/input.txt")

func TestSeating_RunUntilStablePart1(t *testing.T) {
	sampleSeating := NewSeating(sampleInput)
	sampleSeating.RunUntilStable(Part1)
	assert.Equal(t, 37, sampleSeating.CountAllOccupiedSeats())

	realSeating := NewSeating(realInput)
	realSeating.RunUntilStable(Part1)
	assert.Equal(t, 2249, realSeating.CountAllOccupiedSeats())
}

func TestSeating_RunUntilStablePart2(t *testing.T) {
	sampleSeating := NewSeating(sampleInput)
	sampleSeating.RunUntilStable(Part2)
	assert.Equal(t, 26, sampleSeating.CountAllOccupiedSeats())

	realSeating := NewSeating(realInput)
	realSeating.RunUntilStable(Part2)
	assert.Equal(t, 2023, realSeating.CountAllOccupiedSeats())
}
