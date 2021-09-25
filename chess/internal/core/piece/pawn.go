package piece

import (
	"fmt"
)

type Pawn struct {
	Piece
}

func (pawn *Pawn) CanMove(board Board, start Position, end Position) bool {

	fmt.Println("Moving pawn from ", start, "to ", end)

	if pawn.IsSameColor(board, start, end) {
		fmt.Println("Pieces are of same color.")
		return false
	}

	if start[1] == end[1]  {
		endCell := board[end[0]][end[1]]
		if endCell != nil || mod(end[0]-start[0]) > 2 {
			fmt.Println("Cannot play more than 2 moves")
			return false
		} else if mod(end[0] - start[0]) == 2  {
			if pawn.IsMoved {
				fmt.Println("Is already moved, cannot move 2 squares.")
				return false
			}
			if pawn.Color == White && board[start[0]+1][start[1]] != nil {
				fmt.Println("In between square is not empty.")
				return false
			} else if pawn.Color == Black && board[start[0]-1][start[1]] != nil {
				fmt.Println("In between square is not empty.")
				return false
			}
		}
	} else if mod(end[0]-start[0]) != 1 || mod(end[1]-start[1]) != 1 {
		fmt.Println("Not a valid capture.")
		return false
	}
	return true
}

func mod(a int) int {
	if a < 0 {
		return -1*a
	}
	return a
}

func (pawn *Pawn) GetColor() Color {
	return pawn.Color
}

func (pawn *Pawn) GetType() ChessPieceType {
	return pawn.Type
}