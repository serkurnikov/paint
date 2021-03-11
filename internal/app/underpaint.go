package app

import (
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
)

func (a App) ExternalApiTest() {}
func (a App) Scobel()          {}

func (a App) Render(ctx Ctx) {
	r, err := a.imageProcessingClient.FindBlendStructureAmongFabricColorsLUV(ctx, &pb.BlendStructureRequest{
		MainColorS:  "#85739b",
		ColorFabric: mixcolors.MasterColors,
	})

	println(r.BlendStructures[0].ResultHex)
	if err != nil {

	}
}
