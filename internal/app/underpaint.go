package app

const (
	//change your directory for saving results

	basePath        = "D:\\Sergey\\projects\\Go Projects\\paint\\internal\\resources"
	testPathImage0  = basePath + "\\test0.jpg"
	testPathImage1  = basePath + "\\test1.jpg"
	testPathImage2  = basePath + "\\test2.jpg"
	testPathImage3  = basePath + "\\test3.jpg"
	testPathImage4  = basePath + "\\test4.jpg"
	testPathImage5  = basePath + "\\test5.jpg"
	testPathImage6  = basePath + "\\test6.jpg"
	testPathImage7  = basePath + "\\test7.jpg"
	testPathImage8  = basePath + "\\test8.jpg"
	testPathImage9  = basePath + "\\test9.jpg"
	testPathImage10 = basePath + "\\test10.jpg"

	laplacianPathImage = basePath + "\\laplacian.jpg"
	sobelPathImage     = basePath + "\\sobel.jpg"
	thresholdPathImage = basePath + "\\threshold.jpg"
	erodePathImage     = basePath + "\\erode.jpg"
	dilatePathImage    = basePath + "\\dilate.jpg"
	resultPathImage    = basePath + "\\result.jpg"
	watershedPathImage = basePath + "\\watershed.jpg"
	fusionPathImage    = basePath + "\\fusion.jpg"
	meanShiftPathImage = basePath + "\\meanShift.jpg"
	contoursPathImage  = basePath + "\\contours.jpg"
	contoursPathImageR = basePath + "\\contoursR.jpg"
	contoursPathImageL = basePath + "\\contoursL.jpg"
	contoursPathImageC = basePath + "\\contoursC.jpg"
	pencilPathImage    = basePath + "\\pencil.jpg"
	floodFillPathImage = basePath + "\\floodFill.jpg"
	palettePathImage = basePath + "\\palette.jpg"

	morphPathClose    = basePath + "\\morphClose.jpg"
	morphPathOpen     = basePath + "\\morphOpen.jpg"
	morphPathGradient = basePath + "\\morphGradient.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {
	meanShiftFilter(testPathImage0, meanShiftPathImage, []float64{30, 60})
	drawCustomContours(meanShiftPathImage, contoursPathImageC)

	sobel(testPathImage0, sobelPathImage)
	drawCustomContours(meanShiftPathImage, contoursPathImageC)

	dilate(sobelPathImage, dilatePathImage, 3)
	erode(sobelPathImage, erodePathImage, 1)

	watershed(testPathImage0, fusionPathImage, watershedPathImage, 1, 1)
	pencil(sobelPathImage, pencilPathImage)

	threshold(dilatePathImage, thresholdPathImage)
	applyMask(testPathImage0, resultPathImage, watershedPathImage)
}

func (a App) Scobel() {}
