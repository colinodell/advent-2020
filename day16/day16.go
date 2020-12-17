package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	puzzle := NewPuzzle(utils.ReadFile("./day16/input.txt"))

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Calculated error rate: %d\n", puzzle.CalculateErrorRate())

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Product of 'departure' tickets: %d\n", puzzle.MultiplyMyDepartureFields())
}

// Data structures

type Puzzle struct {
	fields []*Field
	myTicket Ticket
	otherTickets []Ticket
}

type Ticket []int

type Field struct {
	name string
	pos int
	validRanges [2][2]int
}

// Parsing functions

func NewPuzzle(input string) Puzzle {
	p := Puzzle{}

	parts := strings.Split(input, "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		p.fields = append(p.fields, NewField(line))
	}

	for i, line := range strings.Split(parts[1], "\n") {
		if i != 0 {
			p.myTicket = NewTicket(line)
		}
	}

	for i, line := range strings.Split(parts[2], "\n") {
		if i != 0 {
			p.otherTickets = append(p.otherTickets, NewTicket(line))
		}
	}

	return p
}

func NewTicket(input string) Ticket {
	t := make(Ticket, 0)

	for _, v := range strings.Split(input, ",") {
		t = append(t, utils.MustParseInt(v))
	}

	return t
}

var fieldRegex = regexp.MustCompile("(.+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")

func NewField(line string) *Field {
	matches := fieldRegex.FindStringSubmatch(line)

	return &Field{
		name: matches[1],
		pos: -1,
		validRanges: [2][2]int{
			{
				utils.MustParseInt(matches[2]),
				utils.MustParseInt(matches[3]),
			},
			{
				utils.MustParseInt(matches[4]),
				utils.MustParseInt(matches[5]),
			},
		},
	}
}

// Part 1

func (p *Puzzle) CalculateErrorRate() int {
	errorRate := 0
	for _, ticket := range p.otherTickets {
		if seemsValid, invalidValue := ticket.naiveValidation(p.fields); !seemsValid {
			errorRate += invalidValue
		}
	}

	return errorRate
}

func (f Field) Validate(val int) bool {
	return (val >= f.validRanges[0][0] && val <= f.validRanges[0][1]) || (val >= f.validRanges[1][0] && val <= f.validRanges[1][1])
}

func (t Ticket) naiveValidation(fields []*Field) (bool, int) {
	for _, val := range t {
		valueIsValid := false
		for _, f := range fields {
			if f.Validate(val) {
				valueIsValid = true
			}
		}

		if !valueIsValid {
			return false, val
		}
	}

	return true, 0
}

// Part 2
func (p *Puzzle) getUnidentifiedFields() map[*Field]bool {
	ret := make(map[*Field]bool)

	for _, f := range p.fields {
		if f.pos == -1 {
			ret[f] = true
		}
	}

	return ret
}

func (p *Puzzle) identifyFields() {
	// Only consider the valid tickets
	validTickets := make([]Ticket, 0)
	for _, t := range p.otherTickets {
		if seemsValid, _ := t.naiveValidation(p.fields); seemsValid {
			validTickets = append(validTickets, t)
		}
	}

	// Build a list of fields we don't know the field index position for
	fieldsToSolve := p.getUnidentifiedFields()
	for len(fieldsToSolve) > 0 {
		for pos := 0; pos < len(p.myTicket); pos++ {
			for _, f := range p.fields {
				for _, t := range validTickets {
					if !f.Validate(t[pos]) {
						delete(fieldsToSolve, f)
						break
					}
				}
			}

			if len(fieldsToSolve) == 1 {
				// We solved a field!
				for f, _ := range fieldsToSolve {
					f.pos = pos
				}
			}

			fieldsToSolve = p.getUnidentifiedFields()
		}
	}
}

func (p *Puzzle) MultiplyMyDepartureFields() int {
	p.identifyFields()

	result := 1

	for _, field := range p.fields {
		if strings.Contains(field.name, "departure") {
			result *= p.myTicket[field.pos]
		}
	}

	return result
}
