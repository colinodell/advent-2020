package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput1 = []string {
	"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
	"mem[8] = 11",
	"mem[7] = 101",
	"mem[8] = 0",
}

var sampleInput2 = []string {
	"mask = 000000000000000000000000000000X1001X",
	"mem[42] = 100",
	"mask = 00000000000000000000000000000000X0XX",
	"mem[26] = 1",
}

var realInput = utils.ReadLines("../day14/input.txt")

func TestComputer_RunPart1(t *testing.T) {
	sampleComputer := NewComputer(&Part1{})
	assert.Equal(t, uint64(165), sampleComputer.Run(sampleInput1))

	realComputer := NewComputer(&Part1{})
	assert.Equal(t, uint64(15172047086292), realComputer.Run(realInput))
}

func TestComputer_RunPart2(t *testing.T) {
	sampleComputer := NewComputer(&Part2{})
	assert.Equal(t, uint64(208), sampleComputer.Run(sampleInput2))

	realComputer := NewComputer(&Part2{})
	assert.Equal(t, uint64(4197941339968), realComputer.Run(realInput))
}
