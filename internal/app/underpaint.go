package app

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"net/http"
	"os"
)

const (
	imageURL          = "http://i.imgur.com/m1UIjW1.jpg"
	originalPathImage = "/tmp/image2.jpg"
	resultPathImage   = "/tmp/result.jpg"
)

func (a App) ExternalApiTest() {}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func calculateAverageColor(img image.Image) [4]float64 {
	bounds := img.Bounds()
	r, g, b, a := 0.0, 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, a1 := img.At(x, y).RGBA()
			r, g, b, a = r+float64(r1), g+float64(g1), b+float64(b1), a+float64(a1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [4]float64{r / totalPixels, g / totalPixels, b / totalPixels, a / totalPixels}
}

func downloadFile(filepath, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func saveImage(filepath string, img image.Image) error {

	f, err := os.Create(filepath)
	if err != nil {

	}
	defer f.Close()

	opt := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(f, img, &opt)
	if err != nil {

	}

	return err
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}

func underTileImage(img image.Image) image.Image {
	bounds := img.Bounds()

	resultImage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X,
		bounds.Max.X, bounds.Max.Y))

	green := color.RGBA{
		uint8(calculateAverageColor(img)[0]),
		uint8(calculateAverageColor(img)[1]),
		uint8(calculateAverageColor(img)[2]),
		250}

	draw.Draw(resultImage, bounds, &image.Uniform{green}, image.Point{0, 0}, draw.Src)

	return resultImage
}

func (a App) UnderPaint(tileSize int) {

	original, _ := getImageFromFilePath(originalPathImage)
	bounds := original.Bounds()

	resultImage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X,
		bounds.Max.X, bounds.Max.Y))

	draw.Draw(resultImage, bounds, original, image.Point{0, 0}, draw.Src)
	saveImage(resultPathImage, resultImage)

	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {

			tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
			tileImage := original.(SubImager).SubImage(tileBounds)
			img := underTileImage(tileImage)

			draw.Draw(resultImage, tileBounds, img, image.Point{-x, -y}, draw.Src)
			saveImage(resultPathImage, resultImage)
		}
	}
}
