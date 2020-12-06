package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumCountsIncorrectly(t *testing.T) {
	sampleInput := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	assert.Equal(t, 11, SumCountsIncorrectly(sampleInput))

	realInput := utils.ReadFile("../day06/input.txt")
	assert.Equal(t, 6297, SumCountsIncorrectly(realInput))
}

func TestSumCountsCorrectly(t *testing.T) {
	sampleInput := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	assert.Equal(t, 6, SumCountsCorrectly(sampleInput))

	realInput := utils.ReadFile("../day06/input.txt")
	assert.Equal(t, 3158, SumCountsCorrectly(realInput))
}
