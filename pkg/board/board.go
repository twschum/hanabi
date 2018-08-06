/* board

Board

*/
package board

import (
	"github.com/twschum/hanabi/pkg/card"
	"fmt"
)

type Board struct {
	Strikes, Information int
	//draw // one array for deck, everything else is a slice or pointer
	//discard
	Fireworks [card.SuitCount]int // starts at 0 for each suit
}

func (board *Board) PlayCard(c *card.Card) bool {
	// use closure to reprint board??
	// returns success or not
	if c.Number == board.Fireworks[c.Color] + 1 {
		board.Fireworks[c.Color] = c.Number
	} else {
		// card to discards
		board.Strikes++
		fmt.Println("Failed to play ", c)
		return false
	}
	fmt.Println("Successfully played ", c)
	if c.Number == 3 || c.Number == 5 {
		board.Information++
	}
	return true
}

func (board *Board) Print() {
	// Prints out the game information
	fmt.Printf("R B G Y W\n")
	for i := range board.Fireworks {
		fmt.Printf("%d ", board.Fireworks[i])
	}
	fmt.Println("")
}

