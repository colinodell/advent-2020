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

func TestPocketDimension_Run(t *testing.T) {
	sampleDimension := NewPocketDimension(sampleInput)
	assert.Equal(t, 112, sampleDimension.Run(6))

	realDimension := NewPocketDimension(realInput)
	assert.Equal(t, 319, realDimension.Run(6))
}
