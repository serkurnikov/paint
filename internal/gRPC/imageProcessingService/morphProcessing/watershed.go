package morphProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
)

const (
	fusionPath    = "fusion.jpg"
	watershedPath = "watershed.jpg"
)

func Watershed(in string, nErode, nDilate int) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}

	grayMat := gocv.NewMat()
	defer grayMat.Close()
	gocv.CvtColor(img, &grayMat, gocv.ColorBGRToGray)

	thresholdMat := gocv.NewMat()
	defer thresholdMat.Close()
	gocv.Threshold(grayMat, &thresholdMat, 0, 255, gocv.ThresholdOtsu)

	fg := gocv.NewMat()
	defer fg.Close()

	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	defer kernel.Close()

	gocv.ErodeWithParams(thresholdMat, &fg, kernel, image.Pt(-1, -1), nErode, 0)

	bgt := gocv.NewMat()
	defer bgt.Close()
	gocv.DilateWithParams(thresholdMat, &bgt, kernel, image.Pt(-1, -1), nDilate, 0)

	bg := gocv.NewMat()
	defer bg.Close()
	gocv.Threshold(bgt, &bg, 1, 128, gocv.ThresholdBinaryInv)

	marker := gocv.NewMat()
	defer marker.Close()
	gocv.Add(fg, bg, &marker)

	/*outF := basePath + fmt.Sprintf("\\e_%v_d_%v_%s",
		nErode, nDilate, fusionPath)

	if ok := gocv.IMWrite(outF, marker); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}*/

	markerCV32FC1 := gocv.NewMat()
	defer markerCV32FC1.Close()
	marker.ConvertTo(&markerCV32FC1, gocv.MatTypeCV32SC1)

	gocv.Resize(img, &img, image.Pt(markerCV32FC1.Cols(), markerCV32FC1.Rows()), 0, 0, gocv.InterpolationLinear)

	gocv.Watershed(img, &markerCV32FC1)

	m := gocv.NewMat()
	defer m.Close()
	gocv.ConvertScaleAbs(markerCV32FC1, &m, 1, 0)

	gocv.Threshold(m, &thresholdMat, 0, 255, gocv.ThresholdOtsu)

	dest := gocv.NewMat()
	defer dest.Close()

	img.CopyToWithMask(&dest, thresholdMat)

	outW := basePath + fmt.Sprintf("\\e_%v_d_%v_%s",
		nErode, nDilate, watershedPath)

	if ok := gocv.IMWrite(outW, thresholdMat); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
