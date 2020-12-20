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

type TileSet map[int]Tile

func (ts TileSet) CountAllPossibleEdges() map[string]int {
	ret := make(map[string]int)

	for _, t := range ts {
		for _, edge := range t.AllPossibleEdges() {
			ret[edge]++
		}
	}

	return ret
}

type Tile struct {
	num int
	grid [][]rune
	size int
}

func NewTile(input string) Tile {
	s := strings.SplitN(input, "\n", 2)
	var id int
	fmt.Sscanf(s[0], "Tile %d:", &id)

	lines := strings.Split(s[1], "\n")
	size := len(lines)

	t := Tile{num: id, grid: make([][]rune, size), size: size}
	for i, line := range lines {
		t.grid[i] = make([]rune, size)
		for j, char := range line {
			t.grid[i][j] = char
		}
	}

	return t
}

func (t *Tile) Top() string {
	return string(t.grid[0])
}

func (t *Tile) Bottom() string {
	return string(t.grid[t.size-1])
}

func (t *Tile) Left() string {
	var left strings.Builder
	for i := 0; i < len(t.grid[0]); i++ {
		left.WriteRune(t.grid[i][0])
	}

	return left.String()
}

func (t *Tile) Right() string {
	var right strings.Builder
	for i := 0; i < t.size; i++ {
		right.WriteRune(t.grid[i][t.size-1])
	}

	return right.String()
}

func (t *Tile) Edges() []string {
	return []string{
		t.Top(),
		t.Bottom(),
		t.Left(),
		t.Right(),
	}
}

func (t *Tile) AllPossibleEdges() []string {
	edges := t.Edges()

	return []string{
		edges[0],
		reverseString(edges[0]),
		edges[1],
		reverseString(edges[1]),
		edges[2],
		reverseString(edges[2]),
		edges[3],
		reverseString(edges[3]),
	}
}

func (t Tile) CountUniqueEdges(edges map[string]int) int {
	ret := 0
	for _, edge := range t.AllPossibleEdges() {
		if edges[edge] == 2 {
			ret++
		}
	}

	return ret
}

func Part1(input string) int {
	tiles := NewTileSet(input)
	allEdges := tiles.CountAllPossibleEdges()

	result := 1

	for id, tile := range tiles {
		if tile.CountUniqueEdges(allEdges) == 4 {
			result *= id
		}
	}

	return result
}

func NewTileSet(input string) TileSet {
	tiles := strings.Split(input, "\n\n")
	result := make(TileSet, len(tiles))

	for _, tileData := range tiles {
		t := NewTile(tileData)
		result[t.num] = t
	}

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
