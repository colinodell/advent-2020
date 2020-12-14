package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main()  {
	input := utils.ReadLines("./day14/input.txt")

	fmt.Println("----- Part 1 -----")
	part1 := NewComputer(&Part1{})
	fmt.Printf("Sum of memory: %d\n", part1.Run(input))

	fmt.Println("----- Part 2 -----")
	part2 := NewComputer(&Part2{})
	fmt.Printf("Sum of memory: %d\n", part2.Run(input))
}

type Computer struct {
	memory map[uint64]uint64
	bitmask string
	setter memorySetter
}

type memorySetter interface {
	Set(c *Computer, address string, value uint64)
}

type Part1 struct {}

func (p *Part1) Set(c *Computer, address string, value uint64) {
	set0s := MustParseUint64(strings.Replace(c.bitmask, "X", "1", -1), 2)
	set1s := MustParseUint64(strings.Replace(c.bitmask, "X", "0", -1), 2)

	c.memory[MustParseUint64(address, 10)] = (value & set0s) | set1s
}

type Part2 struct {}

func (p *Part2) Set(c *Computer, address string, value uint64) {
	if len(address) < 36 {
		addressInBinary := fmt.Sprintf("%036b", MustParseUint64(address, 10))
		var sb strings.Builder
		for i, addressBit := range addressInBinary {
			if c.bitmask[i] == '0' {
				sb.WriteRune(addressBit)
			} else {
				sb.WriteRune(rune(c.bitmask[i]))
			}
		}
		address = sb.String()
	}

	if !strings.Contains(address, "X") {
		c.memory[MustParseUint64(address, 2)] = value
		return
	}

	p.Set(c, strings.Replace(address, "X", "0", 1), value)
	p.Set(c, strings.Replace(address, "X", "1", 1), value)
}

func NewComputer(setter memorySetter) Computer {
	return Computer{
		memory: make(map[uint64]uint64),
		setter: setter,
	}
}

var memRegex = regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")

func (c *Computer) Run(instructions []string) uint64 {
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "mask = ") {
			c.bitmask = strings.Replace(instruction, "mask = ", "", -1)
			continue
		}

		matches := memRegex.FindStringSubmatch(instruction)
		c.setter.Set(c, matches[1], MustParseUint64(matches[2], 10))
	}

	return c.SumMemory()
}

func (c *Computer) SumMemory() (sum uint64) {
	for _, v := range c.memory {
		sum += v
	}

	return sum
}

func MustParseUint64(input string, base int) uint64 {
	val, err := strconv.ParseUint(input, base, 64)
	if err != nil {
		panic("invalid value")
	}

	return val
}
