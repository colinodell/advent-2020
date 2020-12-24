package main

import (
	"advent-2020/utils"
	"strconv"
)

type CircularList struct {
	values []int
}

func NewCircularList(values ...int) CircularList {
	return CircularList{values: values}
}

func (c *CircularList) GetCurrent() int {
	return c.values[0]
}

func (c *CircularList) Advance() {
	c.values = append(c.values[1:], c.values[0])
}


func (c *CircularList) SetCurrent(value int) {
	for i, v := range c.values {
		if v == value {
			c.SetCurrentPosition(i)
			return
		}
	}

	panic("value does not exist")
}

func (c *CircularList) SetCurrentPosition(pos int) {
	if pos == 0 {
		return
	}

	var ret []int
	for len(ret) < len(c.values) {
		ret = append(ret, c.values[pos])
		pos = (pos + 1) % len(c.values)
	}

	c.values = ret
}

func (c *CircularList) Take(count int, offset int) []int {
	ret := make([]int, count)
	copy(ret, c.values[offset:offset+count])

	c.values = append(c.values[0:offset], c.values[count+offset:]...)

	return ret
}

func (c *CircularList) InsertAfter(target int, values []int) {
	for i, v := range c.values {
		if v == target {
			c.values = append(c.values[:i+1], append(values, c.values[i+1:]...)...)
			return
		}
	}
}

func Play(input string, rounds int) string {
	list := NewCircularList(utils.DigitsFromString(input)...)

	round := 1

	for {
		round++


		currentCup := list.GetCurrent()
		pickedUp := list.Take(3, 1)

		destinationCup := currentCup - 1
		if destinationCup == 0 {
			destinationCup = 9
		}
		for utils.SliceContains(pickedUp, destinationCup) {
			destinationCup--
			if destinationCup == 0 {
				destinationCup = 9
			}
		}

		list.InsertAfter(destinationCup, pickedUp)
		list.Advance()

		rounds--
		if rounds == 0 {
			break
		}
	}

	list.SetCurrent(1)

	ret := ""
	oneFound := false
	for _, v := range list.values {
		if v == 1 {
			oneFound = true
			continue
		} else if !oneFound {
			continue
		}

		ret += strconv.Itoa(v)
	}
	return ret
}
