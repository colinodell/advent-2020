package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var realInput = utils.ReadLines("../day18/input.txt")

func TestEvaluate(t *testing.T) {
	assert.Equal(t, 71, Evaluate("1 + 2 * 3 + 4 * 5 + 6", Part1))
	assert.Equal(t, 51, Evaluate("1 + (2 * 3) + (4 * (5 + 6))", Part1))
	assert.Equal(t, 26, Evaluate("2 * 3 + (4 * 5)", Part1))
	assert.Equal(t, 437, Evaluate("5 + (8 * 3 + 9 + 3 * 4 * 3)", Part1))
	assert.Equal(t, 12240, Evaluate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", Part1))
	assert.Equal(t, 13632, Evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", Part1))

	assert.Equal(t, 231, Evaluate("1 + 2 * 3 + 4 * 5 + 6", Part2))
	assert.Equal(t, 51, Evaluate("1 + (2 * 3) + (4 * (5 + 6))", Part2))
	assert.Equal(t, 46, Evaluate("2 * 3 + (4 * 5)", Part2))
	assert.Equal(t, 1445, Evaluate("5 + (8 * 3 + 9 + 3 * 4 * 3)", Part2))
	assert.Equal(t, 669060, Evaluate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", Part2))
	assert.Equal(t, 23340, Evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", Part2))
}

func TestEvaluateSum(t *testing.T) {
	assert.Equal(t, 5374004645253, EvaluateSum(realInput, Part1))

	assert.Equal(t, 88782789402798, EvaluateSum(realInput, Part2))
}
