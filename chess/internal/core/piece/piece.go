package piece

type ChessPieceType int
type Color	int
type Position [2]int
type Board	[8][8]ChessPiece

const (
	KING ChessPieceType = iota
	QUEEN
	ROOK
	BISHOP
	KNIGHT
	PAWN
)

const (
	White Color = iota
	Black
)

type Piece struct {
	Type		ChessPieceType
	Color		Color
	IsMoved		bool
	IsKing		bool
}

type ChessPiece interface {
	CanMove(board Board, start Position, end Position) bool
	GetColor() Color
	GetType() ChessPieceType
}

func (pc *Piece) IsSameColor(board Board, start Position, end Position) bool {
	startCell, endCell := board[start[0]][start[1]], board[end[0]][end[1]]
	if endCell != nil && endCell.GetColor() == startCell.GetColor() {
		return true
	}
	return false
}

func (pc *Piece) CanMoveStraight(board Board, start Position, end Position) bool {

	if start[0] != end[0] && start[1] != end[1] {
		return false
	}
	if pc.IsSameColor(board, start, end) {
		return false
	}

	if start[0] == end[0] {
		inc := 1
		if end[1] < start[1] {
			inc = -1
		}
		for file := start[1]+inc; file != end[1]; file+=inc{
			if board[start[0]][file] != nil {
				return false
			}
		}
	} else {
		inc := 1
		if end[0] < start[0] {
			inc = -1
		}
		for rank := start[0]+inc; rank != end[0]; rank+=inc{
			if board[rank][start[1]] != nil {
				return false
			}
		}
	}
	return true
}

func (pc *Piece) CanMoveDiagonally(board Board, start Position, end Position) bool {

	if pc.IsSameColor(board, start, end) {
		return false
	}
	if abs(end[0]-start[0]) != abs(end[1]-start[1]) {
		return false
	}

	xinc := 1
	if end[0]-start[0] < 0 {
		xinc = -1
	}
	yinc := 1
	if end[1]-start[1] < 0 {
		yinc = -1
	}

	for rank, file := start[0]+xinc, start[1]+yinc; rank != end[0]; rank+=xinc {
		if board[rank][file] != nil {
			return false
		}
		file+=yinc
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -1*a
	}
	return a
}