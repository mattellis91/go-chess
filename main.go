package main

//assets from
//https://brysiaa.itch.io/pixel-chess-assets-pack

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var boardImg *ebiten.Image
const scale = 6

type Game struct {

}

func init() {
	var err error
	boardImg, _, err = ebitenutil.NewImageFromFile("assets/set_wooden/board_empty_wooden.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(boardImg, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 128, 128
}

func main() {
	ebiten.SetWindowSize(128 * scale, 128 * scale)
	ebiten.SetWindowTitle("Go Chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}
}