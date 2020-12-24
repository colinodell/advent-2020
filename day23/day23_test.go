package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleInput = "389125467"
var realInput = "739862541"

func TestPlay(t *testing.T) {
	assert.Equal(t, "92658374", Play(sampleInput, 10))
	assert.Equal(t, "67384529", Play(sampleInput, 100))
	assert.Equal(t, "94238657", Play(realInput, 100))
}
