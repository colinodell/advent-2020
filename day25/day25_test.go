package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var samplePublicKeys = []int{5764801, 17807724}
var realPublicKeys = []int{18499292, 8790390}

func TestPart1(t *testing.T) {
	assert.Equal(t, 14897079, Part1(samplePublicKeys[0], samplePublicKeys[1]))
	assert.Equal(t, 18433997, Part1(realPublicKeys[0], realPublicKeys[1]))
}
