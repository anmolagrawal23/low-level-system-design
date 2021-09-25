package main

import (
	"bufio"
	"fmt"
	game2 "github.com/anmolagrawal23/low-level-system-design/chess/internal/core/game"
	"github.com/anmolagrawal23/low-level-system-design/chess/internal/core/piece"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting new game with Player1 as white & Player2 as black")
	game := &game2.Game{}
	board := game.NewGame()
	fmt.Println("Game initialized...")

	turn := 0
	scanner := bufio.NewScanner(os.Stdin)

	for {
		game2.DisplayBoard(board)
		if turn%2 == 0 {
			fmt.Println("Player1 turn play as White:")
		} else {
			fmt.Println("Player2 turn play as Black")
		}

		fmt.Print("Enter your move: ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		move := strings.Split(input, " ")
		if len(move) < 2 {
			break
		}
		start := strings.Split(move[0], ",")
		end := strings.Split(move[1], ",")
		startRank,_ := strconv.Atoi(start[0])
		startFile,_ := strconv.Atoi(start[1])
		endRank,_ := strconv.Atoi(end[0])
		endFile,_ := strconv.Atoi(end[1])

		var validMove bool
		board, validMove = game.NextMove(board, piece.Position{startRank,startFile}, piece.Position{endRank, endFile})
		if !validMove {
			fmt.Println("Invalid move, play again...")
		} else {
			turn++
		}
	}
}
