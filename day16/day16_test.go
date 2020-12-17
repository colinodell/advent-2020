package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = utils.ReadFile("../day16/sample_input.txt")
var realInput = utils.ReadFile("../day16/input.txt")

func TestPuzzle_CalculateErrorRate(t *testing.T) {
	samplePuzzle := NewPuzzle(sampleInput)
	assert.Equal(t, 71, samplePuzzle.CalculateErrorRate())

	realPuzzle := NewPuzzle(realInput)
	assert.Equal(t, 30869, realPuzzle.CalculateErrorRate())
}

func TestPuzzle_MultiplyMyDepartureFields(t *testing.T) {
	realPuzzle := NewPuzzle(realInput)
	assert.Equal(t, 4381476149273, realPuzzle.MultiplyMyDepartureFields())
}
