package prominentcolor

import (
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"paint/pkg/imageProcessingService/colorProcessing/mixcolors"
	"paint/pkg/utils"
	"sort"
)

// Process images in a directory, for each image it picks out the dominant color and
// prints out an imagemagick call to resize image and use the dominant color as padding for the background
// it saves tmp files in /tmp/ with the masked bit marked as pink
func BuildP(in string, clusters int) (image.Image, []ColorItem) {
	img, err := utils.LoadImage(in)
	if nil != err {
		log.Println(err)
		log.Printf("Error: failed loading %s\n", in)
		return nil, nil
	}
	cols, err := KmeansWithAll(clusters, img, ArgumentDefault, DefaultSize, GetDefaultMasks())
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	return img, cols
}

func DisplayColors(img image.Image, colors []ColorItem, out string) {

	result := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	h := img.Bounds().Max.X / len(colors)
	w := img.Bounds().Max.Y

	for i := 0; i < len(colors); i++ {
		currentColor := color.RGBA{R: uint8(colors[i].Color.R), G: uint8(colors[i].Color.G), B: uint8(colors[i].Color.B), A: 255}
		draw.Draw(result, image.Rect(0, i*h, w, i*h+h), &image.Uniform{C: currentColor}, image.Point{
			X: 0,
			Y: 0,
		}, draw.Src)
	}

	file, err := os.Create(out)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, result)
}

func DisplayPalette(in, out string, clusters int) {
	img, cols := BuildP(in, clusters)
	DisplayColors(img, cols, out)
}

func DisplayPictureInDominatedColors(in, out string, clusters int) {
	img, cols := BuildP(in, clusters)
	result := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			resultColor := img.At(x, y)
			additiveColor, _ := FindAdditiveColorFromDominates(cols, resultColor)
			result.Set(x, y, additiveColor)
		}
	}

	file, err := os.Create(out)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, result)
}

func FindAdditiveColorFromDominates(cols []ColorItem, originalColor color.Color) (color.Color, error) {
	data := make([]mixcolors.ColorAdditive, 0)

	for i := 0; i < len(cols); i++ {
		origCol, _ := colorful.MakeColor(originalColor)
		clusterCol, _ := colorful.Hex("#" + cols[i].AsString())

		data = append(data, mixcolors.ColorAdditive{
			Hex:  clusterCol.Hex(),
			Diff: []float64{origCol.DistanceLab(clusterCol), 0.0, 0.0},
		})
	}

	sort.Slice(data, func(i, j int) bool { return data[i].Diff[0] < data[j].Diff[0] })
	return colorful.Hex(data[0].Hex)
}

func DisplayPictureWithPalette(in, out string, clusters int) {}
