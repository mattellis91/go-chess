package pieces

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 4 pixel border
// 16x16 pixel squares

//TODO: Convert structs to generic piece structs instead of white and black

type Pawn struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewPawn(piecesSource *ebiten.Image, pieceId PieceId) *Pawn {
	newPiece := &Pawn{
		Id: pieceId,
	}
	rect := image.Rect(0, 0, 16, 16)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Pawn) GetLegalMoves(position BoardPosition, square Square) []Square {
	startingY := 6
	multiplier := 1
	if p.Id < 0 { // Black pawn
		startingY = 1
		multiplier = -1
	}
	legalMoves := []Square{}
	if square.Y == startingY {
		legalMoves = append(legalMoves, Square{square.X, square.Y - (2 * multiplier)}, Square{square.X, square.Y - (1 * multiplier)})	
	} else {
		legalMoves = append(legalMoves, Square{square.X, square.Y - (1 * multiplier)})
	}
	return legalMoves
}

func (p *Pawn) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
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

	drawOps.GeoM.Translate(float64((square.X*16 + xOffset)), float64((square.Y*16 + yOffset)))
	screen.DrawImage(p.PieceImage, drawOps)
}

type Bishop struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewBishop(piecesSource *ebiten.Image, pieceId PieceId) *Bishop {
	newPiece := &Bishop{
		Id: pieceId,
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

func (p *Bishop) GetLegalMoves(position BoardPosition, square Square) []Square {
	return nil
}

type Knight struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewKnight(piecesSource *ebiten.Image, pieceId PieceId) *Knight {
	newPiece := &Knight{
		Id: pieceId,
	}
	rect := image.Rect(16, 16, 32, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Knight) GetLegalMoves(position BoardPosition, square Square) []Square {
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

func NewRook(piecesSource *ebiten.Image, pieceId PieceId) *Rook {
	newPiece := &Rook{
		Id: pieceId,
	}
	rect := image.Rect(0, 16, 16, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *Rook) GetLegalMoves(position BoardPosition, square Square) []Square {
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

func NewQueen(piecesSource *ebiten.Image, pieceId PieceId) *Queen {
	newPiece := &Queen{
		Id: pieceId,
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

func (p *Queen) GetLegalMoves(position BoardPosition, square Square) []Square {
	return nil
}

type King struct {
	Id         PieceId
	PieceImage *ebiten.Image
}

func NewKing(piecesSource *ebiten.Image, pieceId PieceId) *King {
	newPiece := &King{
		Id: pieceId,
	}
	rect := image.Rect(62, 16, 74, 32)
	newPiece.PieceImage = piecesSource.SubImage(rect).(*ebiten.Image)
	return newPiece
}

func (p *King) GetLegalMoves(position BoardPosition, square Square) []Square {
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

type LegalSquareIndicator struct {
	IndicatorImage *ebiten.Image
}

func NewLegalSquareIndicator() *LegalSquareIndicator {
	newIndicator := &LegalSquareIndicator{}
	var err error
	indicatorImg, _, err := ebitenutil.NewImageFromFile("assets/set_regular/circle.png")
	if err != nil {
		log.Fatal(err)
	}
	newIndicator.IndicatorImage = indicatorImg
	return newIndicator
}

func (l *LegalSquareIndicator) Draw(screen *ebiten.Image, square Square) {
	drawOps := &ebiten.DrawImageOptions{}
	drawOps.GeoM.Translate(float64(square.X*16), float64(square.Y*16))
	screen.DrawImage(l.IndicatorImage, drawOps)
}
