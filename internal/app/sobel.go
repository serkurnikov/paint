package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

func sobel(in, out string) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {

	}
	defer img.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	//gocv.GaussianBlur(img, &img, image.Pt(3, 3), 0, 0, 4)
	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	gradX := gocv.NewMat()
	absGradX := gocv.NewMat()
	defer gradX.Close()
	defer absGradX.Close()

	gradY := gocv.NewMat()
	absGradY := gocv.NewMat()
	defer gradY.Close()
	defer absGradY.Close()

	gocv.Sobel(img, &gradX, gocv.MatTypeCV16S, 1, 0, 3, 1, 0, gocv.BorderDefault)
	gocv.Sobel(img, &gradY, gocv.MatTypeCV16S, 0, 1, 3, 1, 0, gocv.BorderDefault)

	gocv.ConvertScaleAbs(gradX, &absGradX, 1, 0)
	gocv.ConvertScaleAbs(gradY, &absGradY, 1, 0)

	gocv.AddWeighted(absGradX, 0.5, absGradY, 0.5, 0, &dest)

	if ok := gocv.IMWrite(out, dest); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
