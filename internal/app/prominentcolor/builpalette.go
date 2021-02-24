package prominentcolor

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func loadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	defer f.Close()
	if err != nil {
		log.Println("File not found:", fileInput)
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Process images in a directory, for each image it picks out the dominant color and
// prints out an imagemagick call to resize image and use the dominant color as padding for the background
// it saves tmp files in /tmp/ with the masked bit marked as pink
func BuildP(in, out string, clusters int) {
	img, err := loadImage(in)
	if nil != err {
		log.Println(err)
		log.Printf("Error: failed loading %s\n", in)
		return
	}
	cols, err := KmeansWithAll(5, img, ArgumentDebugImage, DefaultSize, GetDefaultMasks())
	println(len(cols))
	if err != nil {
		log.Println(err)
		return
	}

	displayColors(img, cols, out)
}

func displayColors(img image.Image, colors []ColorItem, out string) {

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
