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
	for _, c := range(p.Cards) {
		if p.canPlay(&c) {
			fmt.Println("Playing %v", c)
			c.Number = 0 // hack for now
			b.PlayCard(&c)
			return
		}
	}
}

func (p *Player) canPlay(c *card.Card) bool {
	if c.IsValid() && (c.Color_known || c.Number_known) {
		return true;
	}
	return false;
}
