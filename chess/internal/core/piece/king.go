package piece

import "fmt"

type King struct {
	Piece
}

func (king *King) CanMove(board Board, start Position, end Position) bool {
	// Castling - TBD

	fmt.Println("Moving King from ", start, "to ", end)

	if king.IsSameColor(board, start, end) {
		return false
	}

	if (start[0] == end[0] && (end[1] == start[1]-1 || end[1] == start[1] + 1)) || (start[1] == end[1] && (end[0] == start[0]-1 || end[0] == start[0] + 1)) {
		return true
	}
	return false
}

func (king *King) GetColor() Color {
	return king.Color
}

func (king *King) GetType() ChessPieceType {
	return king.Type
}