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
	count int
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

		p.count++
	}
}

func (p *Player) HasCards() bool {
	return p.count > 0
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

	p.count--

	return top.value
}

// Returns a representation of the deck as a string
func (p *Player) String() string {
	var sb strings.Builder
	node := p.topCard
	for node != nil {
		sb.WriteString(strconv.Itoa(node.value))
		if node.next != nil {
			sb.WriteString(",")
		}
		node = node.next
	}

	return sb.String()
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

func (p *Player) Copy(count int) *Player {
	newPlayer := &Player{}

	nextCard := p.topCard

	for newPlayer.count < count {
		newPlayer.AddCards(nextCard.value)
		nextCard = nextCard.next
	}

	return newPlayer
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

// Plays a full game of recursive combat and returns the winner
func (g *Game) PlayRecursive() *Player {
	previousRounds := make(map[string]struct{}, len(g.players))

	for {
		// Check if we've played an identical round before
		currentRoundString := g.players[0].String() + "|" + g.players[1].String()
		if _, exists := previousRounds[currentRoundString]; exists {
			return g.players[0]
		}
		// We have not; record the current round to check against in the future
		previousRounds[currentRoundString] = struct{}{}

		// Have each player draw a card
		p1Card, p2Card := g.players[0].PlayTopCard(), g.players[1].PlayTopCard()
		if p1Card <= g.players[0].count && p2Card <= g.players[1].count {
			// The winner of the round is determined by recursing into a sub-game of Recursive Combat.
			newP1, newP2 := g.players[0].Copy(p1Card), g.players[1].Copy(p2Card)
			subGame := Game{players: []*Player{newP1, newP2}}
			
			winner := subGame.PlayRecursive()
			if winner == newP1 {
				g.players[0].AddCards(p1Card, p2Card)
			} else {
				g.players[1].AddCards(p2Card, p1Card)
			}
		} else {
			// At least one player must not have enough cards left in their deck to recurse.
			// The winner of the round is the player with the higher-value card.
			if p1Card > p2Card {
				g.players[0].AddCards(p1Card, p2Card)
			} else {
				g.players[1].AddCards(p2Card, p1Card)
			}
		}

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
