package utils

import (
	"bufio"
	"encoding/csv"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	check(err)

	return strings.TrimSpace(string(buf))
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadLinesOfNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, ToInt(scanner.Text()))
	}
	return numbers
}

func ReadCSVNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	r := csv.NewReader(file)
	row, err := r.Read()

	var numbers []int

	for _, number := range row {
		numbers = append(numbers, ToInt(number))
	}

	return numbers
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
