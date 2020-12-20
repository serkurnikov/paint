package monography

import (
	"fmt"
	"github.com/go-opencv/go-opencv/opencv"
)

func Monograph(input, output string, params []int) {

	img := opencv.LoadImage(input, opencv.CV_LOAD_IMAGE_GRAYSCALE)
	if img == nil {
		fmt.Printf("LoadImage fail: %s\n", input)
		return
	}
	defer img.Release()
	ProcessImage(img, output, params)
}

type NamedConsts struct {
	Name string
	C    int
}

var (
	morphTypes = []NamedConsts{
		NamedConsts{"Open", opencv.CV_MORPH_OPEN},
		NamedConsts{"Close", opencv.CV_MORPH_CLOSE},
		NamedConsts{"Gradient", opencv.CV_MORPH_GRADIENT},
		NamedConsts{"Top Hat", opencv.CV_MORPH_TOPHAT},
		NamedConsts{"Black Hat", opencv.CV_MORPH_BLACKHAT},
	}
	structShapes = []NamedConsts{
		NamedConsts{"Rectangle", opencv.CV_MORPH_RECT},
		NamedConsts{"Ellipse", opencv.CV_MORPH_ELLIPSE},
		NamedConsts{"Cross", opencv.CV_MORPH_CROSS},
	}
)

func ProcessImage(img *opencv.IplImage, output string, params []int) {

	nErode := params[0]
	nDilate := params[1]
	nMorphEx := params[2]
	morphIndx := params[3]
	sWidth := 3
	sHeight := 3
	sIndx := 0

	// create the structuring element:
	var element *opencv.IplConvKernel
	if sWidth != 0 && sHeight != 0 {
		element = opencv.CreateStructuringElement(
			sWidth,                // width
			sHeight,               // height
			sWidth/2,              // X anchor
			sHeight/2,             // Y anchor
			structShapes[sIndx].C, // shape constant
		)
		defer element.ReleaseElement()
	}

	fmt.Println("****************************************")
	fmt.Printf("NErode: %d\n", nErode)
	fmt.Printf("nDilate: %d\n", nDilate)
	fmt.Printf("nMorphEx: %d\n", nMorphEx)
	fmt.Printf("morphType: %v\n", morphTypes[morphIndx])
	if element == nil {
		fmt.Println("Structuring element: nil (default 3x3 rect)")
	} else {
		fmt.Printf("Structuring element: [%dx%d] %v\n", sWidth, sHeight, structShapes[sIndx])
	}

	outputImg := img.Clone()
	defer outputImg.Release()

	// first we will erode the image..
	if nErode > 0 {
		opencv.Erode(
			outputImg, // source
			outputImg, // destination (inplace is okay here)
			element,   // structuring element (nil means default 3x3)
			nErode,    // number of iterations
		)
	}

	// next we will dilate the image...
	if nDilate > 0 {
		opencv.Dilate(
			outputImg, // source
			outputImg, // destination (inplace is okay here)
			element,   // structuring element (nil means default 3x3)
			nDilate,   // number of iterations
		)
	}

	// last we'll do the MorphologyEx()
	if nMorphEx > 0 {
		tempImg := opencv.CreateImage(img.Width(), img.Height(), opencv.IPL_DEPTH_8U, 1)
		defer tempImg.Release()

		opencv.MorphologyEx(
			outputImg,               // source image
			outputImg,               // destination (inplace is okay)
			tempImg,                 // temporary image
			element,                 // structuring element (nil means default 3x3)
			morphTypes[morphIndx].C, // morphological operation constant
			nMorphEx,                // number of iterations
		)
	}

	opencv.SaveImage(output, outputImg, nil)
}
