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

type Walker struct {
	x, y float64
}

func (w *Walker) Update(mean, stdDev float64) {
	xStep := rand.NormFloat64()*stdDev + mean
	yStep := rand.NormFloat64()*stdDev + mean

	w.x += xStep
	w.y += yStep
}

func (w *Walker) Draw(dst *ebiten.Image) {
	vector.DrawFilledCircle(dst, float32(w.x), float32(w.y), 2, color.Black, true)
}

type Game struct {
	bgImg                   *ebiten.Image
	walker                  *Walker
	mean                    float64
	standardDeviation       float64
	meanSlider              *ui.Slider
	standardDeviationSlider *ui.Slider
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	g.meanSlider.Update(mx, my, ebiten.IsMouseButtonPressed(ebiten.MouseButton0))
	g.standardDeviationSlider.Update(mx, my, ebiten.IsMouseButtonPressed(ebiten.MouseButton0))

	g.mean = g.meanSlider.Value
	g.standardDeviation = g.standardDeviationSlider.Value

	g.walker.Update(g.mean, g.standardDeviation)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bgImg, nil)
	g.walker.Draw(g.bgImg)
	g.meanSlider.Draw(screen)
	g.standardDeviationSlider.Draw(screen)

	txtop := &text.DrawOptions{}
	txtop.LineSpacing = ui.FontSize
	txtop.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, fmt.Sprintf("Mean: %.2f\nStandard Deviation: %.2f", g.mean, g.standardDeviation), ui.FontFace, txtop)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}

func NewGame() *Game {
	g := &Game{
		bgImg: ebiten.NewImage(ScreenWidth, ScreenHeight),
		walker: &Walker{
			x: ScreenWidth / 2,
			y: ScreenHeight / 2,
		},
		mean:              0,
		standardDeviation: 3,
		meanSlider: &ui.Slider{
			Label:    "Mean",
			X:        20,
			Y:        200,
			Width:    120,
			Min:      0,
			Max:      60,
			Value:    0,
			Dragging: false,
			Color:    color.Black,
		},
		standardDeviationSlider: &ui.Slider{
			Label: "Standard Deviation",
			X:     20,
			Y:     220,
			Width: 120,
			Min:   0,
			Max:   100,
			Value: 3,
			Color: color.Black,
		},
	}

	g.bgImg.Fill(color.White)
	return g
}

func main() {
	ebiten.SetWindowTitle("Gaussian Walker")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
