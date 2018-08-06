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
	Cards []*card.Card
	// channel here?
}

func (p *Player) ChooseAction(b *board.Board) {
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

	// if nothing else, must discard
	p.randomDiscard(b)
	fmt.Println("Pass")
}

func (p *Player) canPlay(c *card.Card) bool {
	if c.IsValid() && (c.Color_known || c.Number_known) {
		return true;
	}
	return false;
}

func (p *Player) drawCard(index int, b *board.Board) {
	p.Cards[index] = b.Deck.Draw()
	fmt.Println("Drawing ", p.Cards[index])
}

func (p *Player) randomDiscard(b *board.Board) {
	// random, for now 1
	index := 1
	b.DiscardCard(p.Cards[index])
	p.drawCard(index, b)
}
