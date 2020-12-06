package main

import (
	"advent-2020/utils"
	"fmt"
	"strings"
)

func main()  {
	input := utils.ReadFile("./day06/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Sum of the counts (incorrect): %d\n", SumCountsIncorrectly(input))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Sum of the counts (correct): %d\n", SumCountsCorrectly(input))
}

func SumCountsIncorrectly(input string) int {
	sum := 0

	for _, group := range strings.Split(input, "\n\n") {
		yesQuestions := make(map[rune]bool)
		for _, char := range group {
			if char == '\n' {
				continue
			}
			yesQuestions[char] = true
		}

		for range yesQuestions {
			sum++
		}
	}

	return sum
}

func SumCountsCorrectly(input string) int {
	sum := 0

	for _, group := range strings.Split(input, "\n\n") {
		people := 1
		yesQuestions := make(map[rune]int)
		for _, char := range group {
			if char == '\n' {
				people++
			}
			if _, ok := yesQuestions[char]; !ok {
				yesQuestions[char] = 1
			} else {
				yesQuestions[char]++
			}
		}

		for _, count := range yesQuestions {
			if count == people {
				sum++
			}
		}
	}

	return sum
}
