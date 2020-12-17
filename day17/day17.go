package main

import (
	"advent-2020/utils"
)

func main() {
	input := utils.ReadLines("./day17/input.txt")

	fmt.Println("----- Part 1 -----")
	pd3 := NewPocketDimension(input, 3)
	fmt.Printf("3D dimension, after 6 cycles: %d\n", pd3.Run(6))
}

type PocketDimension struct {
	grid map[utils.Vector3]bool
}

func NewPocketDimension(input []string) PocketDimension {
	pd := PocketDimension{
		grid: make(map[utils.Vector3]bool),
	}

	for y, line := range input {
		for x, char := range line {
			v := utils.Vector3{X: x, Y: y, Z: 0}
			pd.grid[v] = char == '#'
		}
	}

	return pd
}

func (pd *PocketDimension) RunCycle() {
	searchGrid := make(map[utils.Vector3]bool)
	for k, v := range pd.grid {
		searchGrid[k] = v

		// Make sure we also check its neighbors
		for _, neighbor := range k.Nearby() {
			if _, ok := searchGrid[neighbor]; !ok {
				searchGrid[neighbor] = false
			}
		}
	}

	for vector, active := range searchGrid {
		activeNeighbors := 0
		for _, neighbor := range vector.Nearby() {
			if searchGrid[neighbor] {
				activeNeighbors++
			}
		}

		if active {
			pd.grid[vector] = activeNeighbors == 2 || activeNeighbors == 3
		} else {
			pd.grid[vector] = activeNeighbors == 3
		}
	}
}

func (pd *PocketDimension) countActive() int {
	count := 0

	for _, active := range pd.grid {
		if active {
			count++
		}
	}

	return count
}

func (pd *PocketDimension) Run(cycles int) int {
	for i := 0; i < cycles; i++ {
		pd.RunCycle()
	}

	return pd.countActive()
}
