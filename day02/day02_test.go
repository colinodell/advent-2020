package main

import (
	"advent-2020/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	sampleInput := strings.Split("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc", "\n")
	entries := parseEntries(sampleInput)

	assert.Equal(t, 2, countCorrect(entries, Entry.IsValidForOldPolicy))
	assert.Equal(t, 1, countCorrect(entries, Entry.IsValidForNewPolicy))

	realInput := utils.ReadLines("../day02/input.txt")
	entries = parseEntries(realInput)

	assert.Equal(t, 439, countCorrect(entries, Entry.IsValidForOldPolicy))
	assert.Equal(t, 584, countCorrect(entries, Entry.IsValidForNewPolicy))
}
