/* simulation

Board, Card, and general rules

*/
package simulation

import (
	"fmt"
)

const (
	red      =  iota
	blue     =  iota
	green    =  iota
	yellow   =  iota
	white    =  iota
	//rainbow  =  iota
	SuitCount = iota
)

func ColorChar(c int) byte {
	return "RBGYW*"[c]
}

type Card struct {
	Color, Number int
	Color_known, Number_known bool
}

type State struct {
	strikes, information, players int
	//draw // one array for deck, everything else is a slice or pointer
	//discard
	board [SuitCount]int // starts at 0 for each suit
}

func Begin(players int) *State {
	state := new(State)
	state.players = players
	state.information = 9
	return state
}

/* Run the thing */
func Run(state *State) {


}

func PrintBoard(state *State) {
	fmt.Printf("R B G Y W\n")
	for value := range state.board {
		fmt.Printf("%d ", value)
	}
	fmt.Println("")
}

func PrintCards(cards []Card) {
	var format string
	for i := range cards {
		card := cards[i]
		if card.Color_known {
			format = "|%c|"
		} else {
			format = " %c "
		}
		fmt.Printf(format, ColorChar(card.Color))
	}
	fmt.Println("")
	for i := range cards {
		card := cards[i]
		if card.Number_known {
			format = "|%d|"
		} else {
			format = " %d "
		}
		fmt.Printf(format, card.Number)
	}
	fmt.Println("")
}
