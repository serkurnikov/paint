package app

import (
	"github.com/disintegration/gift"
	"image"
	"paint/internal/app/monography"
	"paint/internal/app/scobel"
)

const (
	basePath            = "/Users/sergeykurnikov/Documents/GoProjects/paint/internal/resources/"
	originalPathImage   = basePath + "image.jpg"
	resultPathImage     = basePath + "result.jpg"
	contourPathImage    = basePath + "contour.jpg"
	scobelPathImage     = basePath + "scobel.jpg"
	monographyPathImage = basePath + "monography.jpg"
	thresholdPathImage  = basePath + "threshold.jpg"
	invertPathImage     = basePath + "invert.jpg"
	testPathImage       = basePath + "test.jpg"
	testPathImage2      = basePath + "test2.jpg"
	testPathImage3      = basePath + "test3.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {
	src, _ := getImageFromFilePath(originalPathImage)

	g := gift.New(gift.Median(tileSize, true))

	dst := image.NewNRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)

	saveImage(resultPathImage, dst)
}

func Invert() {
	src, _ := getImageFromFilePath(scobelPathImage)

	g := gift.New(gift.Invert())

	dst := image.NewNRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)

	saveImage(invertPathImage, dst)
}
func (a App) Scobel() {
	src, _ := getImageFromFilePath(testPathImage2)
	var edge = sobel.Filter(src)
	saveImage(scobelPathImage, edge)

	monography.Monograph(scobelPathImage, monographyPathImage, []int{0, 2, 0, 0})
	monography.Threshold(monographyPathImage, thresholdPathImage)

	monography.ApplyMask(testPathImage2, resultPathImage, thresholdPathImage)

	CreateCountur(thresholdPathImage, contourPathImage)

	Invert()
}
