/* card

Cards and the deck

*/
package card

import (
	"fmt"
)


const (
	Red      =  iota
	Blue     =  iota
	Green    =  iota
	Yellow   =  iota
	White    =  iota
	//Rainbow  =  iota
	SuitCount = iota
)

type Card struct {
	Color, Number int
	Color_known, Number_known bool
}

func (c *Card) IsValid() bool {
	return c.Number > 0
}

func (c *Card) String() string {
	return fmt.Sprintf("%c%d", ColorChar(c.Color), c.Number)
}

func ColorChar(c int) byte {
	return "RBGYW*"[c]
}

func (c *Card) ColorChar() byte {
	return "RBGYW*"[c.Color]
}

// Only one deck, all card refs are pointers everywhere else
type Deck struct {
	index int
	deck []Card
	nullCard Card
}

var Distribution []int = []int{0, 3, 2, 2, 2, 1} // 3 1s, 2 2s 3s 4s, 1 5
func GenerateDeck() Deck {
	var deck Deck
	for color := 0; color < SuitCount; color++ {
		for n := 1; n < 6; n++ {
			for i := 0; i < Distribution[n]; i++ {
				deck.deck = append(deck.deck, Card{Color: color, Number: n})
			}
		}
	}
	return deck
}

func (deck *Deck) Draw() (*Card, bool) {
	if deck.index == len(deck.deck) {
		return nil, false
	}
	deck.index++
	return &deck.deck[deck.index-1], true
}
