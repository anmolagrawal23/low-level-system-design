package board

import "fmt"

type Board interface {
	AddJumper(start, end int) error
	GetNextPosition(start, value int) int
	HasWon(pos int) bool
}

type GameBoard struct {
	Size   int
	Jumper map[int]int
}

func NewGameBoard(size int) Board {
	board := &GameBoard{Size: size}
	board.Jumper = make(map[int]int)
	return board
}

func (g *GameBoard) AddJumper(start, end int) error {
	err := g.validateJumper(start, end)
	if err != nil {
		return err
	}
	g.Jumper[start] = end
	fmt.Println("Jumper added ", start, ":", end)
	return nil
}

func (g *GameBoard) validateJumper(start, end int) error {
	if start == end {
		return fmt.Errorf("start and end position cannot be same")
	}

	if _, exists := g.Jumper[start]; exists {
		return fmt.Errorf("jumper with start position already exists")
	}

	if start < 1 || start > g.Size {
		return fmt.Errorf("jumper with start position is out of bounds")
	}

	if end < 1 || end > g.Size {
		return fmt.Errorf("jumper with end position is out of bounds")
	}

	return nil
}

func (g *GameBoard) GetNextPosition(start, value int) int {
	newPos := start + value
	if newPos > g.Size {
		return start
	}

	if next, exists := g.Jumper[newPos]; exists {
		return next
	}
	return newPos
}

func (g *GameBoard) HasWon(pos int) bool {
	return pos == g.Size
}
