package main

//assets from
//https://brysiaa.itch.io/pixel-chess-assets-pack

import (
	"fmt"
	"log"

	"github.com/mattellis91/go-chess/pieces"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var boardImg *ebiten.Image
var whitePiecesImg *ebiten.Image
var blackPiecesImg *ebiten.Image

const scale = 6
const boardPixelSize = 128
const pieceSize = 16
const windowSize = boardPixelSize * scale

var whitePawn *pieces.WhitePawn
var whiteBishop *pieces.WhiteBishop
var whiteKnight *pieces.WhiteKnight
var whiteRook *pieces.WhiteRook
var whiteQueen *pieces.WhiteQueen
var whiteKing *pieces.WhiteKing

var startBoard = pieces.BoardPosition{
	{-4, -2, -3, -5, -6, -3, -2, -4},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{4, 2, 3, 5, 6, 3, 2, 4},
}
type Game struct {

}

func init() {
	var err error
	boardImg, _, err = ebitenutil.NewImageFromFile("assets/set_wooden/board_empty_wooden.png")
	if err != nil {
		log.Fatal(err)
	}

	whitePiecesImg, _, err = ebitenutil.NewImageFromFile("assets/set_wooden/pieces_wooden_light.png")
	if err != nil {
		log.Fatal(err)
	}

	blackPiecesImg, _, err = ebitenutil.NewImageFromFile("assets/set_wooden/pieces_wooden_dark.png")
	if err != nil {
		log.Fatal(err)
	}	

	whitePawn = pieces.NewWhitePawn(whitePiecesImg)
	whiteBishop = pieces.NewWhiteBishop(whitePiecesImg)
	whiteKnight = pieces.NewWhiteKnight(whitePiecesImg)
	whiteRook = pieces.NewWhiteRook(whitePiecesImg)
	whiteQueen = pieces.NewWhiteQueen(whitePiecesImg)
	whiteKing = pieces.NewWhiteKing(whitePiecesImg)

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(boardImg, nil)
	for i, row := range startBoard {
		for j, id := range row {
			if id != 0 {
				piece := getPieceFromId(id)
				if piece != nil {
					piece.Draw(screen, pieces.Square{X: j, Y: i})
				}
			}
		}
	
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardPixelSize, boardPixelSize
}

func main() {
	ebiten.SetWindowSize(windowSize, windowSize)
	ebiten.SetWindowTitle("Go Chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}
}

func getPieceFromId(id pieces.PieceId) pieces.GamePiece {
	switch id {
	case pieces.WHITE_PAWN:
		return whitePawn
	case pieces.WHITE_BISHOP:
		return whiteBishop
	case pieces.WHITE_KNIGHT:
		return whiteKnight
	case pieces.WHITE_ROOK:
		return whiteRook
	case pieces.WHITE_QUEEN:
		return whiteQueen
	case pieces.WHITE_KING:
		return whiteKing
	}
	return nil
}