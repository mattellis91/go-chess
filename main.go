package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	WIDTH = 512
	HEIGHT = 512
	DIMENSIONS = 8
	SQUARE_SIZE = HEIGHT / DIMENSIONS
)

var pieceImages = map[string]*ebiten.Image{}
var whiteSquareColor = color.RGBA{238, 238, 210, 255}
var BlackSquareColor = color.RGBA{118, 150, 86, 255}

type Game struct{
	GameState *GameState
}

func (g *Game) Update() error {
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
	for r := 0; r < DIMENSIONS; r++ {
		for c := 0; c < DIMENSIONS; c++ {
			piece := gs.board[r][c]
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