package pieces

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PieceId int
type PieceColor rune
type BoardPosition [8][8]PieceId

type Square struct {
	X int
	Y int
}

type DragPosition struct {
	X int
	Y int
}

type GamePiece interface {
	Draw(screen *ebiten.Image, square Square)
	GetLegalMoves(position BoardPosition, square Square) []Square
}

const (
	EMPTY      = 0
	WHITE_PAWN = iota
	WHITE_KNIGHT
	WHITE_BISHOP
	WHITE_ROOK
	WHITE_QUEEN
	WHITE_KING
)

const (
	BLACK_KING = iota - 6
	BLACK_QUEEN
	BLACK_ROOK
	BLACK_BISHOP
	BLACK_KNIGHT
	BLACK_PAWN
)

const DRAG_START_LOCATION = -10

const WHITE_PIECES = 'w'
const BLACK_PIECES = 'b'
