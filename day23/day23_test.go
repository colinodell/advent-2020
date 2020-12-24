package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = "389125467"
var realInput = "739862541"

func TestCupGame(t *testing.T) {
	{
		game := NewCupGame(sampleInput, 9)
		game.Play(10)
		assert.Equal(t, "92658374", game.LabelsAfterCup1())
	}

	{
		game := NewCupGame(sampleInput, 9)
		game.Play(100)
		assert.Equal(t, "67384529", game.LabelsAfterCup1())
	}

	{
		game := NewCupGame(realInput, 9)
		game.Play(100)
		assert.Equal(t, "94238657", game.LabelsAfterCup1())
	}

	{
		game := NewCupGame(sampleInput, 1000000)
		game.Play(10000000)
		assert.Equal(t, 149245887792, game.GetProductOfStarLabels())
	}

	{
		game := NewCupGame(realInput, 1000000)
		game.Play(10000000)
		assert.Equal(t, 3072905352, game.GetProductOfStarLabels())
	}
}
