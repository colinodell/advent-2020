package main

import (
	"advent-2020/utils"
	"fmt"
	"strconv"
)

func main() {
	input := "739862541"

	{
		fmt.Println("----- Part 1 -----")
		game := NewCupGame(input, 9)
		game.Play(10)
		fmt.Printf("Labels on the cups after cup 1: %s\n", game.LabelsAfterCup1())
	}

	{
		fmt.Println("----- Part 2 -----")
		game := NewCupGame(input, 1000000)
		game.Play(10000000)
		fmt.Printf("Product of the labels containing stars: %d\n", game.GetProductOfStarLabels())
	}
}

type CupGame struct {
	cups []int
	current int
	max int
}

func NewCupGame(input string, max int) CupGame {
	inputDigits := utils.DigitsFromString(input)

	cg := CupGame{cups: make([]int, max+1), current: 0, max: max}
	start := 0
	last := 0

	for i := 0; i < len(inputDigits); i++ {
		if start == 0 {
			start = inputDigits[0]
		}

		if i < (len(inputDigits) - 1) {
			// Point to the next cup
			cg.cups[inputDigits[i]] = inputDigits[i+1]
		} else {
			// We're at the end
			last = inputDigits[i]
			cg.cups[last] = start
		}
	}

	// Do we need extra digits?
	for i := len(inputDigits) + 1; i <= max; i++ {
		cg.cups[last] = i
		last = i
	}

	cg.cups[last] = start

	cg.current = start

	return cg
}

func (cg *CupGame) Play(rounds int) {
	for {
		// Take three cups
		cup1 := cg.cups[cg.current]
		cup2 := cg.cups[cup1]
		cup3 := cg.cups[cup2]
		after := cg.cups[cup3]

		destination := cg.current - 1
		if destination == 0 {
			destination = cg.max
		}
		for cup1 == destination || cup2 == destination || cup3 == destination {
			destination--
			if destination == 0 {
				destination = cg.max
			}
		}

		// Remove the three cups
		cg.cups[cg.current] = after

		// Insert them after the destination
		oldDestValue := cg.cups[destination]
		cg.cups[destination] = cup1
		cg.cups[cup3] = oldDestValue

		cg.current = after

		rounds--
		if rounds == 0 {
			break
		}
	}
}

func (cg *CupGame) LabelsAfterCup1() string {
	cg.current = cg.cups[1]

	ret := ""
	for cg.current = cg.cups[1]; cg.current != 1; cg.current = cg.cups[cg.current] {
		ret += strconv.Itoa(cg.current)
	}
	return ret
}

func (cg *CupGame) GetProductOfStarLabels() int {
	c1 := cg.cups[1]
	c2 := cg.cups[c1]

	return c1 * c2
}
