package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	startingNumbers []int
	turns int
	expected int
}

var part1Tests = []TestData{
	// {startingNumbers: []int{0, 3, 6}, turns: 4, expected: 0},
	// {startingNumbers: []int{0, 3, 6}, turns: 5, expected: 3},
	// {startingNumbers: []int{0, 3, 6}, turns: 6, expected: 3},
	// {startingNumbers: []int{0, 3, 6}, turns: 7, expected: 1},
	// {startingNumbers: []int{0, 3, 6}, turns: 8, expected: 0},
	// {startingNumbers: []int{0, 3, 6}, turns: 9, expected: 4},
	// {startingNumbers: []int{0, 3, 6}, turns: 10, expected: 0},

	// {startingNumbers: []int{0, 3, 6}, turns: 2020, expected: 436},
	// {startingNumbers: []int{1, 3, 2}, turns: 2020, expected: 1},
	// {startingNumbers: []int{2, 1, 3}, turns: 2020, expected: 10},
	// {startingNumbers: []int{1, 2, 3}, turns: 2020, expected: 27},
	// {startingNumbers: []int{2, 3, 1}, turns: 2020, expected: 78},
	// {startingNumbers: []int{3, 2, 1}, turns: 2020, expected: 438},
	// {startingNumbers: []int{3, 1, 2}, turns: 2020, expected: 1836},

	{startingNumbers: []int{0, 8, 15, 2, 12, 1, 4}, turns: 2020, expected: 289},

	// {startingNumbers: []int{0, 3, 6}, turns: 30000000, expected: 175594},
	// {startingNumbers: []int{1, 3, 2}, turns: 30000000, expected: 2578},
	// {startingNumbers: []int{2, 1, 3}, turns: 30000000, expected: 3544142},
	// {startingNumbers: []int{1, 2, 3}, turns: 30000000, expected: 261214},
	// {startingNumbers: []int{2, 3, 1}, turns: 30000000, expected: 6895259},
	// {startingNumbers: []int{3, 2, 1}, turns: 30000000, expected: 18},
	// {startingNumbers: []int{3, 1, 2}, turns: 30000000, expected: 362},

	{startingNumbers: []int{0, 8, 15, 2, 12, 1, 4}, turns: 30000000, expected: 1505722},
}

func TestPlay(t *testing.T) {
	for _, test := range part1Tests {
		assert.Equal(t, test.expected, Play(test.startingNumbers, test.turns))
	}
}
