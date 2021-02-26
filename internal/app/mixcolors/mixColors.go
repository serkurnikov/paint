package mixcolors

import (
	"github.com/Jeffail/gabs/v2"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	HEIGHT = 1000
	WEIGHT = 1000
)

type Blend interface {
	BlendColors(colorS1, colorS2 string, numberOfShades int) *gabs.Container
	BlendCombination(combinationElements []string)
}

type Palette interface {
	CreatePaletteAllColors(data []*gabs.Container)
	CreatePaletteByDominateColors(dominateColors []string)
	SearchColorInPalette(colorS string) (colorResultS string)
	SearchSimilarColorsInPalette(colorS string) (colorResultS []string)
}

type MixColors interface {
	Blend
	Palette
	DisplayMixColors(colors []color.Color, out string)
	GetAllMixColors(colorsDataS []string, numberOfShades int) *gabs.Container
}


func DisplayMixColors(colors []color.Color, out string) {
	result := image.NewRGBA(image.Rect(0, 0, HEIGHT, WEIGHT))

	h := HEIGHT / len(colors)
	w := WEIGHT / len(colors)

	for i := 0; i < len(colors); i++ {
		for j := 0; j < len(colors); j++ {
			draw.Draw(result, image.Rect(0, i*h, j*w+w, i*h+h), &image.Uniform{C: colors[i]}, image.Point{
				X: 0,
				Y: 0,
			}, draw.Src)
		}
	}

	file, err := os.Create(out)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, result)
}
