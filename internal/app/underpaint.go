package app

//https://github.com/golang-standards/project-layout

const (
	basePath       = "D:\\Sergey\\projects\\Go Projects\\paint\\assets\\examples"
	testPathImage0 = basePath + "\\test0.jpg"
	testPathImage1 = basePath + "\\test1.jpg"

	laplacianPathImage = basePath + "\\laplacian.jpg"
	sobelPathImage     = basePath + "\\sobel.jpg"
	thresholdPathImage = basePath + "\\threshold.jpg"
	erodePathImage     = basePath + "\\erode.jpg"
	dilatePathImage    = basePath + "\\dilate.jpg"
	resultPathImage    = basePath + "\\result.jpg"
	watershedPathImage = basePath + "\\watershed.jpg"
	fusionPathImage    = basePath + "\\fusion.jpg"
	meanShiftPathImage = basePath + "\\meanShift.jpg"
	pencilPathImage    = basePath + "\\pencil.jpg"
	floodFillPathImage = basePath + "\\floodFill.jpg"
	palettePathImage   = basePath + "\\palette.jpg"

	morphPathClose    = basePath + "\\morphClose.jpg"
	morphPathOpen     = basePath + "\\morphOpen.jpg"
	morphPathGradient = basePath + "\\morphGradient.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) Render(ctx Ctx) {
	/*r, err := a.imageProcessingClient.PyrMeanShiftFiltering(ctx, &pb.PyrRequest{
		PathPicture: testPathImage0,
		Sp:          15,
		Sr:          30,
		MaxLevel:    3,
	})
	if err != nil {
		log.Fatalf("could not PyrMeanShiftFiltering: %v", err)
	}
	log.Printf("PyrMeanShiftFiltering: %s", r.OutPicture)*/

	FindingMatchingGeometricShapes(meanShiftPathImage)
}

func (a App) Scobel() {

}
