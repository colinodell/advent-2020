package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	rules := NewLuggageRules(utils.ReadFile("./day07/input.txt"))

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of bag colors that can eventually contain 'shiny gold': %d\n", len(rules.GetBagColorsThatContain("shiny gold")))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Individual bags needed inside the single 'shiny gold' bag:: %d\n", rules.CountTotalInnerBagsRequired("shiny gold"))
}

type BagColor string
type Bag struct {
	mustContain map[BagColor]int
}
type LuggageRules map[BagColor]*Bag

func NewLuggageRules(input string) LuggageRules {
	rules := make(LuggageRules)

	ruleRegex := regexp.MustCompile("([\\w ]+) bags contain (.+)\\.")
	bagColorQtyRegex := regexp.MustCompile("(\\d+) ([\\w ]+) bags?")

	for _, line := range strings.Split(input, "\n") {
		parts := ruleRegex.FindStringSubmatch(line)
		color := BagColor(parts[1])

		for _, c := range strings.Split(parts[2], ", ") {
			if c == "no other bags" {
				break
			}

			match := bagColorQtyRegex.FindStringSubmatch(c)
			rules.Get(color).mustContain[BagColor(match[2])] = utils.MustParseInt(match[1])
		}
	}

	return rules
}

func (rules LuggageRules) Get (color BagColor) *Bag {
	if _, ok := rules[color]; !ok {
		rules[color] = &Bag{
			mustContain: make(map[BagColor]int),
		}
	}

	return rules[color]
}

func (rules LuggageRules) GetBagColorsThatContain(target BagColor) map[BagColor]bool {
	result := make(map[BagColor]bool)

	for color, bagDefinition := range rules {
		if _, ok := bagDefinition.mustContain[target]; ok {
			result[color] = true
			for foo, _ := range rules.GetBagColorsThatContain(color) {
				result[foo] = true
			}
		}
	}

	return result
}

func (rules LuggageRules) CountTotalInnerBagsRequired(color BagColor) (count int) {
	for color, qty := range rules.Get(color).mustContain {
		count += qty + (qty * rules.CountTotalInnerBagsRequired(color))
	}

	return
}
