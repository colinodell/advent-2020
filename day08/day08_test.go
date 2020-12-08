package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var sampleInput = "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
var realInput = utils.ReadLines("../day08/input.txt")

func TestDay08(t *testing.T) {
	sampleConsole := NewGameConsole(strings.Split(sampleInput, "\n"))
	assert.Equal(t, InfiniteLoop, sampleConsole.Run())
	assert.Equal(t, 5, sampleConsole.Accumulator)

	realConsole := NewGameConsole(realInput)
	assert.Equal(t, InfiniteLoop, realConsole.Run())
	assert.Equal(t, 1723, realConsole.Accumulator)

	assert.Equal(t, 8, sampleConsole.FixProgram())

	assert.Equal(t, 846, realConsole.FixProgram())
}
