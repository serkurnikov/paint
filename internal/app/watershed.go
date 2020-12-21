package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
)

func watershed(in, out string) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}

	grayMat := gocv.NewMat()
	gocv.CvtColor(img, &grayMat, gocv.ColorBGRToGray)

	thresholdMat := gocv.NewMat()
	gocv.Threshold(grayMat, &thresholdMat, 0, 255, gocv.ThresholdOtsu)

	fg := gocv.NewMat()
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	defer kernel.Close()
	gocv.ErodeWithParams(thresholdMat, &fg, kernel, image.Pt(-1, -1), 2, 0)

	nDilate := 3
	bgt := gocv.NewMat()
	for i := 0; i < nDilate; i++ {
		gocv.Dilate(thresholdMat, &bgt, kernel)
	}

	bg := gocv.NewMat()
	gocv.Threshold(bgt, &bg, 1, 128, gocv.ThresholdBinaryInv)

	fusion := gocv.NewMat()
	gocv.Add(fg, bg, &fusion)

	imageSegment := gocv.NewMat()
	gocv.PyrMeanShiftFiltering(img, &imageSegment, 70, 60, 3)

	if ok := gocv.IMWrite(out, imageSegment); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
