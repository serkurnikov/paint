package app

const (
	basePath        = "D:\\Sergey\\projects\\Go Projects\\paint\\examples\\assets"
	testPathImage0  = basePath + "\\test0.jpg"
	testPathImage1  = basePath + "\\test1.jpg"

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
	palettePathImage   = basePath + "\\palette.jpg"

	morphPathClose    = basePath + "\\morphClose.jpg"
	morphPathOpen     = basePath + "\\morphOpen.jpg"
	morphPathGradient = basePath + "\\morphGradient.jpg"
)

func (a App) ExternalApiTest() {}

func (a App) UnderPaint(tileSize int) {

}

func (a App) Scobel() {}
