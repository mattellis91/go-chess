package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	WIDTH       = 512
	HEIGHT      = 512
	DIMENSIONS  = 8
	SQUARE_SIZE = HEIGHT / DIMENSIONS
)

var pieceImages = map[string]*ebiten.Image{}
var whiteSquareColor = color.RGBA{238, 238, 210, 255}
var BlackSquareColor = color.RGBA{118, 150, 86, 255}
var selectedPieceSquareColor = color.RGBA{255, 255, 0, 50}
var higlightedSquareColor = color.RGBA{255, 0, 0, 50}

type Game struct {
	GameState *GameState
}
type Square struct {
	row int
	col int
}

func (g *Game) Update() error {
	handleInput(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBoard(screen)
	drawPieces(screen, g.GameState)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func (g *Game) Init() {
	loadAssets()
	g.GameState.ValidMoves = g.GameState.GetValidMoves()
}

func handleInput(g *Game) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		fmt.Println("Mouse button pressed")
		mouseX, mouseY := ebiten.CursorPosition()
		row := mouseY / SQUARE_SIZE
		col := mouseX / SQUARE_SIZE

		if g.GameState.SquareSelected.row == row && g.GameState.SquareSelected.col == col {
			resetClicks(g.GameState)
		} else {
			g.GameState.SquareSelected = Square{row, col}
			g.GameState.PlayerClicks = append(g.GameState.PlayerClicks, g.GameState.SquareSelected)
		}

		if len(g.GameState.PlayerClicks) == 2 {
			m := NewMove(g.GameState.PlayerClicks[0], g.GameState.PlayerClicks[1], g.GameState.Board, false, false)

			for _, move := range g.GameState.ValidMoves {
				if  m.MoveId == move.MoveId {
					g.GameState.MakeMove(move)
					g.GameState.MoveMade = true
					resetClicks(g.GameState)
					break
				}
			}

			if !g.GameState.MoveMade {
				g.GameState.PlayerClicks = []Square{g.GameState.SquareSelected}
			}
			
			fmt.Println(m.GetChessNotation())
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		mouseX, mouseY := ebiten.CursorPosition()
		row := mouseY / SQUARE_SIZE
		col := mouseX / SQUARE_SIZE
		selectedSquare := Square{row, col}
		if !g.GameState.SquareAlreadyHighlighted(selectedSquare) {
			g.GameState.HiglightedSquares = append(g.GameState.HiglightedSquares, selectedSquare)
		} else {
			g.GameState.HiglightedSquares = g.GameState.RemoveSquareFromSlice(g.GameState.HiglightedSquares, selectedSquare)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.GameState.UndoMove()
		g.GameState.MoveMade = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.GameState.HiglightedSquares = []Square{}
	}

	if g.GameState.MoveMade {
		g.GameState.ValidMoves = g.GameState.GetValidMoves()
		g.GameState.MoveMade = false
	}
}

func loadAssets() {
	imagesToLoad := []string{"wp", "wR", "wN", "wB", "wQ", "wK", "bp", "bR", "bN", "bB", "bQ", "bK"}
	for _, image := range imagesToLoad {
		img, _, err := ebitenutil.NewImageFromFile("assets/" + image + ".png")
		if err != nil {
			log.Fatalf("Error loading image: %v", err)
		}
		pieceImages[image] = img
	}
}

func resetClicks(gs *GameState) {
	gs.SquareSelected = GetNullSquare()
	gs.PlayerClicks = []Square{}
}

func drawBoard(screen *ebiten.Image) {
	for r := 0; r < DIMENSIONS; r++ {
		for c := 0; c < DIMENSIONS; c++ {
			if (r+c)%2 == 0 {
				vector.DrawFilledRect(screen, float32(c*SQUARE_SIZE), float32(r*SQUARE_SIZE), float32(SQUARE_SIZE), float32(SQUARE_SIZE), whiteSquareColor, false)
			} else {
				vector.DrawFilledRect(screen, float32(c*SQUARE_SIZE), float32(r*SQUARE_SIZE), float32(SQUARE_SIZE), float32(SQUARE_SIZE), BlackSquareColor, false)
			}
		}
	}
}

func drawPieces(screen *ebiten.Image, gs *GameState) {

	if gs.SquareSelected.row != -1 && gs.SquareSelected.col != -1 {
		vector.DrawFilledRect(screen, float32(gs.SquareSelected.col*SQUARE_SIZE), float32(gs.SquareSelected.row*SQUARE_SIZE), float32(SQUARE_SIZE), float32(SQUARE_SIZE), selectedPieceSquareColor, false)
	}

	for _, square := range gs.HiglightedSquares {
		vector.DrawFilledRect(screen, float32(square.col*SQUARE_SIZE), float32(square.row*SQUARE_SIZE), float32(SQUARE_SIZE), float32(SQUARE_SIZE), higlightedSquareColor, false)
	}

	for r := 0; r < DIMENSIONS; r++ {
		for c := 0; c < DIMENSIONS; c++ {
			piece := gs.Board[r][c]
			if piece != "--" {
				img := pieceImages[piece]
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(c*SQUARE_SIZE), float64(r*SQUARE_SIZE))
				screen.DrawImage(img, op)
			}
		}
	}
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	gs := NewGameState()
	g := &Game{gs}
	g.Init()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
