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

	laplacianPathImage = basePath + "laplacian.jpg"
	sobelPathImage     = basePath + "sobel.jpg"
	thresholdPathImage = basePath + "threshold.jpg"
	erodePathImage     = basePath + "erode.jpg"
	dilatePathImage    = basePath + "dilate.jpg"
	resultPathImage    = basePath + "result.jpg"
	fusionPathImage    = basePath + "fusion.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {}

func (a App) Scobel() {

	sobel(testPathImage4, sobelPathImage)
	dilate(sobelPathImage, dilatePathImage)
	threshold(dilatePathImage, thresholdPathImage)
	applyMask(testPathImage4, resultPathImage, thresholdPathImage)

	watershed(testPathImage4, fusionPathImage)
}
