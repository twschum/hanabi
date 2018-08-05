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
}
