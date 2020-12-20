package monography

import (
	"fmt"
	"github.com/go-opencv/go-opencv/opencv"
)

func ApplyMask(input, output, mask string) {
	img := opencv.LoadImage(input, opencv.CV_LOAD_IMAGE_COLOR)
	if img == nil {
		fmt.Printf("LoadImage fail: %s\n", input)
		return
	}
	defer img.Release()

	m := opencv.LoadImage(mask, opencv.CV_LOAD_IMAGE_GRAYSCALE)
	if m == nil {
		fmt.Printf("LoadImage fail: %s\n", input)
		return
	}
	resizedMask := opencv.Resize(m, img.Width(), img.Height(), 0)
	defer m.Release()

	result := opencv.CreateImage(img.Width(), img.Height(), opencv.IPL_DEPTH_8U, 3)
	defer result.Release()

	opencv.Copy(img, result, resizedMask)
	opencv.SaveImage(output, result, nil)
}
