package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = []string{
	".#.",
	"..#",
	"###",
}

var realInput = utils.ReadLines("../day17/input.txt")

func TestPocketDimension_Run3D(t *testing.T) {
	sampleDimension := NewPocketDimension(sampleInput, 3)
	assert.Equal(t, 112, sampleDimension.Run(6))

	realDimension := NewPocketDimension(realInput, 3)
	assert.Equal(t, 319, realDimension.Run(6))
}

func TestPocketDimension_Run4D(t *testing.T) {
	sampleDimension := NewPocketDimension(sampleInput, 4)
	assert.Equal(t, 848, sampleDimension.Run(6))

	realDimension := NewPocketDimension(realInput, 4)
	assert.Equal(t, 2324, realDimension.Run(6))
}
