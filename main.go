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

var whitePawn *pieces.Pawn
var whiteBishop *pieces.Bishop
var whiteKnight *pieces.Knight
var whiteRook *pieces.Rook
var whiteQueen *pieces.Queen
var whiteKing *pieces.King

var blackPawn *pieces.Pawn
var blackBishop *pieces.Bishop
var blackKnight *pieces.Knight
var blackRook *pieces.Rook
var blackQueen *pieces.Queen
var blackKing *pieces.King

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

var legalSquares []pieces.Square

var currentBoard = startBoard

type Game struct {
	strokes map[*Stroke]struct{}
}

func init() {
	var err error
	boardImg, _, err = ebitenutil.NewImageFromFile("assets/set_regular/board_empty.png")
	if err != nil {
		log.Fatal(err)
	}

	whitePiecesImg, _, err = ebitenutil.NewImageFromFile("assets/set_regular/pieces_white_2.png")
	if err != nil {
		log.Fatal(err)
	}

	blackPiecesImg, _, err = ebitenutil.NewImageFromFile("assets/set_regular/pieces_black_2.png")
	if err != nil {
		log.Fatal(err)
	}

	whitePawn = pieces.NewPawn(whitePiecesImg, pieces.WHITE_PAWN)
	whiteBishop = pieces.NewBishop(whitePiecesImg, pieces.WHITE_BISHOP)
	whiteKnight = pieces.NewKnight(whitePiecesImg, pieces.WHITE_KNIGHT)
	whiteRook = pieces.NewRook(whitePiecesImg, pieces.WHITE_ROOK)
	whiteQueen = pieces.NewQueen(whitePiecesImg, pieces.WHITE_QUEEN)
	whiteKing = pieces.NewKing(whitePiecesImg, pieces.WHITE_KING)

	blackPawn = pieces.NewPawn(blackPiecesImg, pieces.BLACK_PAWN)
	blackBishop = pieces.NewBishop(blackPiecesImg, pieces.BLACK_BISHOP)
	blackKnight = pieces.NewKnight(blackPiecesImg, pieces.BLACK_KNIGHT)
	blackRook = pieces.NewRook(blackPiecesImg, pieces.BLACK_ROOK)
	blackQueen = pieces.NewQueen(blackPiecesImg, pieces.BLACK_QUEEN)
	blackKing = pieces.NewKing(blackPiecesImg, pieces.BLACK_KING)

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

	for _, square := range legalSquares {
		indicator := pieces.NewLegalSquareIndicator()
		indicator.Draw(screen, square)
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
	case pieces.BLACK_PAWN:
		return blackPawn
	case pieces.BLACK_BISHOP:
		return blackBishop
	case pieces.BLACK_KNIGHT:
		return blackKnight
	case pieces.BLACK_ROOK:
		return blackRook
	case pieces.BLACK_QUEEN:
		return blackQueen
	case pieces.BLACK_KING:
		return blackKing
	}
	return nil
}

type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

type MouseStrokeSource struct{}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

type Stroke struct {
	source         StrokeSource
	initX          int
	initY          int
	currentX       int
	currentY       int
	released       bool
	draggingObject interface{}
}

func NewStroke(source StrokeSource) *Stroke {
	cx, cy := source.Position()
	return &Stroke{
		source:   source,
		initX:    cx,
		initY:    cy,
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
	legalSquares = getPieceFromId(selectedPiece).GetLegalMoves(currentBoard, pieces.Square{X: cellX, Y: cellY})
	return whitePawn
}

func convertToBoardPosition(x, y int) (int, int) {
	return x / pieceSize, y / pieceSize
}
