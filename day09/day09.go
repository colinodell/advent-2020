package main

import (
	"advent-2020/utils"
	"fmt"
)

func main()  {
	numbers := utils.ReadLinesOfNumbers("./day09/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("First invalid number: %d\n", FindFirstInvalidNumber(numbers, 25))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Encryption weakness: %d\n", FindEncryptionWeakness(numbers, 25))
}

func FindFirstInvalidNumber(numbers []int, preambleSize int) int {
	// Starting after the preamble, check each number to see if it's valid.
	for i := preambleSize; i < len(numbers); i++ {
		if !validate(numbers[i], numbers[i-preambleSize:i]) {
			// Found the invalid number!
			return numbers[i]
		}
	}

	panic("no invalid numbers")
}

func validate(number int, previousNumbers []int) bool {
	// Loop through each possible pair of previous numbers using the two loops below.
	for i := 0; i < len(previousNumbers); i++ {
		for j := 0; j < len(previousNumbers); j++ {
			// if i == j we're not really looking at a pair, it's the same number referenced twice, so skip that.
			// Otherwise see if the pair adds up to the desired number.
			if (i != j) && (previousNumbers[i] + previousNumbers[j] == number) {
				return true
			}
		}
	}

	return false
}

func FindEncryptionWeakness(numbers []int, preambleSize int) int {
	invalidNumber := FindFirstInvalidNumber(numbers, preambleSize)

	// Starting from the beginning of the list, check if we have a sequence of numbers that add up to invalidNumber.
	for i := 0; i < len(numbers); i++ {
		// Start with the first 2 numbers.
		// If we don't find a matching sum, the loop will include yet another number in the potential sequence.
		for cnt := 2; i + cnt < len(numbers); cnt++ {
			// Pull those 2+ numbers into a slice and sum them.
			seq := numbers[i:i+cnt]
			sum := utils.SumSlice(seq)
			if sum == invalidNumber {
				// We found the right contiguous set!
				// Return the min and max of that set as specified by the puzzle.
				return utils.MinIntSlice(seq) + utils.MaxIntSlice(seq)
			}

			// We've exceeded the target sum - there's no reason to keep growing the potential sequence as
			// the sum will only get higher, so let's stop searching at this position (i) and move on to the next one.
			if sum > invalidNumber {
				break
			}
		}
	}

	panic("no encryption weakness found")
}
