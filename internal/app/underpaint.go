package app

import (
	"paint/internal/gRPC/imageProcessingService/colorProcessing/prominentcolor"
)

//https://github.com/golang-standards/project-layout

const (
	basePath       = "D:\\Sergey\\projects\\Go Projects\\paint\\assets\\examples"
	testPathImage0 = basePath + "\\test0.jpg"
	testPathImage1 = basePath + "\\test1.jpg"

	meanShiftPathImage = basePath + "\\meanShift.jpg"
	thresholdPathImage = basePath + "\\threshold.jpg"
	watershedPathImage = basePath + "\\watershed.jpg"

	palettePathImage = basePath + "\\palette.jpg"
)

func (a App) ExternalApiTest() {}
func (a App) Scobel()          {}

func (a App) Render(ctx Ctx) {
	prominentcolor.DisplayPictureInDominatedColors(meanShiftPathImage, palettePathImage, 10)
}
