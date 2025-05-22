package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/Kalebhawkins/natureofgo/ui"
	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 240
)

type Game struct {
	NoiseGenerator *perlin.Perlin
	Pixels         []byte
	Scale          float64
	ScaleSlider    *ui.Slider
	Time           float64
}

func (g *Game) Update() error {
	// Handle the Slider Logic
	mx, my := ebiten.CursorPosition()
	g.ScaleSlider.Update(mx, my, ebiten.IsMouseButtonPressed(ebiten.MouseButton0))
	g.Scale = g.ScaleSlider.Value

	// Perlin Noise determines the brightness of each pixel.
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			index := (y*ScreenWidth + x) * 4

			redNoise := g.NoiseGenerator.Noise3D(float64(x)*g.Scale, float64(y)*g.Scale, g.Time)
			greenNoise := g.NoiseGenerator.Noise3D(float64(x)*g.Scale, float64(y)*g.Scale, g.Time+4)
			blueNoise := g.NoiseGenerator.Noise3D(float64(x)*g.Scale, float64(y)*g.Scale, g.Time+8)

			red := uint8((redNoise + 1) / 2 * 255)
			green := uint8((greenNoise + 1) / 2 * 255)
			blue := uint8((blueNoise + 1) / 2 * 255)

			g.Pixels[index] = red
			g.Pixels[index+1] = green
			g.Pixels[index+2] = blue
			g.Pixels[index+3] = 255
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.Pixels)
	g.ScaleSlider.Draw(screen)

	txtop := &text.DrawOptions{}
	txtop.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, fmt.Sprintf("Scale: %.2f", g.Scale), ui.FontFace, txtop)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	g := &Game{
		NoiseGenerator: perlin.NewPerlin(2, 2, 8, rand.Int64()),
		Pixels:         make([]byte, ScreenWidth*ScreenHeight*4),
		Scale:          0.01,
		ScaleSlider: &ui.Slider{
			Label: "Noise Scale",
			X:     20,
			Y:     200,
			Width: 120,
			Min:   0.001,
			Max:   0.99,
			Value: 0.03,
			Color: color.Black,
		},
	}
	return g
}

func main() {
	ebiten.SetWindowTitle("2D Perlin Noise")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
