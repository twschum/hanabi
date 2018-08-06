/* hanabi

An implementation of a simulation of the game Hanabi, by Antoine Bauza

*/
package main

import (
	"github.com/twschum/hanabi/pkg/simulation"
)

func main() {
	game := simulation.Begin(3)
	simulation.Run(game)
}
