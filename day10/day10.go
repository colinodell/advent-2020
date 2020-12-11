package main

import (
	"advent-2020/utils"
	"fmt"
	"sort"
)

func main()  {
	joltages := utils.ReadLinesOfNumbers("./day10/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Product of differences: %d\n", CalculateProductOfDifferences(joltages))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Number of distinct arrangements: %d\n", CalculateDistinctArrangements(joltages))
}

func CalculateProductOfDifferences(joltages []int) int {
	sort.Ints(joltages)

	lastJoltage := 0
	differences := map[int]int{1: 0, 2: 0, 3: 1}

	for _, joltage := range joltages {
		differences[joltage - lastJoltage]++
		lastJoltage = joltage
	}

	return differences[1] * differences[3]
}

func CalculateDistinctArrangements(joltages []int) int {
	// Prepend 0; sort joltages; append joltage of final device
	joltages = append([]int{0}, joltages...)
	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages) - 1] + 3)

	seqStart, seqEnd := 0, 0
	result := 1

	for seqStart < len(joltages) - 1 && seqEnd < len(joltages) - 1 {
		// Look for consecutive sequences of numbers with a difference of 1
		for joltages[seqEnd + 1] - joltages[seqEnd] == 1 {
			seqEnd++
		}

		seqLength := seqEnd - seqStart + 1

		if seqLength > 2 {
			// Use Tribonacci numbers to calculate the number of possible combinations
			result *= utils.Tribonacci(seqLength + 1)
		}

		seqStart = seqEnd + 1
		seqEnd = seqStart
	}

	return result
}
