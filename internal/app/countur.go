package app

import (
	"github.com/go-opencv/go-opencv/opencv"
)

var (
	seq            *opencv.Seq
	redColor       = opencv.NewScalar(0, 0, 255, 255) // (blue, green, red, alpha)
	blackColor     = opencv.NewScalar(0, 0, 0, 255)   // (blue, green, red, alpha)
	blueColor      = opencv.NewScalar(255, 0, 0, 255) // (blue, green, red, alpha)
)

func CreateCountur(input, output string) {

	var image = opencv.LoadImage(input)
	if image == nil {
		panic("LoadImage failed")
	}
	defer image.Release()

	width := image.Width()
	height := image.Height()

	// Convert to grayscale
	gray := opencv.CreateImage(width, height, opencv.IPL_DEPTH_8U, 1)
	defer gray.Release()

	opencv.CvtColor(image, gray, opencv.CV_BGR2GRAY)

	// for edge detection
	cannyImage := opencv.CreateImage(width, height, opencv.IPL_DEPTH_8U, 1)
	defer cannyImage.Release()

	// Run the edge detector on grayscale
	//opencv.Canny(gray, cannyImage, float64(position), float64(position)*2, 3)

	// ** For better result, use 50 for the canny threshold instead of tying the value
	//    to the track bar position
	opencv.Canny(gray, cannyImage, float64(50), float64(50)*2, 3)

	// Find contours sequence from canny edge processed image
	// see http://docs.opencv.org/2.4/modules/imgproc/doc/structural_analysis_and_shape_descriptors.html
	// for mode and method

	seq = cannyImage.FindContours(opencv.CV_RETR_TREE, opencv.CV_CHAIN_APPROX_NONE, opencv.Point{0, 0})
	defer seq.Release()

	// based on the sequence, draw the contours
	// back on the original image
	finalImage := image.Clone()

	maxLevel := 1
	opencv.DrawContours(finalImage, seq, redColor, blueColor, maxLevel, 2, 8, opencv.Point{0, 0})
	opencv.SaveImage(output, finalImage, nil)
}