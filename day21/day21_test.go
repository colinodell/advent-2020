package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = []string{
	"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
	"trh fvjkl sbzzf mxmxvkd (contains dairy)",
	"sqjhc fvjkl (contains soy)",
	"sqjhc mxmxvkd sbzzf (contains fish)",
}

var realInput = utils.ReadLines("../day21/input.txt")

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, Part1(sampleInput))
	assert.Equal(t, 2211, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", Part2(sampleInput))
	assert.Equal(t, "vv,nlxsmb,rnbhjk,bvnkk,ttxvphb,qmkz,trmzkcfg,jpvz", Part2(realInput))
}
