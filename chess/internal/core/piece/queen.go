package piece

import "fmt"

type Queen struct {
	Piece
}

func (queen *Queen) CanMove(board Board, start Position, end Position) bool {

	fmt.Println("Moving queen from ", start, "to ", end)
	return queen.CanMoveStraight(board, start, end) || queen.CanMoveDiagonally(board, start, end)
}

func (queen *Queen) GetColor() Color {
	return queen.Color
}

func (queen *Queen) GetType() ChessPieceType {
	return queen.Type
}