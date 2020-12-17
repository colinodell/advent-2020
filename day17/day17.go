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
			var v utils.Vector
			if dimensions == 3 {
				v = utils.Vector3{X: x, Y: y, Z: 0}
			} else if dimensions == 4 {
				v = utils.Vector4{X: x, Y: y, Z: 0, W: 0}
			} else {
				panic("invalid number of dimensions")
			}

			pd.grid[v] = char == '#'
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
