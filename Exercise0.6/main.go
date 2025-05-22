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

type Walker struct {
	x, y float64
}

func (w *Walker) Update() {
	var step float64 = 5
	xStep := acceptReject() * step
	if rand.IntN(2) == 0 {
		xStep *= -1
	}
	yStep := acceptReject() * step
	if rand.IntN(2) == 0 {
		xStep *= -1
	}

	w.x += xStep
	w.y += yStep
}

func (w *Walker) Draw(dst *ebiten.Image) {
	vector.DrawFilledCircle(dst, float32(w.x), float32(w.y), 2, color.Black, true)
}

type Game struct {
	bgImg *ebiten.Image
	*Walker
}

func (g *Game) Update() error {
	g.Walker.Update()
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
		bgImg: ebiten.NewImage(ScreenWidth, ScreenHeight),
		Walker: &Walker{
			x: ScreenWidth / 2,
			y: ScreenHeight / 2,
		},
	}
	g.bgImg.Fill(color.White)
	return g
}

func main() {
	ebiten.SetWindowTitle("Custom Distribution Walker")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func acceptReject() float64 {
	for {
		r1 := rand.Float64()
		probability := r1 * r1
		r2 := rand.Float64()
		if r2 < probability {
			return r1
		}
	}
}
