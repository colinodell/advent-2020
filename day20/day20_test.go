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
