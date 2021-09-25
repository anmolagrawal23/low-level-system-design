package piece

import "fmt"

type Knight struct {
	Piece
}

func (knight *Knight) CanMove(board Board, start Position, end Position) bool {

	fmt.Println("Moving Knight from ", start, "to ", end)

	if knight.IsSameColor(board, start, end) {
		return false
	}

	top, down, left, right := start[0]+2, start[0]-2, start[1]-2, start[1]+2

	if isValidEndCell(top, start[1]-1, end) || isValidEndCell(top, start[1]+1, end) || isValidEndCell(down, start[1]-1, end) || isValidEndCell(down, start[1]+1, end) ||
		isValidEndCell(start[0]-1, left, end) || isValidEndCell(start[0]+1, left, end) || isValidEndCell(start[0]-1, right, end) || isValidEndCell(start[0]+1, right, end) {
		return true
	}
	return false
}

func isValidEndCell(x, y int, end Position) bool {
	if x >= 0 && x < 8 && y >= 0 && y < 8 && x == end[0] && y == end[1] {
		return true
	}
	return false
}

func (knight *Knight) GetColor() Color {
	return knight.Color
}

func (knight *Knight) GetType() ChessPieceType {
	return knight.Type
}