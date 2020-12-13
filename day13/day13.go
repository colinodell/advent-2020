package main

import (
	"advent-2020/utils"
	"fmt"
	"strings"
)

func main()  {
	input := utils.ReadLines("./day13/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("ID of the next bus * wait time: %d\n", SolvePart1(input))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Earliest timestamp with sequential departures: %d\n", SolvePart2(input[1]))
}

func SolvePart1(input []string) int {
	// Parse the timestamp and bus schedule
	timestamp := utils.MustParseInt(input[0])
	var buses []int
	for _, id := range strings.Split(input[1], ",") {
		if id == "x" {
			continue
		}
		buses = append(buses, utils.MustParseInt(id))
	}

	// Find the next departure
	for i := 0; ; i++ {
		for _, bus := range buses {
			if (timestamp + i) % bus == 0 {
				return i * bus
			}
		}
	}
}

func SolvePart2(schedule string) int {
	// Parse the bus schedule.
	// A given bus will depart when the timestamp is evenly divisible by the bus ID, so we'll put those schedules
	// into a slice where the key is the timestamp offset and the value is the bus ID (factor to divide by).
	var buses []int
	for _, id := range strings.Split(schedule, ",") {
		if id == "x" {
			// Bus departure time doesn't matter, so we'll use a factor of 1 which will always match
			id = "1"
		}
		buses = append(buses, utils.MustParseInt(id))
	}

	timestamp := 1

	for {
		timeToSkipIfNoMatch := 1
		valid := true

		for offset := 0; offset < len(buses); offset++ {
			// A given bus will depart when the timestamp is evenly divisible by the bus ID
			if (timestamp + offset) % buses[offset] != 0 {
				// No match here; abort and we'll try the next potential timestamp
				valid = false
				break
			}

			// This particular bus schedule matches, but we don't know if subsequent ones will. However, we
			// do know when this particular schedule will match again, so let's keep track of that.
			// For example, if the first bus is Bus 7, we know it won't depart again for another 7 minutes,
			// so we'll skip ahead by 7 minutes and ignore the timestamps we know won't match.
			//
			// If we find a partial match where some schedules match but not the whole thing, we can calculate
			// the next time those schedules align by multiplying their values together; worst case, we
			// still only have 2 aligned schedules, or best case we find yet another matching departure!
			//
			// For example, let's say we find a timestamp where the first two buses (7 and 11) align but none
			// of the others do; in that case, we know that buses 7 and 11 won't align again for another
			// 77 (7*11) minutes, so we'll skip ahead 77 minutes. Eventually we might find that now buses
			// 7, 11, and 13 align, but none others do. Well, that means the next time that these three buses
			// align will be in 7*11*13 minutes, so skip ahead by that much and try again there.
			//
			// This approach significantly speeds up the search and the speed improves the bigger your
			// partial match is!
			//
			// (Note that technically we'd need to find the LCM of those bus IDs, but luckily our inputs are
			// always prime numbers so the LCM will always equal the product of those IDs.)
			timeToSkipIfNoMatch *= buses[offset]
		}

		// Did we find a full match?
		if valid {
			return timestamp
		}

		timestamp += timeToSkipIfNoMatch
	}
}
