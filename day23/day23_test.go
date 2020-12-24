package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = "389125467"
var realInput = "739862541"

func TestCupGame(t *testing.T) {
	{
		game := NewCupGame(sampleInput)
		game.Play(10)
		assert.Equal(t, "92658374", game.LabelsAfterCup1())
	}

	{
		game := NewCupGame(sampleInput)
		game.Play(100)
		assert.Equal(t, "67384529", game.LabelsAfterCup1())
	}

	{
		game := NewCupGame(realInput)
		game.Play(100)
		assert.Equal(t, "94238657", game.LabelsAfterCup1())
	}
}
