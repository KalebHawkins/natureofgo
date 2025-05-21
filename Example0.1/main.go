package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Walker struct {
	*ebiten.Image
	x, y float64
}

func (w *Walker) Step() {
	dir := rand.Intn(4)

	if dir == 0 {
		w.x++
	} else if dir == 1 {
		w.x--
	} else if dir == 2 {
		w.y++
	} else {
		w.y--
	}
}

func (w *Walker) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(w.x, w.y)
	dst.DrawImage(w.Image, op)
}

type Game struct {
	BgImg *ebiten.Image
	Walker
}

func (g *Game) Update() error {
	g.Walker.Step()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.BgImg, nil)
	g.Walker.Draw(g.BgImg)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	walker := Walker{
		Image: ebiten.NewImage(2, 2),
		x:     320,
		y:     120,
	}
	walker.Image.Fill(color.Black)

	bgImg := ebiten.NewImage(640, 240)
	bgImg.Fill(color.White)

	g := &Game{
		Walker: walker,
		BgImg:  bgImg,
	}

	return g
}

func main() {
	ebiten.SetWindowTitle("A Traditional Random Walk")
	ebiten.SetWindowSize(640, 240)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
