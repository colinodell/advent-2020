package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
)

func main() {
	lines := utils.ReadLines("./day02/input.txt")
	entries := parseEntries(lines)

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of valid passwords (old policy): %d\n", countCorrect(entries, Entry.IsValidForOldPolicy))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Number of valid passwords (new policy): %d\n", countCorrect(entries, Entry.IsValidForNewPolicy))
}

func countCorrect(entries []Entry, policy func(Entry) bool) int {
	i := 0
	for _, entry := range entries {
		if policy(entry) {
			i++
		}
	}

	return i
}

func parseEntries(lines []string) []Entry {
	entries := make([]Entry, 0)

	re := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)")

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		entries = append(entries, Entry{
			low:      utils.MustParseInt(matches[1]),
			high:     utils.MustParseInt(matches[2]),
			letter:   rune(matches[3][0]),
			password: matches[4],
		})
	}

	return entries
}

type Entry struct {
	low      int
	high     int
	letter   rune
	password string
}

func (e Entry) IsValidForOldPolicy() bool {
	matches := 0
	for _, char := range e.password {
		if char == e.letter {
			matches++
		}
	}

	return matches >= e.low && matches <= e.high
}

func (e Entry) CharMatches(pos int, char rune) bool {
	return rune(e.password[pos - 1]) == char
}

func (e Entry) IsValidForNewPolicy() bool {
	return e.CharMatches(e.low, e.letter) != e.CharMatches(e.high, e.letter)
}
