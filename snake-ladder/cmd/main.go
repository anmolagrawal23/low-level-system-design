package main

import (
	"fmt"
	game "github.com/anmolagrawal23/low-level-system-design/snake-ladder/internal/services/game"
	"github.com/anmolagrawal23/low-level-system-design/snake-ladder/internal/services/player"
	"strconv"
)

func main() {
	players := createPlayers(4)
	fmt.Println("Created players...")

	game := game.NewSnakeLadder(1, 1, 100, 4)
	for _, p := range players {
		err := game.JoinGame(p)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Plyayer", p.GetName(), "joined the game")
		}
	}

	err := game.Start()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Starting the game")
		for !game.Play() {
			fmt.Println("Winner: ", game.GetWinner())
		}
	}
}

func createPlayers(n int) []player.Player {
	playerList := make([]player.Player, 0)
	for i := 1; i <= n; i++ {
		playerName := strconv.Itoa(i)
		playerName = "Player" + playerName
		email := playerName + "@random.com"
		newPlayer := player.NewPlayer(i, playerName, email)
		playerList = append(playerList, newPlayer)
	}
	return playerList
}
