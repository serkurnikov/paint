package app

import (
	pb "paint/internal/gRPC/imageProcessingService/service"
)

//https://github.com/golang-standards/project-layout

const (
	BasePath       = "C:\\Users\\master\\go\\src\\projects\\paint\\assets\\examples\\"
	TestPathImage0 = BasePath + "test0.jpg"
	TestPathImage1 = BasePath + "test1.jpg"

	MeanShiftPathImage = BasePath + "meanShift.jpg"
	ThresholdPathImage = BasePath + "threshold.jpg"
	WatershedPathImage = BasePath + "watershed.jpg"

	PalettePathImage = BasePath + "palette.jpg"
)

func (a App) ExternalApiTest() {}
func (a App) Scobel()          {}

func (a App) Render(ctx Ctx) {
	_, _ = a.imageProcessingClient.DisplayPictureInDominatedColors(ctx, &pb.PictureInDominatedColorsRequest{
		InPicture:        MeanShiftPathImage,
		OutPicture:       PalettePathImage,
		NumberOfClusters: 10})
}
