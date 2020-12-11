package main

import (
	"advent-2020/utils"
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("----- Part 1 -----")
	seatingPart1 := NewSeating(utils.ReadFile("./day11/input.txt"))
	seatingPart1.RunUntilStable(Part1)
	fmt.Printf("Number of occupied seats: %d\n", seatingPart1.CountAllOccupiedSeats())

	fmt.Println("----- Part 2 -----")
	seatingPart2 := NewSeating(utils.ReadFile("./day11/input.txt"))
	seatingPart2.RunUntilStable(Part2)
	fmt.Printf("Number of occupied seats: %d\n", seatingPart2.CountAllOccupiedSeats())
}

type Rules struct {
	adjacentOnly bool
	tolerance int
}

var Part1 = Rules{adjacentOnly: true, tolerance: 4}
var Part2 = Rules{adjacentOnly: false, tolerance: 5}

var directions = []utils.Vector2{
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
}

type Seating struct {
	grid map[utils.Vector2]rune
	width, height int
}

func NewSeating(str string) Seating {
	input := strings.Split(str, "\n")

	s := Seating{
		grid: make(map[utils.Vector2]rune),
		width: len(input[0]),
		height: len(input),
	}

	for y, row := range input {
		for x, char := range row {
			s.grid[utils.Vector2{X: x, Y: y}] = char
		}
	}

	return s
}

func (s *Seating) get(pos utils.Vector2) rune {
	if val, ok := s.grid[pos]; ok {
		return val
	}

	return ' '
}

func (s *Seating) countNearbyOccupiedSeats(pos utils.Vector2, rules Rules) int {
	if rules.adjacentOnly {
		return s.countAdjacentOccupiedSeats(pos)
	} else {
		return s.countVisibleOccupiedSeats(pos)
	}
}

func (s *Seating) countAdjacentOccupiedSeats(pos utils.Vector2) int {
	cnt := 0

	for _, dir := range directions {
		if s.get(pos.Add(dir)) == '#' {
			cnt++
		}
	}

	return cnt
}

func (s *Seating) countVisibleOccupiedSeats(pos utils.Vector2) int {
	cnt := 0

	for _, dir := range directions {
		distance := 1
		for {
			v := s.get(pos.Add(dir.Multiply(distance)))
			if v == ' ' || v == 'L' {
				break
			} else if v == '.' {
				distance++
				continue
			}

			cnt++
			break
		}
	}

	return cnt
}

func (s *Seating) CountAllOccupiedSeats() int {
	cnt := 0
	for _, char := range s.grid {
		if char == '#' {
			cnt++
		}
	}

	return cnt
}

func (s *Seating) Step(rules Rules) {
	// Create a new grid by applying the rules against the old grid
	newGrid := make(map[utils.Vector2]rune)
	for pos, v := range s.grid {
		if v == 'L' && s.countNearbyOccupiedSeats(pos, rules) == 0 {
			newGrid[pos] = '#'
		} else if v == '#' && s.countNearbyOccupiedSeats(pos, rules) >= rules.tolerance {
			newGrid[pos] = 'L'
		} else {
			newGrid[pos] = v
		}
	}

	s.grid = newGrid
}

func (s *Seating) RunUntilStable(rules Rules) {
	occupied := s.CountAllOccupiedSeats()
	for {
		s.Step(rules)

		newOccupied := s.CountAllOccupiedSeats()
		if newOccupied == occupied {
			return
		}

		occupied = newOccupied
	}
}

