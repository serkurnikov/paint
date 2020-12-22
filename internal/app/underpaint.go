package app

const (
	basePath       = "/Users/sergeykurnikov/Documents/GoProjects/paint/internal/resources/"
	testPathImage0 = basePath + "test0.jpg"
	testPathImage1 = basePath + "test1.jpg"
	testPathImage2 = basePath + "test2.jpg"
	testPathImage3 = basePath + "test3.jpg"
	testPathImage4 = basePath + "test4.jpg"
	testPathImage5 = basePath + "test5.jpg"
	testPathImage6 = basePath + "test6.jpg"
	testPathImage7 = basePath + "test7.jpg"
	testPathImage8 = basePath + "test8.jpg"

	laplacianPathImage = basePath + "laplacian.jpg"
	sobelPathImage     = basePath + "sobel.jpg"
	thresholdPathImage = basePath + "threshold.jpg"
	erodePathImage     = basePath + "erode.jpg"
	dilatePathImage    = basePath + "dilate.jpg"
	resultPathImage    = basePath + "result.jpg"
	watershedPathImage = basePath + "watershed.jpg"
	meanShiftPathImage = basePath + "meanShift.jpg"
	contoursPathImage  = basePath + "contours.jpg"
	pencilPathImage    = basePath + "pencil.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {}

func (a App) Scobel() {

	sobel(testPathImage0, sobelPathImage)
	pencil(sobelPathImage, pencilPathImage)
	meanShiftFilter(pencilPathImage, meanShiftPathImage, []float64{5, 10})
	watershed(testPathImage0, watershedPathImage)
	drawContoursCustom(testPathImage0, contoursPathImage)

	/*
		meanShiftFilter(testPathImage2, meanShiftPathImage, []float64{15, 30})
		sobel(meanShiftPathImage, sobelPathImage)
		dilate(sobelPathImage, dilatePathImage, 3)
		threshold(dilatePathImage, thresholdPathImage)
		applyMask(testPathImage2, resultPathImage, thresholdPathImage)
		watershed(testPathImage2, watershedPathImage)
	*/
}
