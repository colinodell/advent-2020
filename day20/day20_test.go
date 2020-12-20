package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = utils.ReadFile("../day20/sample_input.txt")
var realInput = utils.ReadFile("../day20/input.txt")

func TestParseTile(t *testing.T) {
	tile := "..##.#..#.\n##..#.....\n#...##..#.\n####.#...#\n##.##.###.\n##...#.###\n.#.#.#..##\n..#....#..\n###...#.#.\n..###..###"
	expected := []string{
		"..##.#..#.",
		".#..#.##..",
		"...#.##..#",
		"#..##.#...",
		"..###..###",
		"###..###..",
		".#####..#.",
		".#..#####.",
	}

	assert.Equal(t, expected, ParseTile(tile))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 20899048083289, Part1(sampleInput))
	assert.Equal(t, 27803643063307, Part1(realInput))
}
