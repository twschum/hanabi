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
	Deck card.Deck
	Discards []*card.Card
	Fireworks [card.SuitCount]int // starts at 0 for each suit
}

func (board *Board) PlayCard(c *card.Card) bool {
	// use closure to reprint board??
	// returns success or not
	if c.Number == board.Fireworks[c.Color] + 1 {
		board.Fireworks[c.Color] = c.Number
	} else {
		// card to discards
		board.Discards = append(board.Discards, c)
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

func (board *Board) DiscardCard(c *card.Card) {
	// whatever traking we may do for cards
	fmt.Println("Discarded ", c)
	board.Information++
	board.Discards = append(board.Discards, c)
}

func (board *Board) Print() {
	// Prints out the game information
	fmt.Printf("-------------\n| R B G Y W | Info: %d  Strikes: %d\n| ", board.Information, board.Strikes)
	for _, v := range board.Fireworks {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("| Score: %d / %d\n-------------\n", board.Score(), board.MaxScore())
}

func (board *Board) Score() (score int) {
	for _, v := range board.Fireworks {
		score += v
	}
	return
}

func (board *Board) ShouldContinue() bool {
	// checks against end game conditions
	// 3 strikes
	if board.Strikes == 3 {
		fmt.Println("Game Over - 3 Strikes and You're Out!")
		return false
	}
	// score is maximized based on discards
	if board.Score() == board.MaxScore() {
		fmt.Println("Game Over - No more playable cards")
		return false
	}
	// a player is out of cards and the deck is out of cards
	return true
}

func (board *Board) MaxScore() (sum int) {
	maxFireworks := []int{5, 5, 5, 5, 5} // ???
	// for each color, determine max
	for color := range maxFireworks {
		// get counts of each number for this color
		counts := map[int]int{}
		for _, c := range board.Discards {
			if c.Color == color {
				counts[c.Number]++
			}
		}
		// for the counts figure out the max firework
		for i := 1; i < len(distribution); i++ {
			if counts[i] == distribution[i] {
				maxFireworks[color] = i - 1
				break
			}
		}
	}
	// tally the score
	for _, v := range maxFireworks {
		sum += v
	}
	return
}

