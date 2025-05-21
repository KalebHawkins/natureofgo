package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 240
)

type Game struct {
	BgImg *ebiten.Image

	NormalRandom float64
	Mean         float64
	StdDeviation float64
}

func (g *Game) Update() error {
	g.NormalRandom = rand.NormFloat64()*g.StdDeviation + g.Mean
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.BgImg, nil)
	vector.DrawFilledCircle(g.BgImg, float32(g.NormalRandom), ScreenHeight/2, 8, color.RGBA{0, 0, 0, 8}, true)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	bgImg := ebiten.NewImage(ScreenWidth, ScreenHeight)
	bgImg.Fill(color.White)

	g := &Game{
		BgImg:        bgImg,
		Mean:         320,
		StdDeviation: 60,
	}

	return g
}

func main() {
	ebiten.SetWindowTitle("A Gaussian Distribution")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
