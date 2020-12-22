package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

func meanShiftFilter(in, out string, params []float64) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}

	imageSegment := gocv.NewMat()
	gocv.PyrMeanShiftFiltering(img, &imageSegment, params[0], params[1], 3)

	if ok := gocv.IMWrite(out, imageSegment); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
