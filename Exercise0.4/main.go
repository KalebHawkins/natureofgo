package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/Kalebhawkins/natureofgo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 240
)

type Game struct {
	bg                      *ebiten.Image
	meanSlider              *ui.Slider
	standardDeviationSlider *ui.Slider
	mean                    float64
	standardDeviation       float64
	x, y                    float64
	splatterRadius          float64
	splatterRadiusScale     float64
}

func (g *Game) Update() error {
	// We are going to implement the use of normal distribution to simulate that of paint splattering.
	g.x = (rand.NormFloat64()*g.standardDeviation + g.mean) + ScreenWidth/2
	g.y = (rand.NormFloat64()*g.standardDeviation + g.mean) + ScreenHeight/2

	g.splatterRadius = rand.NormFloat64() * g.splatterRadiusScale

	mx, my := ebiten.CursorPosition()
	g.meanSlider.Update(mx, my, ebiten.IsMouseButtonPressed(ebiten.MouseButton0))
	g.standardDeviationSlider.Update(mx, my, ebiten.IsMouseButtonPressed(ebiten.MouseButton0))

	g.mean = g.meanSlider.Value
	g.standardDeviation = g.standardDeviationSlider.Value

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, nil)
	vector.DrawFilledCircle(g.bg, float32(g.x), float32(g.y), float32(g.splatterRadius), color.RGBA{0, 0, 0, 255}, true)

	g.meanSlider.Draw(screen)
	g.standardDeviationSlider.Draw(screen)

	txtop := &text.DrawOptions{}
	txtop.ColorScale.ScaleWithColor(color.Black)
	txtop.LineSpacing = 14
	text.Draw(screen, fmt.Sprintf("Mean: %.2f\nStandard Deviation: %.2f", g.mean, g.standardDeviation), ui.FontFace, txtop)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	g := &Game{
		bg:                  ebiten.NewImage(ScreenWidth, ScreenHeight),
		mean:                20,
		standardDeviation:   300,
		splatterRadiusScale: 1,
		meanSlider: &ui.Slider{
			Label: "Mean",
			X:     20,
			Y:     200,
			Width: 120,
			Min:   0,
			Max:   30,
			Value: 0,
			Color: color.Black,
		},
		standardDeviationSlider: &ui.Slider{
			Label: "Standard Deviation",
			X:     20,
			Y:     220,
			Width: 120,
			Min:   0,
			Max:   100,
			Value: 0,
			Color: color.Black,
		},
	}
	g.bg.Fill(color.White)
	return g
}

func main() {
	ebiten.SetWindowTitle("Paint Splatter")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
