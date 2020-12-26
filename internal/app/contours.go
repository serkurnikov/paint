package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"sort"
)

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

func drawDefaultContours(in, out string) {
	red := color.RGBA{255, 0, 0, 255}

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
	contours := gocv.FindContours(matCanny, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

	gocv.DrawContours(&img, contours, -1, red, 2)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func drawLines(in, out string) {
	red := color.RGBA{255, 0, 0, 255}

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

	matLines := gocv.NewMat()
	defer matLines.Close()

	gocv.HoughLinesP(matCanny, &matLines, 0.5, math.Pi/360, 20)

	gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

	for index := 0; index < matLines.Rows(); index++ {
		pt1 := image.Pt(int(matLines.GetVeciAt(index, 0)[0]),
			int(matLines.GetVeciAt(index, 0)[1]))

		pt2 := image.Pt(int(matLines.GetVeciAt(index, 0)[2]),
			int(matLines.GetVeciAt(index, 0)[3]))

		gocv.Line(&img, pt1, pt2, red, 2)
	}

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func drawCircles(in, out string) {
	red := color.RGBA{255, 0, 0, 255}

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
	gocv.HoughCirclesWithParams(matCanny, &matCircles, gocv.HoughGradient, 2,
		float64(img.Rows()/8), 100, 100, img.Rows()/50, img.Rows()/4)

	//gocv.HoughCircles(matCanny, &matCircles, 3, 2, float64(img.Rows()/4))

	gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

	for index := 0; index < matCircles.Cols(); index++ {

		point := image.Pt(
			int(matCircles.GetVecfAt(0, index)[0]),
			int(matCircles.GetVecfAt(0, index)[1]))

		radius := int(matCircles.GetVecfAt(0, index)[2])
		gocv.Circle(&img, point, radius, red, 1)
	}

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func drawCustomContours(in, out string) {
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
	gocv.Canny(morph, &edges, 50, 100)
	//possibly instead of canny, use threshold?: threshold (gray, bw, 0, 255, THRESH_BINARY|THRESH_OTSU);

	contours := gocv.FindContours(edges, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	var toSort CustomContour
	toSort.c = contours
	//find the contour with the largest area
	sort.Sort(CustomContour(toSort))

	statusColor := color.RGBA{255, 0, 0, 0}

	if len(toSort.c) > 0 {
		gocv.FillPoly(&img, toSort.c, statusColor)
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

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func getHoles(in, out string) {
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

	statusColor := color.RGBA{255, 0, 0, 0}
	gocv.DrawContours(&binImageInv, contour, 0, statusColor, 1)

	nt := gocv.NewMat()
	defer nt.Close()

	gocv.BitwiseNot(binImage, &nt)
	gocv.BitwiseOr(binImageInv, nt, &binImageInv)

	if ok := gocv.IMWrite(out, binImageInv); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func grabCut(in string) {
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
