package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

func sobel(in, out string) {
	img := gocv.IMRead(in, gocv.IMReadGrayScale)
	if img.Empty() {

	}
	defer img.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	gocv.Sobel(img, &dest, gocv.MatTypeCV16S, 0, 1, 3, 1, 0, gocv.BorderDefault)

	if ok := gocv.IMWrite(out, dest); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
