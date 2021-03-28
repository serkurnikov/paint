package contoursProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"sort"
)

const (
	basePath             = "D:\\Sergey\\projects\\Go Projects\\paint\\assets\\examples"
	defaultContour       = "defaultContour.jpg"
	customContours       = "customContours.jpg"
	houghLinesWithParams = "houghLinesWithParams.jpg"
	houghCircles         = "houghCircles.jpg"
	holes                = "holes.jpg"
)

var colorCountours = color.RGBA{R: 255, G: 0, B: 0, A: 255}

type CustomContour struct {
	c     [][]image.Point
	index int
	area  float64
}

func (cc CustomContour) Len() int {
	return len(cc.c)
}

func (cc CustomContour) Less(i, j int) bool {
	aI := gocv.ContourArea(cc.c[i])
	aJ := gocv.ContourArea(cc.c[j])
	if aI > aJ {
		return true
	}
	return false
}

func (cc CustomContour) Swap(i, j int) {
	cc.c[i], cc.c[j] = cc.c[j], cc.c[i]
}

func DrawDefaultContours(in string, t1 float32, t2 float32) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	matCanny := gocv.NewMat()
	defer matCanny.Close()

	gocv.Canny(img, &matCanny, t1, t2)
	contours := gocv.FindContours(matCanny, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

	gocv.DrawContours(&img, contours, -1, colorCountours, 2)

	out := basePath + fmt.Sprintf("\\t1_%f_t2_%f_%s", t1, t2, defaultContour)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func DrawHoughLinesWithParams(in string, rho float32, theta float32, threshold int, minLineLength float32, maxLineGap float32) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	matCanny := gocv.NewMat()
	defer matCanny.Close()
	gocv.Canny(img, &matCanny, 100, 200)

	matLines := gocv.NewMat()
	defer matLines.Close()

	gocv.HoughLinesPWithParams(matCanny, &matLines, rho, theta, threshold, minLineLength, maxLineGap)

	gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

	for index := 0; index < matLines.Rows(); index++ {
		pt1 := image.Pt(int(matLines.GetVeciAt(index, 0)[0]),
			int(matLines.GetVeciAt(index, 0)[1]))

		pt2 := image.Pt(int(matLines.GetVeciAt(index, 0)[2]),
			int(matLines.GetVeciAt(index, 0)[3]))

		gocv.Line(&img, pt1, pt2, colorCountours, 2)
	}

	out := basePath + fmt.Sprintf("\\rho_%f_theta_%f_threshold_%v_minL_%f_maxG_%f_%s",
		rho, theta, threshold, minLineLength, maxLineGap, houghLinesWithParams)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func DrawHoughCircles(in string, method gocv.HoughMode, dp, minDist, param1, param2 float64, minRadius, maxRadius int) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	matCanny := gocv.NewMat()
	defer matCanny.Close()
	gocv.Canny(img, &matCanny, 50, 100)

	matCircles := gocv.NewMat()
	defer matCircles.Close()
	gocv.HoughCirclesWithParams(matCanny, &matCircles, method, dp,
		minDist, param1, param2, minRadius, maxRadius)

	gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

	for index := 0; index < matCircles.Cols(); index++ {

		point := image.Pt(
			int(matCircles.GetVecfAt(0, index)[0]),
			int(matCircles.GetVecfAt(0, index)[1]))

		radius := int(matCircles.GetVecfAt(0, index)[2])
		gocv.Circle(&img, point, radius, colorCountours, 1)
	}

	out := basePath + fmt.Sprintf("\\dp_%f_minDist_%f_p1_%f_p2_%f_minR_%v_minR_%v_%s",
		dp, minDist, param1, param2, minRadius, maxRadius, houghCircles)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func DrawCustomContours(in string, t1 float32, t2 float32) [][]image.Point {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	blur := gocv.NewMat()
	defer blur.Close()
	gocv.GaussianBlur(img, &blur, image.Pt(3, 3), 1, 1, gocv.BorderDefault)

	//erode edges so that 'more is inside of the edge detection
	eroded := gocv.NewMat()
	{
		kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
		defer kernel.Close()
		gocv.Erode(blur, &eroded, kernel)
	}
	defer eroded.Close()

	//median blur to remove 'salt and pepper'
	medianBlur := gocv.NewMat()
	defer medianBlur.Close()

	gocv.MedianBlur(eroded, &medianBlur, 9)

	//use morphology to try and join lines up to create a continuous line
	morph := gocv.NewMat()
	{
		kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
		defer kernel.Close()

		gocv.MorphologyEx(medianBlur, &morph, gocv.MorphClose, kernel)
	}
	defer morph.Close()

	edges := gocv.NewMat()
	defer edges.Close()
	/*
		FYI
			// The smallest value between threshold1 and threshold2 is used
			// for edge linking. The largest value is used to
			// find initial segments of strong edges.
	*/

	//finally detect edges
	gocv.Canny(morph, &edges, t1, t2)
	//possibly instead of canny, use threshold?: threshold (gray, bw, 0, 255, THRESH_BINARY|THRESH_OTSU);

	contours := gocv.FindContours(edges, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	var toSort CustomContour
	toSort.c = contours
	//find the contour with the largest area
	sort.Sort(toSort)

	statusColor := color.RGBA{R: 255, G: 255, B: 255}

	if len(toSort.c) > 0 {
		gocv.DrawContours(&img, toSort.c, -1, statusColor, 1)
	}

	hull := gocv.NewMat()
	defer hull.Close()

	if len(toSort.c) > 0 {
		gocv.ConvexHull(toSort.c[0], &hull, false, false)
		var hullPoints []image.Point
		for i := 0; i < hull.Cols(); i++ {
			for j := 0; j < hull.Rows(); j++ {
				p := hull.GetIntAt(j, i)
				hullPoints = append(hullPoints, toSort.c[0][p])
			}
		}

		rect := gocv.MinAreaRect(hullPoints)

		pts := gocv.NewMat()
		defer pts.Close()

		gocv.BoxPoints(rect, &pts)
	}

	out := basePath + fmt.Sprintf("\\t1_%f_t2_%f_%s", t1, t2, customContours)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}

	return toSort.c
}

func GetHoles(in string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)

	binImage := gocv.NewMat()
	defer binImage.Close()
	gocv.Threshold(img, &binImage, 128, 255, gocv.ThresholdBinaryInv)

	binImageInv := gocv.NewMat()
	defer binImageInv.Close()
	gocv.BitwiseNot(binImage, &binImageInv)

	contour := gocv.FindContours(binImageInv, gocv.RetrievalCComp, gocv.ChainApproxSimple)

	statusColor := color.RGBA{R: 255}
	gocv.DrawContours(&binImageInv, contour, 0, statusColor, 1)

	nt := gocv.NewMat()
	defer nt.Close()

	gocv.BitwiseNot(binImage, &nt)
	gocv.BitwiseOr(binImageInv, nt, &binImageInv)

	out := basePath + fmt.Sprintf("\\%s", holes)

	if ok := gocv.IMWrite(out, binImageInv); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func GrabCut(in string) {
	img := gocv.IMRead(in, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to read image\n")
	}
	defer img.Close()

	src := gocv.NewMat()
	defer src.Close()
	gocv.CvtColor(img, &img, gocv.ColorRGBAToBGR)
	img.ConvertTo(&src, gocv.MatTypeCV8UC3)

	mask := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)
	defer mask.Close()

	bgdModel := gocv.NewMat()
	defer bgdModel.Close()
	fgdModel := gocv.NewMat()
	defer fgdModel.Close()

	r := image.Rect(10, 10, 500, 500)

	gocv.GrabCut(src, &mask, r, &bgdModel, &fgdModel, 1, gocv.GCEval)
}
