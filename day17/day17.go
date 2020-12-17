package main

import (
	"advent-2020/utils"
	"fmt"
)

func main() {
	input := utils.ReadLines("./day17/input.txt")

	fmt.Println("----- Part 1 -----")
	pd3 := NewPocketDimension(input, 3)
	fmt.Printf("3D dimension, after 6 cycles: %d\n", pd3.Run(6))

	fmt.Println("----- Part 2 -----")
	pd4 := NewPocketDimension(input, 4)
	fmt.Printf("4D dimension, after 6 cycles: %d\n", pd4.Run(6))
}

type PocketDimension struct {
	grid map[utils.Vector]bool
}

func NewPocketDimension(input []string, dimensions int) PocketDimension {
	pd := PocketDimension{
		grid: make(map[utils.Vector]bool),
	}

	for y, line := range input {
		for x, char := range line {
			if char != '#' {
				continue
			}

			if dimensions == 3 {
				pd.grid[utils.Vector3{X: x, Y: y, Z: 0}] = true
			} else if dimensions == 4 {
				pd.grid[utils.Vector4{X: x, Y: y, Z: 0, W: 0}] = true
			} else {
				panic("invalid number of dimensions")
			}
		}
	}

	return pd
}

func (pd *PocketDimension) RunCycle() {
	searchGrid := make(map[utils.Vector]bool)
	for k, v := range pd.grid {
		searchGrid[k] = v

		// Make sure we also check its neighbors
		for _, neighbor := range k.Nearby() {
			if _, ok := searchGrid[neighbor]; !ok {
				searchGrid[neighbor] = false
			}
		}
	}

	newGrid := make(map[utils.Vector]bool)

	for vector, active := range searchGrid {
		activeNeighbors := 0
		for _, neighbor := range vector.Nearby() {
			if searchGrid[neighbor] {
				activeNeighbors++
			}
		}

		if active && (activeNeighbors == 2 || activeNeighbors == 3) {
			newGrid[vector] = true
		} else if !active && activeNeighbors == 3 {
			newGrid[vector] = true
		}
	}

	pd.grid = newGrid
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
