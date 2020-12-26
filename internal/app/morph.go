package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
	"path/filepath"
)


func morphClose(in, out string) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	morph := gocv.NewMat()
	kernel := gocv.GetStructuringElement(gocv.MorphEllipse, image.Pt(11, 11))
	defer kernel.Close()

	gocv.MorphologyEx(img, &morph, gocv.MorphClose, kernel)

	defer morph.Close()

	if ok := gocv.IMWrite(out, morph); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func morphOpen(in, out string) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	morph := gocv.NewMat()
	kernel := gocv.GetStructuringElement(gocv.MorphEllipse, image.Pt(11, 11))
	defer kernel.Close()

	gocv.MorphologyEx(img, &morph, gocv.MorphOpen, kernel)

	defer morph.Close()

	if ok := gocv.IMWrite(out, morph); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func morphGradient(in, out string) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	morph := gocv.NewMat()
	kernel := gocv.GetStructuringElement(gocv.MorphEllipse, image.Pt(11, 11))
	defer kernel.Close()

	gocv.MorphologyEx(img, &morph, gocv.MorphGradient, kernel)

	defer morph.Close()

	if ok := gocv.IMWrite(out, morph); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
