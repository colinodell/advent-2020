package main

import (
	"advent-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	boardingPasses := utils.ReadLines("./day05/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Highest sead ID on a boarding pass: %d\n", FindHighestSeatId(boardingPasses))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("The ID of your seat: %d\n", FindEmptySeat(boardingPasses))
}

func FindHighestSeatId(passes []string) int {
	max := -1
	for _, pass := range passes {
		id := BoardingPassToId(pass)
		max = utils.Max(max, id)
	}

	return max
}

func FindEmptySeat(passes []string) int {
	seats := make([]bool, 1 << 10)

	// Mark the already-claimed seats
	for _, pass := range passes {
		seats[BoardingPassToId(pass)] = true
	}

	// Find the first empty one with a seat on either side (ID-wise)
	for i, taken := range seats {
		if i > 0 && !taken && seats[i-1] && seats[i+1] {
			return i
		}
	}

	panic("no seats left")
}

var charsToBits = map[string]string {
	"F": "0",
	"B": "1",
	"R": "1",
	"L": "0",
}

func BoardingPassToId(pass string) int {
	for char, bit := range charsToBits {
		pass = strings.ReplaceAll(pass, char, bit)
	}

	if i, err := strconv.ParseInt(pass, 2, 16); err == nil {
		return int(i)
	}

	panic("invalid boarding pass")
}
