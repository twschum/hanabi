/* hanabi

An implementation of a simulation of the game Hanabi, by Antoine Bauza

*/
package main

import (
	"github.com/twschum/hanabi/pkg/simulation"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world.\n")
	game := simulation.Begin(3)
	simulation.PrintBoard(game)
	cards := make([]simulation.Card, 3)
	cards[0] = simulation.Card{2, 3, false, true}
	cards[1] = simulation.Card{0, 5, true, false}
	cards[2] = simulation.Card{1, 5, false, false}
	simulation.PrintCards(cards)
}
