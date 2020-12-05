package main

import (
	"advent-2020/utils"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	batchFile := utils.ReadFile("./day04/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of valid passports (based on required fields): %d\n", countValid(batchFile, false))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Number of valid passports (validating field values too): %d\n", countValid(batchFile, true))
}

func countValid(input string, checkValues bool) (count int) {
	// Normalize the data so each key:value pair is on a single line, then split the batch up into single passports
	passports := strings.Split(strings.Replace(input, " ", "\n", -1), "\n\n")

	// Validate and count the valid passports
	for _, passport := range passports {
		if validatePassport(passport, checkValues) {
			count++
		}
	}

	return
}

var fields = map[string]*regexp.Regexp {
	"byr": regexp.MustCompile("^19[2-9][0-9]|200[0-2]$"),
	"iyr": regexp.MustCompile("^20(?:1[0-9]|20)$"),
	"eyr": regexp.MustCompile("^20(?:2[0-9]|30)$"),
	"hgt": regexp.MustCompile("^(?:(?:59|6[0-9]|7[0-6])in)|(?:(?:1[5-8][0-9]|19[0-3])cm)$"),
	"hcl": regexp.MustCompile("^#[0-9a-f]{6}$"),
	"ecl": regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$"),
	"pid": regexp.MustCompile("^\\d{9}$"),
}

var lineRegex = regexp.MustCompile("(\\w+):(.+)")

func validatePassport(input string, checkValues bool) bool {
	passportData := lineRegex.FindAllStringSubmatch(input, -1)

	// Verify each required field exists and optionally validate the values
	for fieldName, validator := range fields {
		value, err := getFieldValue(fieldName, passportData)
		if err != nil || (checkValues && !validator.MatchString(value)) {
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
