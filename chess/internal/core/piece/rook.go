package piece

import "fmt"

type Rook struct {
	Piece
}

func (rook *Rook) CanMove(board Board, start Position, end Position) bool {
	fmt.Println("Moving rook from ", start, "to ", end)
	return rook.CanMoveStraight(board, start, end)
}

func (rook *Rook) GetColor() Color {
	return rook.Color
}

func (rook *Rook) GetType() ChessPieceType {
	return rook.Type
}