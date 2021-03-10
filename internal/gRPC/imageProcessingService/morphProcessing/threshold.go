package morphProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

const (
	basePath      = "D:\\Sergey\\projects\\Go Projects\\paint\\assets\\examples"
	thresholdPath = "threshold.jpg"
)

func Threshold(in string, thresh float32, maxvalue float32) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}

	grayImage := gocv.NewMat()
	gocv.CvtColor(img, &grayImage, gocv.ColorBGRToGray)

	binImage := gocv.NewMat()
	gocv.Threshold(grayImage, &binImage, thresh, maxvalue, gocv.ThresholdBinaryInv)

	out := basePath + fmt.Sprintf("\\thresh_%f_maxvalue_%f_%s",
		thresh, maxvalue, thresholdPath)

	if ok := gocv.IMWrite(out, binImage); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
