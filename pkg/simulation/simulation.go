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

type Card struct {
	color, number int
}

type State struct {
	strikes, information, players int
	//draw // one array for deck, everything else is a slice or pointer
	//discard
	board [SuitCount]Card // starts at 0 for each suit
}

func Begin(players int) State {
	state := make(State)
	state.players = players
	state.information = 9
}

func PrintBoard(state State) {
	fmt.Printf("R B G Y W\n")
	fmt.Printf("%d %d %d %d %d\n", state.board[red], state.board[blue],
		state.board[green], state.board[yellow], state.board[white])
}
