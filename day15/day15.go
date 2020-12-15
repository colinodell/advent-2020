package main

import (
	"fmt"
)

func main() {
	input := []int{0, 8, 15, 2, 12, 1, 4}

	fmt.Println("----- Part 1 -----")
	fmt.Printf("After 2020 turns: %d\n", Play(input, 2020))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("After 30000000 turns: %d\n", Play(input, 30000000))
}

func Play(starting []int, maxTurns int) int {
	lastSpoken := make(map[int]int)

	var previous int
	speak := starting[0]

	for turn := 0; turn < maxTurns; turn++ {
		previous = speak

		if turn < len(starting) {
			speak = starting[turn]
		} else if _, ok := lastSpoken[speak]; !ok {
			speak = 0
		} else {
			speak = turn - lastSpoken[speak] - 1
		}

		lastSpoken[previous] = turn - 1
	}

	return speak
}
