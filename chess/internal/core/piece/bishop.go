package piece

import "fmt"

type Bishop struct {
	Piece
}

func (bishop *Bishop) CanMove(board Board, start Position, end Position) bool {
	fmt.Println("Moving bishop from ", start, "to ", end)
	return bishop.CanMoveDiagonally(board, start, end)
}

func (bishop *Bishop) GetColor() Color {
	return bishop.Color
}

func (bishop *Bishop) GetType() ChessPieceType {
	return bishop.Type
}