package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
	"os"
	"path/filepath"
)

func drawContours(in, out string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
	gocv.Threshold(img, &img, 128, 255, gocv.ThresholdBinaryInv)

	contours := gocv.FindContours(img, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	gocv.DrawContours(&img, contours, -1, color.RGBA{0,255,75,1}, 1)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func drawContoursCustom(in, out string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
	gocv.Canny(img, &img, 50, 100)

	contours := gocv.FindContours(img, gocv.RetrievalTree, gocv.ChainApproxNone)
	gocv.DrawContours(&img, contours, -1, color.RGBA{0, 255, 75, 1}, 1)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}