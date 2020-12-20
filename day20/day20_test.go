package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = utils.ReadFile("../day20/sample_input.txt")
var realInput = utils.ReadFile("../day20/input.txt")

func TestPart1(t *testing.T) {
	assert.Equal(t, 20899048083289, Part1(sampleInput))
	assert.Equal(t, 27803643063307, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 273, Part2(sampleInput))
	assert.Equal(t, 1644, Part2(realInput))
}

func TestTile_Rotate(t *testing.T) {
	tile := NewTile("Tile 1:\n123\n456\n789")

	rotated := tile.Rotate()
	assert.Equal(t, "741\n852\n963", rotated.String())

	rotated = rotated.Rotate()
	assert.Equal(t, "987\n654\n321", rotated.String())

	rotated = rotated.Rotate()
	assert.Equal(t, "369\n258\n147", rotated.String())

	rotated = rotated.Rotate()
	assert.Equal(t, "123\n456\n789", rotated.String())
}

func TestTile_FlipTopBottom(t *testing.T) {
	tile := NewTile("Tile 1:\n123\n456\n789")

	flipped := tile.FlipTopBottom()
	assert.Equal(t, "789\n456\n123", flipped.String())

	flipped = flipped.FlipTopBottom()
	assert.Equal(t, "123\n456\n789", flipped.String())
}

func TestTile_FlipLeftRight(t *testing.T) {
	tile := NewTile("Tile 1:\n123\n456\n789")

	flipped := tile.FlipLeftRight()
	assert.Equal(t, "321\n654\n987", flipped.String())

	flipped = flipped.FlipLeftRight()
	assert.Equal(t, "123\n456\n789", flipped.String())
}

func TestTile_WithoutBorder(t *testing.T) {
	tile1 := NewTile("Tile 1:\n123\n456\n789")

	withoutBorder1 := tile1.WithoutBorder()
	assert.Equal(t, "5", withoutBorder1.String())

	tile2 := NewTile("Tile 2:\nabcd\nefgh\nijkl\nmnop")

	withoutBorder2 := tile2.WithoutBorder()
	assert.Equal(t, "fg\njk", withoutBorder2.String())
}

func TestTile_CountUniqueEdges(t *testing.T) {
	ts := NewTileSet(sampleInput)
	edges := ts.CountAllPossibleEdges()

	assert.Equal(t, 4, ts[1951].CountUniqueEdges(edges))
	assert.Equal(t, 6, ts[2311].CountUniqueEdges(edges))
	assert.Equal(t, 4, ts[3079].CountUniqueEdges(edges))
	assert.Equal(t, 6, ts[2729].CountUniqueEdges(edges))
	assert.Equal(t, 8, ts[1427].CountUniqueEdges(edges))
	assert.Equal(t, 6, ts[2473].CountUniqueEdges(edges))
	assert.Equal(t, 4, ts[2971].CountUniqueEdges(edges))
	assert.Equal(t, 6, ts[1489].CountUniqueEdges(edges))
	assert.Equal(t, 4, ts[1171].CountUniqueEdges(edges))
}

func TestTileSet_FindRandomCorner(t *testing.T) {
	ts := NewTileSet(sampleInput)
	edges := ts.CountAllPossibleEdges()

	corner := ts.FindRandomCorner(edges)

	assert.Contains(t, []int{1951, 3079, 2971, 1171}, corner.num)
}
