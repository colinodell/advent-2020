package main

import (
	"advent-2020/utils"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := utils.ReadFile("./day19/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Matching rules: %d\n", Part1(input))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Matching rules: %d\n", Part2(input))
}

type Rule struct {
	contents string
	literal bool
	needs []int
}

func Part1 (input string) int {
	s := strings.Split(input, "\n\n")
	rules, inputs := s[0], s[1]
	regex := regexp.MustCompile(ParseRegex(strings.Split(rules, "\n")))

	cnt := 0
	for _, line := range strings.Split(inputs, "\n") {
		if regex.MatchString(line) {
			cnt++
		}
	}

	return cnt
}

func Part2 (input string) int {
	// TL;DR: Dirty hacks
	//
	// Longer explanation:
	//
	// My part 1 approach wasn't going to work for part 2 as-is; it just kept looping forever waiting for
	// its dependencies (itself) to be resolved which would never happen.
	// I could have modified it to detect those loops and then program it try "recursing" itself but that would
	// lead to infinite loops.  So my next thought was that I could limit the depth so we know the input strings
	// will never exceed a certain length, meaning we don't need to "recurse" infinitely, only to some low "maximum"
	// depth that would work for our input set.
	// But then I had another idea - we already know exactly which rules are going to have this problem, so instead
	// of writing code to detect and implement that logic, why not just hard-code the rules it would have dynamically
	// generated as hard-coded inputs?
	// And thus the approach below was born:
	input = regexp.MustCompile("(?m)^8: .+$").ReplaceAllString(input, strings.Join([]string{
		"8: 42 | 42 99081",
		"99081: 42 | 42 99082",
		"99082: 42 | 42 99083",
		"99083: 42 | 42 99084",
		"99084: 42 | 42 99085",
		"99085: 42 | 42 99086",
		"99086: 42 | 42 99087",
		"99087: 42 | 42 99088",
		"99088: 42 | 42 99089",
		"99089: 42",
	}, "\n"))

	input = regexp.MustCompile("(?m)^11: .+$").ReplaceAllString(input, strings.Join([]string{
		"11: 42 31 | 42 99111 31",
		"99111: 42 31 | 42 99112 31",
		"99112: 42 31 | 42 99113 31",
		"99113: 42 31 | 42 99114 31",
		"99114: 42 31 | 42 99115 31",
		"99115: 42 31 | 42 99116 31",
		"99116: 42 31 | 42 99117 31",
		"99117: 42 31 | 42 99118 31",
		"99118: 42 31 | 42 99119 31",
		"99119: 42 31",
	}, "\n"))

	return Part1(input)
}

func ParseRegex(lines []string) string {
	resolved, unresolved := make(map[int]Rule), make(map[int]Rule)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		index, rule := utils.MustParseInt(split[0]), Rule{
			contents: split[1],
		}

		if rule.contents[0] == '"' {
			rule.contents = strings.Replace(rule.contents, "\"", "", 2)
			rule.literal = true
			resolved[index] = rule
		} else {
			rule.needs = parseUniqueNumbersFromString(strings.Replace(rule.contents, " | ", " ", -1))
			unresolved[index] = rule
		}
	}

	for len(unresolved) > 0 {
		for i, rule := range unresolved {
			if !containsAllKeys(resolved, rule.needs) {
				continue
			}

			// We can resolve this one!
			scanner := bufio.NewScanner(strings.NewReader(rule.contents))
			scanner.Split(bufio.ScanWords)
			var sb strings.Builder

			for scanner.Scan() {
				token := scanner.Text()
				if token == "|" {
					sb.WriteString("|")
					continue
				}

				otherRule := resolved[utils.MustParseInt(token)]
				sb.WriteString(otherRule.contents)
			}

			rule.contents = "(" + sb.String() + ")"

			resolved[i] = rule
			delete(unresolved, i)
		}
	}

	// All rules have been parsed
	return "^" + resolved[0].contents + "$"
}

func parseUniqueNumbersFromString(s string) []int {
	keys := make(map[int]bool)
	var list []int
	for _, n := range strings.Split(s, " ") {
		num := utils.MustParseInt(n)
		if _, value := keys[num]; !value {
			keys[num] = true
			list = append(list, num)
		}
	}
	return list
}

func containsAllKeys(haystack map[int]Rule, needle []int) bool {
	for _, search := range needle {
		if _, ok := haystack[search]; !ok {
			return false
		}
	}

	return true
}
