package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
    ScreenWidth = {{.ScreenWidth}}
    ScreenHeight = {{.ScreenHeight}}
)

type Game struct {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	g := &Game{}
	return g
}

func main() {
	ebiten.SetWindowTitle("{{.WindowTitle}}")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
