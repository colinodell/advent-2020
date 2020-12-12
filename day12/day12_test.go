package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = []string{"F10", "N3", "F7", "R90", "F11"}
var realInput = utils.ReadLines("../day12/input.txt")

func TestNewShipWithBadNavigator(t *testing.T) {
	sampleShip := NewShipWithBadNavigator()
	assert.Equal(t, 25, sampleShip.Follow(sampleInput))

	realShip := NewShipWithBadNavigator()
	assert.Equal(t, 882, realShip.Follow(realInput))
}

func TestNewShipWithGoodNavigator(t *testing.T) {
	sampleShip := NewShipWithGoodNavigator()
	assert.Equal(t, 286, sampleShip.Follow(sampleInput))

	realShip := NewShipWithGoodNavigator()
	assert.Equal(t, 28885, realShip.Follow(realInput))
}
