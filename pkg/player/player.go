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
	Cards []card.Card
	// channel here?
}

func (p *Player) ChooseAction(b *board.Board) {
	for i := range(p.Cards) {
		c := &p.Cards[i]
		if p.canPlay(c) {
			fmt.Println("Playing ", c)
			b.PlayCard(c)
			p.Cards[i] = card.DrawCard()
			fmt.Println("Drawing ", &p.Cards[i])
			return
		}
	}
	fmt.Println("Pass")
}

func (p *Player) canPlay(c *card.Card) bool {
	if c.IsValid() && (c.Color_known || c.Number_known) {
		return true;
	}
	return false;
}
