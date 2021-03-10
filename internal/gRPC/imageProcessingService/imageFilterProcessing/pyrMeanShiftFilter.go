package imageFilterProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

const (
	basePath           = "D:\\Sergey\\projects\\Go Projects\\paint\\assets\\examples"
	meanShiftPathImage = basePath + "\\meanShift.jpg"
)

func PyrMeanShiftFiltering(in string, sp float32, sr float32, maxLevel int32) string {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}

	imageSegment := gocv.NewMat()
	gocv.PyrMeanShiftFiltering(img, &imageSegment, float64(sp), float64(sr), int(maxLevel))

	if ok := gocv.IMWrite(meanShiftPathImage, imageSegment); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
	return meanShiftPathImage
}
