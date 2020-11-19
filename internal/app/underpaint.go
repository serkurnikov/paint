package app

import (
	"github.com/disintegration/gift"
	"image"
)

const (
	originalPathImage = "D:\\Sergey\\projects\\Go Projects\\tmp\\image.jpg"
	resultPathImage   = "D:\\Sergey\\projects\\Go Projects\\tmp\\result.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {
	src, _ := getImageFromFilePath(originalPathImage)

	g := gift.New(gift.Median(100, true))

	dst := image.NewNRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)

	saveImage(resultPathImage, dst)
}
