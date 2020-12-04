package main

import (
	"advent-2020/utils"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	batchFile := utils.ReadFile("./day04/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of valid passports (based on required fields): %d\n", countValid(batchFile, false))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Number of valid passports (validating field values too): %d\n", countValid(batchFile, true))
}

func countValid(input string, checkValues bool) int {
	count := 0

	// Normalize the data so each key:value pair is on a single line, then split the batch up into single passports
	passports := strings.Split(strings.Replace(input, " ", "\n", -1), "\n\n")

	// Validate and count the valid passports
	for _, passport := range passports {
		if validatePassport(passport, checkValues) {
			count++
		}
	}

	return count
}

var fields = map[string]func(string) bool {
	"byr": func(input string) bool {
		year, err := strconv.Atoi(input)
		return err == nil && year >= 1920 && year <= 2002
	},
	"iyr": func(input string) bool {
		year, err := strconv.Atoi(input)
		return err == nil && year >= 2010 && year <= 2020
	},
	"eyr": func(input string) bool {
		year, err := strconv.Atoi(input)
		return err == nil && year >= 2020 && year <= 2030
	},
	"hgt": func(input string) bool {
		value, units := input[:len(input)-2], input[len(input)-2:]
		height, err := strconv.Atoi(value)
		if err != nil {
			return false
		} else if units == "cm" {
			return height >= 150 && height <= 193
		} else if units == "in" {
			return height >= 59 && height <= 76
		} else {
			return false
		}
	},
	"hcl": func(input string) bool {
		re := regexp.MustCompile("^#[0-9a-f]{6}$")
		return re.MatchString(input)
	},
	"ecl": func(input string) bool {
		return input == "amb" || input == "blu" || input == "brn" || input == "gry" || input == "grn" || input == "hzl"|| input == "oth"
	},
	"pid": func(input string) bool {
		number, err := strconv.Atoi(input)
		return err == nil && len(input) == 9 && number > 0 && number <= 999999999
	},
}

var lineRegex = regexp.MustCompile("(\\w+):(.+)")

func validatePassport(input string, checkValues bool) bool {
	matches := lineRegex.FindAllStringSubmatch(input, -1)

	for fieldName, validator := range fields {
		value, err := getFieldValue(fieldName, matches)
		if err != nil || (checkValues && !validator(value)) {
			return false
		}
	}

	return true
}

func getFieldValue(needle string, haystack [][]string) (string, error) {
	for _, row := range haystack {
		if row[1] == needle {
			return row[2], nil
		}
	}

	return "", errors.New("field not found")
}
