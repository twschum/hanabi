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
	// try to give information
	// TODO get players going clockwise from this player
	// maybe secondary sort based on how much information they have
	for _, player := range otherPlayers {
		if player.Id == p.Id {
			continue
		}
		// TODO ficure out dat logic yo
		// first sort and choose players
		//   least amount of information first
		//   followed by going to the next player
		// Choose a playable card
		//   prefer cards nothing is known about
		//   try to figure out multi-card clues eventually
		if s, valid := player.choosePlayableCard(player.getPlayableCards(b)); valid {
			if c.ColorKnown {
			}
			return
		}
	}

	// if nothing else, must discard
	p.randomDiscard(b)
	fmt.Println("Pass")
}

func (p Player) HandIsEmpty() bool {
	return len(p.Cards) == 0
}

func (p Player) canPlay(c *card.Card) bool {
	// self
	if c.Color_known || c.Number_known {
		return true;
	}
	return false;
}

func (p Player) drawCard(index int, b *board.Board) {
	// self
	if c, valid := b.Deck.Draw(); valid {
		p.Cards[index] = c
		fmt.Println("Drawing ", p.Cards[index])
	} else {
		p.Cards = append(p.Cards[:index], p.Cards[index+1:]...)
	}
}

func (p Player) randomDiscard(b *board.Board) {
	// self
	// random, for now 1
	for i, c := range p.Cards {
		b.DiscardCard(c)
		p.drawCard(i, b)
		return
	}
	fmt.Println("ERROR: no valid cards to discard")
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
		if c.Number == (b.Fireworks[c.Color] + 1) && !(c.Number_known || c.Color_known) {
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
