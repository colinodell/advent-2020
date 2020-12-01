package main

import (
	"advent-2020/utils"
	"fmt"
)

func main() {
	entries := utils.ReadLinesOfNumbers("./day01/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("The two entries multiply out to: %d\n", findAndMultiplyTwoMatchingEntries(entries))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("The three entries multiply out to: %d\n", findAndMultiplyThreeMatchingEntries(entries))
}

func findAndMultiplyTwoMatchingEntries(entries []int) int {
	for _, i := range entries {
		for _, j := range entries {
			if i+j == 2020 {
				return i * j
			}
		}
	}

	return -1
}

func findAndMultiplyThreeMatchingEntries(entries []int) int {
	for _, i := range entries {
		for _, j := range entries {
			for _, k := range entries {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}

	return -1
}
