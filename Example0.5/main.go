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
	randomCounts []int
}

func (g *Game) Update() error {
	idx := g.acceptReject()
	g.randomCounts[idx]++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	w := ScreenWidth / len(g.randomCounts)

	for i := 0; i < len(g.randomCounts); i++ {
		vector.StrokeRect(screen, float32(i*w), float32(240-g.randomCounts[i]), float32(w-1), float32(g.randomCounts[i]), 2, color.Black, true)
		vector.DrawFilledRect(screen, float32(i*w), float32(240-g.randomCounts[i]), float32(w-1), float32(g.randomCounts[i]), color.RGBA{0, 0, 0, 127}, true)
	}
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	g := &Game{}

	for i := 0; i < 20; i++ {
		g.randomCounts = append(g.randomCounts, 0)
	}

	return g
}

func main() {
	ebiten.SetWindowTitle("Accept-Reject Distribution")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) acceptReject() int {
	for {
		r1 := rand.IntN(len(g.randomCounts))
		probability := r1
		r2 := rand.IntN(len(g.randomCounts))
		if r2 < probability {
			return r1
		}
	}
}
