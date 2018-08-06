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

type color int

func ColorChar(c int) byte {
	return "RBGYW*"[c]
}

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

// Only one deck, all card refs are pointers everywhere else
type Deck struct {
	index int
	deck [25]Card
	nullCard Card
}

func GenerateDeck() Deck {
	var deck Deck
	deck.deck = [...]Card {
		Card{Color: Red, Number: 1, Number_known: true},
		Card{Color: Red, Number: 2, Color_known: true},
		Card{Color: Red, Number: 3},
		Card{Color: Red, Number: 4},
		Card{Color: Red, Number: 5},
		Card{Color: Blue, Number: 1, Number_known: true},
		Card{Color: Blue, Number: 2},
		Card{Color: Blue, Number: 3},
		Card{Color: Blue, Number: 4},
		Card{Color: Blue, Number: 5},
		Card{Color: Green, Number: 1, Number_known: true},
		Card{Color: Green, Number: 2},
		Card{Color: Green, Number: 3},
		Card{Color: Green, Number: 4},
		Card{Color: Green, Number: 5},
		Card{Color: Yellow, Number: 1},
		Card{Color: Yellow, Number: 2},
		Card{Color: Yellow, Number: 3},
		Card{Color: Yellow, Number: 4},
		Card{Color: Yellow, Number: 5},
		Card{Color: White, Number: 1},
		Card{Color: White, Number: 2},
		Card{Color: White, Number: 3},
		Card{Color: White, Number: 4},
		Card{Color: White, Number: 5},
	}
	return deck
}

func (deck *Deck) Draw() *Card {
	// TODO handling the deck running out
	if deck.index == len(deck.deck) {
		return &deck.nullCard
	}
	deck.index++
	return &deck.deck[deck.index-1]
}
