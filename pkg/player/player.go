/* player

Player decision making and state

*/
package player

import (
	"github.com/twschum/hanabi/pkg/card"
	"github.com/twschum/hanabi/pkg/board"
	"fmt"
)

type Player struct {
	Id int
	Cards []*card.Card
	// channel here?
}

func (p Player) ChooseAction(b *board.Board, otherPlayers []Player) {
	// Try to play a card
	for i, c := range(p.Cards) {
		if p.canPlay(c) {
			fmt.Println("Playing ", c)
			b.PlayCard(c)
			p.drawCard(i, b)
			return
		}
	}
	// cheat, check their own playable cards
	if c, valid := p.choosePlayableCard(p.getPlayableCards(b)); valid && false {
		fmt.Println("(CHEAT)Playing ", c)
		b.PlayCard(c)
		p.drawCard(p.indexFromCard(c), b)
		return
	}

	// try to give information
	for _, player := range otherPlayers {
		if player.Id == p.Id {
			continue
		}
	}

	// if nothing else, must discard
	p.randomDiscard(b)
	fmt.Println("Pass")
}

func (p Player) canPlay(c *card.Card) bool {
	// self
	if c.IsValid() && (c.Color_known || c.Number_known) {
		return true;
	}
	return false;
}

func (p Player) drawCard(index int, b *board.Board) {
	// self
	p.Cards[index] = b.Deck.Draw()
	fmt.Println("Drawing ", p.Cards[index])
}

func (p Player) randomDiscard(b *board.Board) {
	// self
	// random, for now 1
	index := 1
	b.DiscardCard(p.Cards[index])
	p.drawCard(index, b)
}

func (p Player) informationKnown() (count int) {
	// any
	for _, c := range p.Cards {
		if c.Number_known {
			count++
		}
		if c.Color_known {
			count++
		}
	}
	return
}

func (p Player) getPlayableCards(b *board.Board) (playable []*card.Card) {
	// other
	for _, c := range p.Cards {
		if c.Number == b.Fireworks[c.Color] + 1 {
			playable = append(playable, c)
		}
	}
	return
}

func (p Player) choosePlayableCard(playable []*card.Card) (*card.Card, bool) {
	// given playable cards, choose the best one
	if len(playable) == 0 {
		return nil, false
	}
	for _, c := range playable {
		if c.Number == 5 {
			return c, true
		}
	}
	for _, c := range playable {
		if c.Number == 3 {
			return c, true
		}
	}
	return playable[0], true
}

func (p Player) indexFromCard(card *card.Card) (int) {
	// for a player get index of card pointer
	for i, c := range p.Cards {
		if *card == *c {
			return i//, true
		}
	}
	return 0
	//return 0, false
}
