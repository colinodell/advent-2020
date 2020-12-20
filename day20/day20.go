package main

import (
	"advent-2020/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile("./day20/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Product of corner tile IDs: %d\n", Part1(input))
}

func Part1(input string) int {
	tiles := ParseInput(input)

	allEdges := make(map[string]int)
	for _, tile := range tiles {
		for _, edge := range tile {
			allEdges[edge]++
		}
	}

	result := 1

	for id, tile := range tiles {
		cnt := 0
		for _, edge := range tile {
			if x := allEdges[edge]; x > 1 {
				cnt++
			}
		}

		// Corner tiles will only align with exactly "2" other tiles
		// (but we have two variants of each edge, so we're searching for "4")
		if cnt == 4 {
			result *= id
		}
	}

	return result
}

func ParseInput (input string) map[int][]string {
	tiles := strings.Split(input, "\n\n")
	result := make(map[int][]string, len(tiles))

	for _, tile := range tiles {
		s := strings.SplitN(tile, "\n", 2)
		var id int
		fmt.Sscanf(s[0], "Tile %d:", &id)
		result[id] = ParseTile(s[1])
	}

	return result
}

func ParseTile (input string) []string {
	result := make([]string, 8)

	lines := strings.Split(input, "\n")

	result[0] = lines[0]
	result[1] = reverseString(lines[0])

	result[4] = lines[len(lines)-1]
	result[5] = reverseString(result[4])

	var left, right strings.Builder
	for i := 0; i < len(lines[0]); i++ {
		left.WriteRune(rune(lines[i][0]))
		right.WriteRune(rune(lines[i][len(lines[i])-1]))
	}

	result[2] = right.String()
	result[3] = reverseString(result[2])

	result[6] = left.String()
	result[7] = reverseString(result[6])

	return result
}

func reverseString(input string) string {
	var sb strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}
