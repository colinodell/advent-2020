package main

import (
	"advent-2020/utils"
	"errors"
	"strconv"
	"strings"
)

type Game struct {
	players []*Player
}

type Player struct {
	topCard    *Card
	bottomCard *Card
}

// Card deck is implemented as a doubly-linked list
type Card struct {
	value int
	prev *Card
	next *Card
}

// Adds one or more cards to the bottom of the Player's deck
func (p *Player) AddCards(values ...int) {
	for _, value := range values {
		card := &Card{
			prev:  p.bottomCard,
			next:  nil,
			value: value,
		}

		if p.bottomCard != nil {
			p.bottomCard.next = card
		}

		p.bottomCard = card

		if p.topCard == nil {
			p.topCard = card
		}
	}
}

func (p *Player) HasCards() bool {
	return p.topCard != nil
}

func (p *Player) PlayTopCard() int {
	top := p.topCard

	p.topCard = top.next
	if p.topCard == nil {
		// No cards left
		p.bottomCard = nil
	} else {
		p.topCard.prev = nil
	}

	return top.value
}

func (p *Player) CalculateScore() int {
	score := 0
	i := 1
	card := p.bottomCard
	for card != nil {
		score += i * card.value
		i++
		card = card.prev
	}

	return score
}

func (g *Game) getWinner() (*Player, error) {
	var winner *Player
	for _, p := range g.players {
		if p.HasCards() {
			if winner != nil {
				return nil, errors.New("game has not ended yet")
			}

			winner = p
		}
	}

	return winner, nil
}

// Plays a full game and returns the winner
func (g *Game) Play() *Player {
	for {
		g.PlayRound()
		if winner, _ := g.getWinner(); winner != nil {
			return winner
		}
	}
}

// Plays a single round
func (g *Game) PlayRound() {
	p1Card, p2Card := g.players[0].PlayTopCard(), g.players[1].PlayTopCard()

	if p1Card > p2Card {
		g.players[0].AddCards(p1Card, p2Card)
	} else {
		g.players[1].AddCards(p2Card, p1Card)
	}
}

func NewGame(input string) Game {
	s := strings.Split(input, "\n\n")
	players := make([]*Player, len(s))
	for i, data := range s {
		p := Player{}

		c := strings.Split(data, "\n")
		for cardIndex := 1; cardIndex < len(c); cardIndex++ {
			p.AddCards(utils.MustParseInt(c[cardIndex]))
		}

		players[i] = &p
	}

	return Game{players: players}
}
