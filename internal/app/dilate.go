package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
	"path/filepath"
)

func dilate(in, out string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	defer kernel.Close()

	gocv.Dilate(img, &dest, kernel)

	if ok := gocv.IMWrite(out, dest); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
