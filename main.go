package main

//assets from
//https://brysiaa.itch.io/pixel-chess-assets-pack

import (
	"fmt"
	"log"

	"github.com/mattellis91/go-chess/pieces"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

var currentBoard = startBoard

type Game struct {
	strokes map[*Stroke]struct{}
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
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		//TODO: BEGIN DRAG / STROKE
		fmt.Println("Mouse button pressed")
		s := NewStroke(&MouseStrokeSource{})
		s.SetDraggingObject(pieceAtPosition(s.initX, s.initY))
		g.strokes[s] = struct{}{}
	}

	for s := range g.strokes {
		g.updateStroke(s)
		if s.IsReleased() {
			delete(g.strokes, s)
		}
	}
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
	if err := ebiten.RunGame(&Game{
		strokes: map[*Stroke]struct{}{},
	}); err != nil {
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

type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

type MouseStrokeSource struct {}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

type Stroke struct {
	source StrokeSource
	initX int
	initY int
	currentX int
	currentY int
	released bool
	draggingObject interface{}
}

func NewStroke(source StrokeSource) *Stroke {
	cx, cy := source.Position()
	return &Stroke{
		source: source,
		initX: cx,
		initY: cy,
		currentX: cx,
		currentY: cy,
	}
}

func (s *Stroke) Update() {
	if s.released {
		return
	}
	if s.source.IsJustReleased() {
		s.released = true
		return
	}
	x, y := s.source.Position()
	s.currentX = x
	s.currentY = y
}

func (s *Stroke) IsReleased() bool {
	return s.released
}

func (s *Stroke) Position() (int, int) {
	return s.currentX, s.currentY
}

func (s *Stroke) PositionDiff() (int, int) {
	dx := s.currentX - s.initX
	dy := s.currentY - s.initY
	return dx, dy
}

func (s *Stroke) DraggingObject() interface{} {
	return s.draggingObject
}

func (s *Stroke) SetDraggingObject(obj interface{}) {
	s.draggingObject = obj
}

func (g *Game) updateStroke(stroke *Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}
	p := stroke.DraggingObject().(pieces.GamePiece)
	if p == nil {
		return
	}
	//TODO: update position of piece
}

func pieceAtPosition(x, y int) pieces.GamePiece {
	fmt.Println("Piece at position", x, y)
	cellX, cellY := convertToBoardPosition(x, y)
	selectedPiece := currentBoard[cellY][cellX]
	fmt.Println("Cell at position", cellX, cellY)
	fmt.Println("Selected piece", selectedPiece)
	return whitePawn
}

func convertToBoardPosition(x, y int) (int, int) {
	return x / pieceSize, y / pieceSize
}