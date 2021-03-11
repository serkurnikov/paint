package app

import (
	"fmt"
	"paint/internal/gRPC/imageProcessingService/colorProcessing/mixcolors"
)

//https://github.com/golang-standards/project-layout

const (
	basePath       = "D:\\Sergey\\projects\\Go Projects\\paint\\assets\\examples"
	testPathImage0 = basePath + "\\test0.jpg"
	testPathImage1 = basePath + "\\test1.jpg"

	meanShiftPathImage = basePath + "\\meanShift.jpg"
	thresholdPathImage = basePath + "\\threshold.jpg"
	watershedPathImage = basePath + "\\watershed.jpg"
)

func (a App) ExternalApiTest() {}
func (a App) Scobel()          {}

func (a App) Render(ctx Ctx) {
	fmt.Printf("%v", mixcolors.FindBlendStructureAmongFabricColorsLUV("#123f31", mixcolors.MasterColors))
}
