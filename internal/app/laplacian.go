package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

func laplacian(in, out string) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
	}
	defer img.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	gocv.Laplacian(img, &dest, gocv.MatTypeCV16S, 1, 1, 0, gocv.BorderDefault)

	if ok := gocv.IMWrite(out, dest); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
