package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 240
)

type Walker struct {
	tx, ty float64
	x, y   float64
}

func (w *Walker) Update(NoiseGenerator *perlin.Perlin) {
	stepSize := Map(NoiseGenerator.Noise1D(w.tx), 0, 1, 0, 10)

	w.x += Map(NoiseGenerator.Noise1D(w.tx), 0, 1, 0, stepSize)
	w.y += Map(NoiseGenerator.Noise1D(w.ty), 0, 1, 0, stepSize)

	w.tx += 0.01
	w.ty += 0.01
}

func (w *Walker) Draw(dst *ebiten.Image) {
	vector.DrawFilledCircle(dst, float32(w.x), float32(w.y), 2, color.Black, true)
}

type Game struct {
	bgImg          *ebiten.Image
	NoiseGenerator *perlin.Perlin
	Walker         *Walker
}

func (g *Game) Update() error {
	g.Walker.Update(g.NoiseGenerator)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bgImg, nil)
	g.Walker.Draw(g.bgImg)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	g := &Game{
		bgImg:          ebiten.NewImage(ScreenWidth, ScreenHeight),
		NoiseGenerator: perlin.NewPerlin(2, 2, 3, rand.Int64()),
		Walker: &Walker{
			tx: 0,
			ty: 10000,
			x:  ScreenWidth / 2,
			y:  ScreenHeight / 2,
		},
	}
	g.bgImg.Fill(color.White)
	return g
}

func main() {
	ebiten.SetWindowTitle("A Perlin Noise Walker Using Noise with Step Size")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func Map(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)/(inMax-inMin)*(outMax-outMin) + outMin
}
