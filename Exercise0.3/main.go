package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 240
)

type Walker struct {
	*ebiten.Image
	x, y float64
}

func (w *Walker) Step() {
	xstep := rand.Float64()*1.0 - 1.0
	ystep := rand.Float64()*1.0 - 1.0

	percentageToFollowMouse := 0.5
	movetowardMouse := rand.Float64()

	mouseX, mouseY := ebiten.CursorPosition()
	if movetowardMouse < percentageToFollowMouse {
		if w.x < float64(mouseX) {
			w.x++
		} else {
			w.x--
		}
		if w.y < float64(mouseY) {
			w.y++
		} else {
			w.y--
		}
	} else {
		w.x += xstep
		w.y += ystep
	}
}

func (w *Walker) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(w.x, w.y)
	dst.DrawImage(w.Image, op)
}

type Game struct {
	BgImg  *ebiten.Image
	Walker *Walker
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
	bgImage := ebiten.NewImage(ScreenWidth, ScreenHeight)
	bgImage.Fill(color.White)

	walker := &Walker{
		Image: ebiten.NewImage(2, 2),
		x:     ScreenWidth / 2,
		y:     ScreenHeight / 2,
	}
	walker.Image.Fill(color.Black)

	g := &Game{
		BgImg:  bgImage,
		Walker: walker,
	}
	return g
}

func main() {
	ebiten.SetWindowTitle("A Mouse Following Random Walker")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
