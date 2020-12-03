package main

import (
	"advent-2020/utils"
	"fmt"
)

func main() {
	grid := utils.ReadLines("./day03/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of trees found on a 3,1 slope: %d\n", countTrees(grid, utils.Vector2{X: 3, Y:1}))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Product of slopes: %d\n", productTrees(grid, []utils.Vector2{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}))
}

func countTrees (grid []string, slope utils.Vector2) int {
	gridWidth := len(grid[0])
	row, column, treesFound := 0, 0, 0

	for {
		row += slope.Y
		column += slope.X

		if row >= len(grid) {
			break
		}

		if grid[row][column % gridWidth] == '#' {
			treesFound++
		}
	}

	return treesFound
}

func productTrees(grid []string, slopes []utils.Vector2) int {
	product := 1

	for _, slope := range slopes {
		product *= countTrees(grid, slope)
	}

	return product
}
