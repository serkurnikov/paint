package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

func threshold(in, out string) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}
	// turn the img to gray image
	grayImage := gocv.NewMat()
	gocv.CvtColor(img, &grayImage, gocv.ColorBGRToGray)
	// binary threshold
	binImage := gocv.NewMat()
	gocv.Threshold(img, &binImage, 128, 255, gocv.ThresholdBinary)
	// write img to filesystem
	if ok := gocv.IMWrite(out, binImage); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

