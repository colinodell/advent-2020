package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
)

func main()  {
	console := NewGameConsole(utils.ReadLines("./day08/input.txt"))

	fmt.Println("----- Part 1 -----")
	console.Run()
	fmt.Printf("Value of the accumulator immediately before the infinite loop: %d\n", console.Accumulator)

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Value of the accumulator after fixing the program: %d\n", console.FixProgram())
}

type Instruction struct {
	operation string
	argument int
}

type GameConsole struct {
	program     []Instruction
	pos         int
	Accumulator int
}

func NewGameConsole(program []string) GameConsole {
	gc := GameConsole{
		program: make([]Instruction, len(program)),
	}

	for i, line := range program {
		gc.program[i] = parseInstruction(line)
	}

	return gc
}

type Result int
const (
	Terminated Result = iota
	InfiniteLoop
)

func (gc *GameConsole) Run() Result {
	instructionsExecuted := make(map[int]bool, len(gc.program))
	gc.Accumulator = 0
	gc.pos = 0

	for {
		if gc.pos >= len(gc.program) {
			return Terminated
		}

		if instructionsExecuted[gc.pos] {
			return InfiniteLoop
		}

		instructionsExecuted[gc.pos] = true

		nextInstruction := gc.program[gc.pos]

		switch nextInstruction.operation {
		case "nop":
			gc.pos++
			break
		case "acc":
			gc.Accumulator += nextInstruction.argument
			gc.pos++
			break
		case "jmp":
			gc.pos += nextInstruction.argument
			break
		}
	}
}

func (gc *GameConsole) FixProgram() int {
	// Make a copy of the original program for our reference
	originalProgram := make([]Instruction, len(gc.program))
	copy(originalProgram, gc.program)

	for i := 0; i < len(originalProgram); i++ {
		// Swap the operation if needed
		if originalProgram[i].operation == "nop" {
			gc.program[i].operation = "jmp"
		} else if originalProgram[i].operation == "jmp" {
			gc.program[i].operation = "nop"
		} else {
			continue
		}

		// Run the modified program and see if it terminated successfully
		if gc.Run() == Terminated {
			return gc.Accumulator
		}

		// Reset the program so we can try changing the next instruction
		gc.program[i].operation = originalProgram[i].operation
	}

	panic("could not fix program")
}

var instructionRegex = regexp.MustCompile("(\\w+) \\+?(-?\\d+)")
func parseInstruction(line string) Instruction {
	match := instructionRegex.FindStringSubmatch(line)

	return Instruction{
		operation: match[1],
		argument: utils.MustParseInt(match[2]),
	}
}


