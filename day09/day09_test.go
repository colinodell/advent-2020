package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleNumbers = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
var realNumbers = utils.ReadLinesOfNumbers("../day09/input.txt")

func TestFindFirstInvalidNumber(t *testing.T) {
	assert.Equal(t, 127, FindFirstInvalidNumber(sampleNumbers, 5))
	assert.Equal(t, 32321523, FindFirstInvalidNumber(realNumbers, 25))
}

func TestFindEncryptionWeakness(t *testing.T) {
	assert.Equal(t, 62, FindEncryptionWeakness(sampleNumbers, 5))
	assert.Equal(t, 4794981, FindEncryptionWeakness(realNumbers, 25))
}
