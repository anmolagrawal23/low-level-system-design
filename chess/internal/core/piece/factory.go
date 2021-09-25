package piece

func NewPiece(pieceType ChessPieceType, color Color) ChessPiece {

	switch pieceType {
	case KING:
		return &King{Piece: Piece{Type: KING, Color: color, IsMoved: false, IsKing: true}}
	case QUEEN:
		return &Queen{Piece: Piece{Type: QUEEN, Color: color, IsMoved: false, IsKing: false}}
	case KNIGHT:
		return &Knight{Piece: Piece{Type: KNIGHT, Color: color, IsMoved: false, IsKing: false}}
	case BISHOP:
		return &Bishop{Piece: Piece{Type: BISHOP, Color: color, IsMoved: false, IsKing: false}}
	case ROOK:
		return &Rook{Piece: Piece{Type: ROOK, Color: color, IsMoved: false, IsKing: false}}
	case PAWN:
		return &Pawn{Piece: Piece{Type: PAWN, Color: color, IsMoved: false, IsKing: false}}
	}

	return nil
}
