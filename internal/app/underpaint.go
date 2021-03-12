package app

import (
	"fmt"
	"paint/internal/gRPC/imageProcessingService/colorProcessing/mixcolors"
	pb "paint/internal/gRPC/imageProcessingService/service"
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
	r, _ := a.imageProcessingClient.FindBlendStructureAmongFabricColorsLUV(ctx,
		&pb.BlendStructureRequest{MainColorS: "#ADD8E6", ColorFabric: mixcolors.MasterColors})

	fmt.Printf("\nresult Hex %v, Combination %v + %v + %v Portions %v + %v",
		r.BlendStructures[0].ResultHex,
		r.BlendStructures[0].C1Hex,
		r.BlendStructures[0].C2Hex,
		r.BlendStructures[0].C3Hex,

		r.BlendStructures[0].C2Portion,
		r.BlendStructures[0].C3Portion)
}
