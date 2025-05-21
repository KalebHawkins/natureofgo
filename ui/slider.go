package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Slider struct {
	X, Y     float64
	Width    float64
	Min, Max float64
	Value    float64
	Dragging bool
	color.Color
}

func (s *Slider) posFromValue() float64 {
	return s.X + ((s.Value-s.Min)/(s.Max-s.Min))*s.Width
}

func (s *Slider) valueFromPos(px float64) float64 {
	clamped := math.Max(s.X, math.Min(s.X+s.Width, px))
	return s.Min + ((clamped-s.X)/s.Width)*(s.Max-s.Min)
}

func (s *Slider) Update(mx, my int, mousePressed bool) {
	if mousePressed {
		if my > int(s.Y)-10 && my < int(s.Y)+10 {
			s.Dragging = true
		}

		if s.Dragging {
			s.Value = s.valueFromPos(float64(mx))
		}
	} else {
		s.Dragging = false
	}
}

func (s *Slider) Draw(screen *ebiten.Image) {
	vector.StrokeLine(screen, float32(s.X), float32(s.Y), float32(s.X+s.Width), float32(s.Y), 2, s.Color, false)
	handleX := s.posFromValue()
	vector.DrawFilledCircle(screen, float32(handleX), float32(s.Y), 6, s.Color, true)
}
