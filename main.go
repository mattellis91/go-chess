package main

//assets from
//https://brysiaa.itch.io/pixel-chess-assets-pack

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)


type Game struct {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}
}