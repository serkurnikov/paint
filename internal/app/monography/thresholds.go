package monography

import (
	"fmt"
	"github.com/go-opencv/go-opencv/opencv"
)

func Threshold(input, output string) {

	img := opencv.LoadImage(input, opencv.CV_LOAD_IMAGE_GRAYSCALE)
	if img == nil {
		fmt.Printf("LoadImage fail: %s\n", input)
		return
	}
	defer img.Release()

	ThresholdImage(img, output)
}

var (
	threshTypes = []string{
		"Binary",          // opencv.CV_THRESH_BINARY
		"Binary Inverse",  // opencv.CV_THRESH_BINARY_INV
		"Truncated",       // opencv.CV_THRESH_TRUNC
		"To Zero",         // opencv.CV_THRESH_TOZERO
		"To Zero Inverse", // opencv.CV_THRESH_TOZERO_INV
	}

	adaptiveMethods = []string{
		"Mean",     // opencv.CV_ADAPTIVE_THRESH_MEAN_C
		"Gaussian", // opencv.CV_ADAPTIVE_THRESH_GAUSSIAN_C
	}
)

func ThresholdImage(img *opencv.IplImage, output string) {

	threshImg := opencv.CreateImage(img.Width(), img.Height(), opencv.IPL_DEPTH_8U, 1)
	defer threshImg.Release()

	threshType := 0      //Max 4
	thresh := 128        //Max 255
	maxVal := 255        //Max 255
	fixedOrAdaptive := 0 //Max 1
	adaptMethd := 0      //Max 1
	blockSize := 3       //Max 100

	if fixedOrAdaptive == 0 {
		// fixed threshold:
		fmt.Println("*************************")
		fmt.Printf("Fixed threshold: %d\n", thresh)
		fmt.Printf("maxValue: %d\n", maxVal)
		fmt.Printf("thresType: '%s'\n", threshTypes[threshType])

		opencv.Threshold(
			img,
			threshImg,
			float64(thresh),
			float64(maxVal),
			threshType,
		)
	} else {
		// check to make sure we have a valid threshold type. (Binary or Binary inverse only)
		if threshType > 1 {
			fmt.Println("*** Only threshold types of Binary or Binary Inverse are allowed for adaptive thresholds!")
			return
		}

		// blockSize must be odd and >= 3
		if blockSize < 3 {
			blockSize = 3
		}
		if blockSize%2 == 0 {
			blockSize++
		}
		fmt.Println("*************************")
		fmt.Printf("Adaptive threshold: %d\n", thresh)
		fmt.Printf("blockSize: %d\n", blockSize)
		fmt.Printf("maxValue: %d\n", maxVal)
		fmt.Printf("threshType: '%s'\n", threshTypes[threshType])
		fmt.Printf("adaptiveMethod: '%s'\n", adaptiveMethods[adaptMethd])
		// adaptive threshold:
		opencv.AdaptiveThreshold(
			img,
			threshImg,
			float64(maxVal),
			adaptMethd,
			threshType,
			blockSize,
			float64(thresh),
		)
	}

	opencv.SaveImage(output, threshImg, nil)
}
