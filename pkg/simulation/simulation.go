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
	Board board.Board
	Players []player.Player
}

func Begin(nPlayers int) *Game {
	game := new(Game)
	game.Players = make([]player.Player, nPlayers)
	game.Board.Information = 9
	// deal out starting hand cards
	game.Board.Deck = card.GenerateDeck()
	for i := 0; i < nPlayers*4; i++ {
		game.Players[i % nPlayers].Cards = append(game.Players[i % nPlayers].Cards, game.Board.Deck.Draw())
	}
	// assign IDs
	for i, player := range game.Players {
		player.Id = i
	}
	return game
}

/* Run the thing */
func Run(game *Game) {
	// for each player
	for game.Board.Strikes < 3 {
		for i, p := range game.Players {
			// choose&do an action
			PrintCards(p.Cards)
			fmt.Printf("Player %d: ", i)
			p.ChooseAction(&game.Board, game.Players)
			game.Board.Print()
		}
		game.Board.Strikes++
	}
	fmt.Println("Final Score: ", game.Board.Score())
	fmt.Println("Game Over - 3 Strikes and You're Out!")
}

func PrintCards(cards []*card.Card) {
	for i := range cards {
		format := getFormatter(cards[i].Color_known, "%c")
		fmt.Printf(format, card.ColorChar(cards[i].Color))
	}
	fmt.Println("")
	for i := range cards {
		format := getFormatter(cards[i].Number_known, "%d")
		fmt.Printf(format, cards[i].Number)
	}
	fmt.Println("")
}
func getFormatter(known bool, fmtChar string) string {
	if known {
		return fmt.Sprintf("*%s*", fmtChar)
	} else {
		return fmt.Sprintf(" %s ", fmtChar)
	}
}
