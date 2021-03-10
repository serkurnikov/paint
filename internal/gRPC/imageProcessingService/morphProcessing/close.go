package morphProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
	"path/filepath"
)

const (
	closePath = "close.jpg"
)

func Close(in string, kernelSize int) {

	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	morph := gocv.NewMat()
	kernel := gocv.GetStructuringElement(gocv.MorphEllipse, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()

	gocv.MorphologyEx(img, &morph, gocv.MorphClose, kernel)

	defer morph.Close()

	out := basePath + fmt.Sprintf("\\kernelSize_%v_%s",
		kernelSize, closePath)

	if ok := gocv.IMWrite(out, morph); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
