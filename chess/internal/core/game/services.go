package game

import (
	"fmt"
	"github.com/anmolagrawal23/low-level-system-design/chess/internal/core/piece"
)

// NewGame - Function to create a new game
func (game *Game) NewGame() piece.Board {
	fmt.Println("Initializing board...")
	board := InitializeBoard()
	return board
}

func InitializeBoard() piece.Board {
	var board piece.Board
	board[0][0] = piece.NewPiece(piece.ROOK, piece.White)
	board[0][1] = piece.NewPiece(piece.KNIGHT, piece.White)
	board[0][2] = piece.NewPiece(piece.BISHOP, piece.White)
	board[0][3] = piece.NewPiece(piece.QUEEN, piece.White)
	board[0][4] = piece.NewPiece(piece.KING, piece.White)
	board[0][5] = piece.NewPiece(piece.BISHOP, piece.White)
	board[0][6] = piece.NewPiece(piece.KNIGHT, piece.White)
	board[0][7] = piece.NewPiece(piece.ROOK, piece.White)

	board[7][0] = piece.NewPiece(piece.ROOK, piece.Black)
	board[7][1] = piece.NewPiece(piece.KNIGHT, piece.Black)
	board[7][2] = piece.NewPiece(piece.BISHOP, piece.Black)
	board[7][3] = piece.NewPiece(piece.QUEEN, piece.Black)
	board[7][4] = piece.NewPiece(piece.KING, piece.Black)
	board[7][5] = piece.NewPiece(piece.BISHOP, piece.Black)
	board[7][6] = piece.NewPiece(piece.KNIGHT, piece.Black)
	board[7][7] = piece.NewPiece(piece.ROOK, piece.Black)

	for i:=0; i<8; i++ {
		board[1][i] = piece.NewPiece(piece.PAWN, piece.White)
		board[6][i] = piece.NewPiece(piece.PAWN, piece.Black)
	}
	return board
}

func DisplayBoard(board piece.Board) {
	fmt.Println("    BLACK    ")
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank)
		for file := 0; file < 8; file++ {
			if board[rank][file] == nil {
				fmt.Printf("- ")
				continue
			}
			switch board[rank][file].GetType() {
			case piece.PAWN:
				if board[rank][file].GetColor() == piece.White {
					fmt.Printf("P ")
				} else {
					fmt.Printf("p ")
				}
			case piece.KING:
				if board[rank][file].GetColor() == piece.White {
					fmt.Printf("K ")
				} else {
					fmt.Printf("k ")
				}
			case piece.QUEEN:
				if board[rank][file].GetColor() == piece.White {
					fmt.Printf("Q ")
				} else {
					fmt.Printf("q ")
				}
			case piece.ROOK:
				if board[rank][file].GetColor() == piece.White {
					fmt.Printf("R ")
				} else {
					fmt.Printf("r ")
				}
			case piece.BISHOP:
				if board[rank][file].GetColor() == piece.White {
					fmt.Printf("B ")
				} else {
					fmt.Printf("b ")
				}
			case piece.KNIGHT:
				if board[rank][file].GetColor() == piece.White {
					fmt.Printf("N ")
				} else {
					fmt.Printf("n ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Printf("  ")
	for file:=0; file<8; file++ {
		fmt.Printf("%d ", file)
	}
	fmt.Println()
	fmt.Println("    WHITE")
}

func (game *Game) NextMove(board piece.Board, start piece.Position, end piece.Position) (piece.Board,bool) {
	if board[start[0]][start[1]].CanMove(board, start, end) {
		board[end[0]][end[1]], board[start[0]][start[1]] = board[start[0]][start[1]], nil
		return board,true
	}
	return board,false
}

// CheckGameStatus - Verify if the game is still running, won, draw
func (game *Game) CheckGameStatus(board piece.Board) bool {
	return false
}
