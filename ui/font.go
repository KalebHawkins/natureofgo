package ui

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	FontFaceSource *text.GoTextFaceSource
	FontFace       *text.GoTextFace
	FontSize       float64     = 14
	FontColor      color.Color = color.White
)

func init() {
	FontFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}

	FontFace = &text.GoTextFace{
		Source: FontFaceSource,
		Size:   float64(FontSize),
	}
}

func SetFontSize(size int) {
	FontFace = &text.GoTextFace{
		Source: FontFaceSource,
		Size:   FontSize,
	}
}
