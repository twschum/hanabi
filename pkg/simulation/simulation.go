/* simulation

Top level simulation

*/
package simulation

import (
	"github.com/twschum/hanabi/pkg/card"
	"github.com/twschum/hanabi/pkg/board"
	"github.com/twschum/hanabi/pkg/player"
	"fmt"
)

type Game struct {
	GameBoard board.Board
	Players []player.Player
}

func Begin(nPlayers int) *Game {
	game := new(Game)
	game.Players = make([]player.Player, nPlayers)
	game.GameBoard.Information = 9
	for i := 0; i < nPlayers*4; i++ {
		game.Players[i % nPlayers].Cards = append(game.Players[i % nPlayers].Cards, card.DrawCard())
	}
	return game
}

/* Run the thing */
func Run(game *Game) {
	// for each player
	for game.GameBoard.Strikes < 3 {
		for i, p := range game.Players {
			// choose&do an action
			PrintCards(p.Cards)
			fmt.Printf("Player %d: ", i)
			p.ChooseAction(&game.GameBoard)
		}
		game.GameBoard.Print()
		game.GameBoard.Strikes++
	}
	fmt.Println("Game Over - 3 Strikes and You're Out!")
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
