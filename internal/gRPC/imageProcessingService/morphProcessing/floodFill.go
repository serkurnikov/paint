package morphProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"path/filepath"
)

func floodFill(in, out string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	// turn the img to gray image
	grayImage := gocv.NewMat()
	gocv.CvtColor(img, &grayImage, gocv.ColorBGRToGray)
	// binary threshold
	binImage := gocv.NewMat()
	gocv.Threshold(img, &binImage, 220, 255, gocv.ThresholdBinary)

	//mask := gocv.NewMatWithSize(thresed.Rows()+2, thresed.Cols()+2, gocv.MatTypeCV8U)
	//defer mask.Close()

	/*gocv.FloodFill(thresed, &mask, image.Point{},
	color.RGBA{R: 255, G: 255, B: 255},
	color.RGBA{},
	color.RGBA{}, 4)*/

	// write img to filesystem
	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
