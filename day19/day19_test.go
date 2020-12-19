package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput1 = utils.ReadFile("../day19/sample_input.txt")
var sampleInput2 = utils.ReadFile("../day19/sample_input2.txt")
var realInput = utils.ReadFile("../day19/input.txt")

func TestParseRegex(t *testing.T) {
	assert.Equal(t, "^(a(ab|ba))$", ParseRegex([]string{
		"0: 1 2",
		"1: \"a\"",
		"2: 1 3 | 3 1",
		"3: \"b\"",
	}))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, Part1(sampleInput1))
	assert.Equal(t, 3, Part1(sampleInput2))
	assert.Equal(t, 233, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, Part2(sampleInput2))
	assert.Equal(t, 396, Part2(realInput))
}
