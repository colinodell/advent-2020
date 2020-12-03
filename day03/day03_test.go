package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {
	sampleInput := strings.Split("..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#", "\n")

	assert.Equal(t, 7, countTrees(sampleInput, utils.Vector2{X: 3, Y: 1}))
	assert.Equal(t, 336, productTrees(sampleInput, []utils.Vector2{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}))
	assert.Equal(t, utils.Vector2{X: 1, Y: 5}, findBestSlope(sampleInput))

	realInput := utils.ReadLines("../day03/input.txt")
	assert.Equal(t, 274, countTrees(realInput, utils.Vector2{X: 3, Y: 1}))
	assert.Equal(t, 6050183040, productTrees(realInput, []utils.Vector2{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}))
	assert.Equal(t, utils.Vector2{X: 1, Y: 322}, findBestSlope(realInput))
}
