package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
)

func main()  {
	input := utils.ReadLines("./day12/input.txt")

	fmt.Println("----- Part 1 -----")
	shipWithBadNavigator := NewShipWithBadNavigator()
	fmt.Printf("Manhattan distance based on faulty navigation: %d\n", shipWithBadNavigator.Follow(input))

	fmt.Println("----- Part 2 -----")
	shipWithGoodNavigator := NewShipWithGoodNavigator()
	fmt.Printf("Manhattan distance based on correct navigation: %d\n", shipWithGoodNavigator.Follow(input))
}

var (
	Up = utils.Vector2{X:0, Y:1}
	Down = utils.Vector2{X:0, Y:-1}
	Left = utils.Vector2{X:-1, Y:0}
	Right = utils.Vector2{X:1, Y:0}

	directions = map[int]*utils.Vector2{0: &Up, 90: &Right, 180: &Down, 270: &Left}
)

type Navigator interface {
	Navigate(ship *Ship, direction string, amount int)
}

type Ship struct {
	pos utils.Vector2
	nav Navigator
}

func NewShipWithBadNavigator() Ship {
	ship := Ship{}
	ship.nav = &BadNavigator{heading: 90}

	return ship
}

func NewShipWithGoodNavigator() Ship {
	ship := Ship{}
	ship.nav = &GoodNavigator{waypoint: utils.Vector2{X: 10, Y: 1}}

	return ship
}

var actionRegex = regexp.MustCompile("(\\w)(\\d+)")

func (s *Ship) Follow(instructions []string) int {
	for _, instruction := range instructions {
		match := actionRegex.FindStringSubmatch(instruction)
		direction, amount := match[1], utils.MustParseInt(match[2])
		s.nav.Navigate(s, direction, amount)
	}

	return s.pos.ManhattanDistance()
}

type BadNavigator struct {
	heading int
}

func (nav *BadNavigator) Navigate(ship *Ship, direction string, amount int) {
	switch direction {
	case "N":
		ship.pos = ship.pos.Add(Up.Multiply(amount))
		break
	case "S":
		ship.pos = ship.pos.Add(Down.Multiply(amount))
		break
	case "E":
		ship.pos = ship.pos.Add(Right.Multiply(amount))
		break
	case "W":
		ship.pos = ship.pos.Add(Left.Multiply(amount))
		break
	case "F":
		ship.pos = ship.pos.Add(directions[nav.heading].Multiply(amount))
		break
	case "L":
		nav.heading = (360 + nav.heading - amount) % 360
		break
	case "R":
		nav.heading = (360 + nav.heading + amount) % 360
		break
	}
}

type GoodNavigator struct {
	waypoint utils.Vector2
}

func (nav *GoodNavigator) Navigate(ship *Ship, direction string, amount int) {
	switch direction {
	case "N":
		nav.waypoint = nav.waypoint.Add(Up.Multiply(amount))
		break
	case "S":
		nav.waypoint = nav.waypoint.Add(Down.Multiply(amount))
		break
	case "E":
		nav.waypoint = nav.waypoint.Add(Right.Multiply(amount))
		break
	case "W":
		nav.waypoint = nav.waypoint.Add(Left.Multiply(amount))
		break
	case "F":
		ship.pos = ship.pos.Add(nav.waypoint.Multiply(amount))
		break
	case "L":
		nav.waypoint = nav.waypoint.RotateCounterClockwise(amount / 90)
		break
	case "R":
		nav.waypoint = nav.waypoint.RotateClockwise(amount / 90)
		break
	}
}
