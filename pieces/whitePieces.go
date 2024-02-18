package pieces

import (
	"image"
	"github.com/hajimehoshi/ebiten/v2"
)

// 4 pixel border
// 16x16 pixel squares

//TODO: Convert structs to generic piece structs instead of white and black

type WhitePawn struct {
	Id         PieceId
	PieceImage *ebiten.Image
	DrawOps    *ebiten.DrawImageOptions
}

func NewWhitePawn(piecesSource *ebiten.Image) *WhitePawn {
	newPiece := &WhitePawn{
		Id: WHITE_PAWN,
	}
	rect := image.Rect(0, 0, 16, 16)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *WhitePawn) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *WhitePawn) Draw(screen *ebiten.Image, square Square) {
	p.DrawOps = &ebiten.DrawImageOptions{}
	xOffset := 0
	switch square.X {
	case 0:
		xOffset = 3
	case 1, 2:
		xOffset = 2
	case 3:
		xOffset = 1
	case 5:
		xOffset = -1
	case 6:
		xOffset = -2
	case 7:
		xOffset = -3
	}
	yOffset := 0 
	switch square.Y {
		case 1:
			yOffset = 3
		case 6:
			yOffset = - 2
	}

	p.DrawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64((square.Y * 16 + yOffset)))
	screen.DrawImage(p.PieceImage, p.DrawOps)
}

type WhiteBishop struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewWhiteBishop(piecesSource *ebiten.Image) *WhiteBishop {
	newPiece := &WhiteBishop{
		Id: WHITE_BISHOP,
	}
	rect := image.Rect(32, 16, 48, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *WhiteBishop) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	xOffset := 0
	switch square.X {
	case 2:
		xOffset = 4
	case 5:
		xOffset = 1
	}
	yOffset := 0
	switch square.Y {
		case 0:
			yOffset = 5
		case 7:
			yOffset = -2
	}
	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64(square.Y*16 + yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

func (p *WhiteBishop) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

type WhiteKnight struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewWhiteKnight(piecesSource *ebiten.Image) *WhiteKnight {
	newPiece := &WhiteKnight{
		Id: WHITE_KNIGHT,
	}
	rect := image.Rect(16, 16, 32, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *WhiteKnight) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *WhiteKnight) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	xOffset := 0
	switch square.X {
	case 1:
		xOffset = 4
	case 6:
		xOffset = -1
	}
	yOffset := 0
	switch square.Y {
		case 0:
			yOffset = 5
		case 7:
			yOffset = -2
	}
	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64(square.Y*16 + yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

type WhiteRook struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewWhiteRook(piecesSource *ebiten.Image) *WhiteRook {
	newPiece := &WhiteRook{
		Id: WHITE_ROOK,
	}
	rect := image.Rect(0, 16, 16, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *WhiteRook) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *WhiteRook) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	xOffset := 0
	switch square.X {
	case 0:
		xOffset = 3
	case 7:
		xOffset = -3
	}
	yOffset := 0
	switch square.Y {
	case 0:
		yOffset = 5
	case 7:
		yOffset = -2
	}
	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64(square.Y*16 + yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

type WhiteQueen struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewWhiteQueen(piecesSource *ebiten.Image) *WhiteQueen {
	newPiece := &WhiteQueen{
		Id: WHITE_QUEEN,
	}
	rect := image.Rect(44, 16, 60, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *WhiteQueen) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	yOffset := 0
	switch square.Y {
		case 0:
			yOffset = 5
		case 7:
			yOffset = -2
	}
	drawOps.GeoM.Translate(float64(square.X*16), float64(square.Y*16 + yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

func (p *WhiteQueen) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

type WhiteKing struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewWhiteKing(piecesSource *ebiten.Image) *WhiteKing {
	newPiece := &WhiteKing{
		Id: WHITE_KING,
	}
	rect := image.Rect(62, 16, 74, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *WhiteKing) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *WhiteKing) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	xOffset := 0
	switch square.X {
	case 4:
		xOffset = 2
	}
	yOffset := 0
	switch square.Y {
		case 0:
			yOffset = 5
		case 7:
			yOffset = -2
	}
	drawOps.GeoM.Translate(float64(square.X*16 + xOffset), float64(square.Y*16 + yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}
