package app

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"net/http"
	"os"
)

const (
	imageURL = "http://i.imgur.com/m1UIjW1.jpg"
	pathImage = "/tmp/image.jpg"
)

func (a App) ExternalApiTest() {}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

func DownloadFile(filepath string, url string) error {

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

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}

func underPaintImage(img image.Image) image.Image {
	bounds := img.Bounds()
	aColor := averageColor(img)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			img.(draw.Image).Set(x, y, color.RGBA{uint8(aColor[0]),
				uint8(aColor[1]),
				uint8(aColor[2]), 1})
		}
	}

	return img
}

func (a App) UnderPaint(tileSize int) map[string]string {

	DownloadFile(pathImage, imageURL)

	original, _ := getImageFromFilePath(pathImage)
	bounds := original.Bounds()

	newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X,
		bounds.Max.X, bounds.Max.Y))

	sp := image.Point{0, 0}

	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {

			tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
			tileImage := original.(SubImager).SubImage(tileBounds)

			draw.Draw(newimage, tileBounds, underPaintImage(tileImage), sp, draw.Src)
		}
	}

	return output(original, newimage)
}

func output(imgOrigin, imgUnderPaint image.Image) map[string]string {
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, imgOrigin, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, imgUnderPaint, nil)
	underPaint := base64.StdEncoding.EncodeToString(buf2.Bytes())

	result := map[string]string{
		"original":   originalStr,
		"underPaint": underPaint,
	}

	return result
}
