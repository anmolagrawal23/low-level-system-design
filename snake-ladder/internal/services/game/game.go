package game

import (
	"fmt"
	"github.com/anmolagrawal23/low-level-system-design/snake-ladder/internal/services/board"
	"github.com/anmolagrawal23/low-level-system-design/snake-ladder/internal/services/player"
	"math/rand"
)

const (
	IDLE int = iota
	RUNNING
	OVER
)

type Game interface {
	Play() bool
	GetWinner() player.Player
	JoinGame(player.Player) error
	Start() error
}

type SnakeLadder struct {
	Id         int
	maxPlayers int
	numDices   int
	status     int
	players    []player.Player
	winner     player.Player
	position   map[player.Player]int
	turn       int
	Board      board.Board
}

func (s *SnakeLadder) Play() bool {
	s.turn = (s.turn + 1) % len(s.players)
	score := rollDice(s.numDices)
	currPlayer := s.players[s.turn]
	nextPos := s.Board.GetNextPosition(s.position[currPlayer], score)
	s.position[currPlayer] = nextPos
	fmt.Println("Player ", currPlayer.GetName(), "rolled ", score, ". New position = ", nextPos)

	if s.Board.HasWon(nextPos) {
		s.status = OVER
		s.winner = currPlayer
		fmt.Println("Player ", currPlayer.GetName(), "has won the game")
		return true
	}
	return false
}

func (s *SnakeLadder) GetWinner() player.Player {
	return s.winner
}

func NewSnakeLadder(id int, numDices int, size int, maxPlayers int) Game {
	newBoard := board.NewGameBoard(size)
	newGame := &SnakeLadder{Id: id, numDices: numDices, maxPlayers: maxPlayers}
	newGame.turn = -1
	newGame.Board = newBoard
	newGame.position = make(map[player.Player]int)
	newGame.players = make([]player.Player, 0)
	newGame.addJumpers()
	return newGame
}

func (s *SnakeLadder) JoinGame(player player.Player) error {

	if s.status == RUNNING {
		return fmt.Errorf("cannot join as game has already start")
	}

	if s.status == OVER {
		return fmt.Errorf("cannot join as game is already over")
	}

	if len(s.players) == s.maxPlayers {
		return fmt.Errorf("game is full, cannot join more players")
	}

	if _, found := s.position[player]; found {
		return fmt.Errorf("player has already joined the game")
	}

	fmt.Println("Adding player", player.GetName(), " to the game")
	s.players = append(s.players, player)
	s.position[player] = 0

	return nil
}

func (s *SnakeLadder) Start() error {
	if s.status != IDLE {
		return fmt.Errorf("game has already started or finished")
	}
	return nil
}

func rollDice(n int) int {
	return n + rand.Int()%(5*n+1)
}

func (s *SnakeLadder) addJumpers() {
	s.Board.AddJumper(5, 25)
	s.Board.AddJumper(10, 20)
	s.Board.AddJumper(37, 43)
	s.Board.AddJumper(45, 85)
	s.Board.AddJumper(78, 92)

	s.Board.AddJumper(31, 2)
	s.Board.AddJumper(44, 23)
	s.Board.AddJumper(51, 12)
	s.Board.AddJumper(80, 50)
	s.Board.AddJumper(97, 61)
}
