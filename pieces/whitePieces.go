package pieces

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// 4 pixel border
// 16x16 pixel squares

//TODO: Convert structs to generic piece structs instead of white and black

type Pawn struct {
	Id         PieceId
	PieceImage *ebiten.Image
	DrawOps    *ebiten.DrawImageOptions
}

func NewPawn(piecesSource *ebiten.Image) *Pawn {
	newPiece := &Pawn{
		Id: WHITE_PAWN,
	}
	rect := image.Rect(0, 0, 16, 16)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Pawn) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *Pawn) Draw(screen *ebiten.Image, square Square) {
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
		yOffset = -2
	}

	p.DrawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64((square.Y*16 + yOffset)))
	screen.DrawImage(p.PieceImage, p.DrawOps)
}

type Bishop struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewBishop(piecesSource *ebiten.Image) *Bishop {
	newPiece := &Bishop{
		Id: WHITE_BISHOP,
	}
	rect := image.Rect(32, 16, 48, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Bishop) Draw(screen *ebiten.Image, square Square) {
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
	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64(square.Y*16+yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

func (p *Bishop) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

type Knight struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewKnight(piecesSource *ebiten.Image) *Knight {
	newPiece := &Knight{
		Id: WHITE_KNIGHT,
	}
	rect := image.Rect(16, 16, 32, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Knight) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *Knight) Draw(screen *ebiten.Image, square Square) {
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
	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64(square.Y*16+yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

type Rook struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewRook(piecesSource *ebiten.Image) *Rook {
	newPiece := &Rook{
		Id: WHITE_ROOK,
	}
	rect := image.Rect(0, 16, 16, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Rook) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *Rook) Draw(screen *ebiten.Image, square Square) {
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
	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64(square.Y*16+yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

type Queen struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewQueen(piecesSource *ebiten.Image) *Queen {
	newPiece := &Queen{
		Id: WHITE_QUEEN,
	}
	rect := image.Rect(44, 16, 60, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Queen) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	yOffset := 0
	switch square.Y {
	case 0:
		yOffset = 5
	case 7:
		yOffset = -2
	}
	drawOps.GeoM.Translate(float64(square.X*16), float64(square.Y*16+yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}

func (p *Queen) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

type King struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewKing(piecesSource *ebiten.Image) *King {
	newPiece := &King{
		Id: WHITE_KING,
	}
	rect := image.Rect(62, 16, 74, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *King) GetLegalMoves(position BoardPosition, square Square) []int {
	return nil
}

func (p *King) Draw(screen *ebiten.Image, square Square) {
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
	drawOps.GeoM.Translate(float64(square.X*16+xOffset), float64(square.Y*16+yOffset))
	screen.DrawImage(p.PieceImage, drawOps)
}
