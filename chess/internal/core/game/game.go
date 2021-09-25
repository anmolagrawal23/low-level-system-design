package game

import "github.com/anmolagrawal23/low-level-system-design/chess/internal/core/player"

type Status int

const (
	Running Status = iota
	Won
	Draw
	Forfeit
	Resigned
)

type Game struct {
	id					int
	status        		Status
	PlayerWhite   		player.Player
	PlayerBlack   		player.Player
	NotationSheet 		[]string
	isPlayerInCheck		bool
}