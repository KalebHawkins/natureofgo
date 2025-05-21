package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	randomCounts []int
}

func (g *Game) Update() error {
	idx := rand.Intn(len(g.randomCounts))
	g.randomCounts[idx]++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	w := 640 / len(g.randomCounts)

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
	ebiten.SetWindowTitle("A Random-Number Distribution")
	ebiten.SetWindowSize(640, 240)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
