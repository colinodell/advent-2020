package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = "Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"
var infiniteRecursionSampleInput = "Player 1:\n43\n19\n\nPlayer 2:\n2\n29\n14"
var realInput = utils.ReadFile("../day22/input.txt")

func TestGame_Play(t *testing.T) {
	sampleGame := NewGame(sampleInput)
	assert.Equal(t, 306, sampleGame.Play().CalculateScore())

	realGame := NewGame(realInput)
	assert.Equal(t, 32495, realGame.Play().CalculateScore())
}

func TestGame_PlayRecursive(t *testing.T) {
	sampleGame := NewGame(sampleInput)
	assert.Equal(t, 291, sampleGame.PlayRecursive().CalculateScore())

	infiniteRecursionGame := NewGame(infiniteRecursionSampleInput)
	assert.Equal(t, 105, infiniteRecursionGame.PlayRecursive().CalculateScore())

	realGame := NewGame(realInput)
	assert.Equal(t, 32665, realGame.PlayRecursive().CalculateScore())
}
