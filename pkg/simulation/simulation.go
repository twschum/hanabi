/* simulation

Board, and general rules

*/
package simulation

import (
	"github.com/twschum/hanabi/pkg/player"
	"github.com/twschum/hanabi/pkg/card"
	"fmt"
)

type Board struct {
	strikes, information int
	//draw // one array for deck, everything else is a slice or pointer
	//discard
	board [card.SuitCount]int // starts at 0 for each suit
}

func Begin(players int) *Board {
	board := new(Board)
	board.players = players
	board.information = 9
	return board
}

/* Run the thing */
func Run(board *Board) {

	// for each player
		// choose an action
		// do that action
		// check end game
}

func (board *Board) Print() {
	// Prints out the game information
	fmt.Printf("R B G Y W\n")
	for i := range board.board {
		fmt.Printf("%d ", board.board[i])
	}
	fmt.Println("")
}

func PrintCards(cards []card.Card) {
	var format string
	for i := range cards {
		if cards[i].Color_known {
			format = "*%c*"
		} else {
			format = " %c "
		}
		fmt.Printf(format, card.ColorChar(cards[i].Color))
	}
	fmt.Println("")
	for i := range cards {
		if cards[i].Number_known {
			format = "*%d*"
		} else {
			format = " %d "
		}
		fmt.Printf(format, cards[i].Number)
	}
	fmt.Println("")
}
