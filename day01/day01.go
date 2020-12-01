package main

import (
	"advent-2020/utils"
	"fmt"
)

func main() {
	entries := utils.ReadLinesOfNumbers("./day01/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("The two entries multiply out to: %d\n", findMatchingEntryProduct(entries, 2, 2020))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("The three entries multiply out to: %d\n", findMatchingEntryProduct(entries, 3, 2020))
}

func findMatchingEntryProduct(numbers []int, count int, target int) int {
	if count == 0 {
		// We've reached the desired count; return 1 to bubble the solution back up
		return 1
	}

	if count == 1 {
		if utils.SliceContains(numbers, target) {
			// A matching number was found!
			return target
		} else {
			// We needed to find a specific number but did not
			// Return 0 so the product becomes 0 and we can check the next potential solution
			return 0
		}
	}

	for _, i := range numbers {
		// Try to find (count-1) numbers that add up to (target-i)
		product := findMatchingEntryProduct(numbers, count - 1, target - i)
		if product > 0 {
			// Such a number was found - bubble the product back up
			return i * product
		}
	}

	return 0
}

