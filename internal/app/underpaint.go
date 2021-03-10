package app

import (
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

	for i := 0; i < 10; i++ {
		_, _ = a.imageProcessingClient.Watershed(ctx, &pb.WatershedRequest{
			PathPicture: watershedPathImage,
			NErode:      int32(i)*2,
			NDilate:     int32(i),
		})
	}
}
