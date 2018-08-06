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
	for i := 0; i < len(card.Deck) && i < nPlayers*4; i++ {
	}
	return game
}

/* Run the thing */
func Run(game *Game) {
	// for each player
	for i, p := range game.Players {
		// choose&do an action
		fmt.Printf("Player %d: ", i)
		p.ChooseAction(&game.GameBoard)
		// check end game
		if game.GameBoard.Strikes == 3 {
			fmt.Println("Game Over - 3 Strikes and You're Out!")
			break
		}
		game.GameBoard.Print()
	}
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
