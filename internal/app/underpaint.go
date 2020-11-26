package app

import (
	"github.com/disintegration/gift"
	"image"
	sobel "paint/internal/app/scobel"
)

const (
	//originalPathImage = "D:\\Sergey\\projects\\Go Projects\\tmp\\image.jpg"
	//resultPathImage   = "D:\\Sergey\\projects\\Go Projects\\tmp\\result.jpg"

	originalPathImage = "/Users/sergeykurnikov/Documents/GoProjects/paint/internal/resources/image.jpg"
	resultPathImage   = "/Users/sergeykurnikov/Documents/GoProjects/paint/internal/resources/result.jpg"
	scobelPathImage   = "/Users/sergeykurnikov/Documents/GoProjects/paint/internal/resources/scobel.jpg"

)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {
	src, _ := getImageFromFilePath(originalPathImage)

	g := gift.New(gift.Median(tileSize, true))

	dst := image.NewNRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)

	saveImage(resultPathImage, dst)
}

func (a App) Scobel() {
	src, _ := getImageFromFilePath(originalPathImage)
	var edge = sobel.Filter(src)

	saveImage(scobelPathImage, edge)
}
