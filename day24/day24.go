package main

import (
	"advent-2020/utils"
)

var offsets = map[string]utils.VectorHexagonal{
	"e": {Q: 1, R: 0},
	"se": {Q: 0, R: 1},
	"sw": {Q: -1, R: 1},
	"w": {Q: -1, R: 0},
	"nw": {Q: 0, R: -1},
	"ne": {Q: 1, R: -1},
}

func parseTiles(input []string) map[utils.VectorHexagonal]struct{} {
	black := make(map[utils.VectorHexagonal]struct{})

	for _, line := range input {
		position := utils.VectorHexagonal{}

		for i := 0; i < len(line); i++ {
			direction := string(line[i])
			if direction == "n" || direction == "s" {
				direction += string(line[i+1])
				i++
			}

			position = position.Add(offsets[direction])
		}

		if _, ok := black[position]; ok {
			delete(black, position)
		} else {
			black[position] = struct{}{}
		}
	}

	return black
}

func Part1(input []string) int {
	return len(parseTiles(input))
}

func Part2(input []string, rounds int) int {
	blackTiles := parseTiles(input)

	for rounds > 0 {
		searchGrid := make(map[utils.VectorHexagonal]int)
		for pos := range blackTiles {
			for _, adjacentOffset := range offsets {
				searchGrid[pos.Add(adjacentOffset)]++
			}
		}

		newTiles := make(map[utils.VectorHexagonal]struct{})

		for pos, blackAdjacentCount := range searchGrid {
			if _, isBlack := blackTiles[pos]; (isBlack && blackAdjacentCount == 1) || blackAdjacentCount == 2 {
				newTiles[pos] = struct{}{}
			}
		}

		blackTiles = newTiles
		rounds--
	}

	return len(blackTiles)
}
