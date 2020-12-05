package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoardingPassToId(t *testing.T) {
	assert.Equal(t, 357, BoardingPassToId("FBFBBFFRLR"))
	assert.Equal(t, 567, BoardingPassToId("BFFFBBFRRR"))
	assert.Equal(t, 119, BoardingPassToId("FFFBBBFRRR"))
	assert.Equal(t, 820, BoardingPassToId("BBFFBBFRLL"))
}

func TestFindHighestSeatId(t *testing.T) {
	assert.Equal(t, 820, FindHighestSeatId([]string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}))
}

func TestDay05(t *testing.T) {
	realInput := utils.ReadLines("../day05/input.txt")

	assert.Equal(t, 850, FindHighestSeatId(realInput))
	assert.Equal(t, 599, FindEmptySeat(realInput))
}
