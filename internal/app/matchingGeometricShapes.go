package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
	"paint/internal/gRPC/imageProcessingService/contoursProcessing"
	"path/filepath"
)

func FindingMatchingGeometricShapes(in string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	countours := contoursProcessing.DrawCustomContours(in)
	for _, countour := range countours {
		displayLineByImagePoints(img, countour)
	}

	if ok := gocv.IMWrite(palettePathImage, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}

func displayLineByImagePoints(mat gocv.Mat, imagePoints []image.Point) {

	colorCountours := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	a, b, x1, x2 := getApproxLine(imagePoints)
	pt1 := image.Point{X: x1, Y: int(a)*x1 + int(b)}
	pt2 := image.Point{X: x2, Y: int(a)*x2 + int(b)}

	gocv.Line(&mat, pt1, pt2, colorCountours, 1)
}

func getApproxLine(point []image.Point) (a, b float64, x1, x2 int) {
	var sumX = 0.0
	var sumY = 0.0
	var sumX2 = 0.0
	var sumXY = 0.0
	N := float64(len(point))

	for i := 0; i < len(point); i++ {
		sumX += float64(point[i].X)
		sumY += float64(point[i].Y)
		sumX2 += float64(point[i].X * point[i].X)
		sumXY += float64(point[i].X * point[i].Y)
	}

	x1, x2 = FindMinAndMax(point)

	a = (N*sumXY - (sumX * sumY)) / (N*sumX2 - sumX*sumX)
	b = (sumY - a*sumX) / N
	return a, b, x1, x2
}

func FindMinAndMax(point []image.Point) (min int, max int) {
	min = point[0].X
	max = point[0].X
	for _, value := range point {
		if value.X < min {
			min = value.X
		}
		if value.X > max {
			max = value.X
		}
	}
	return min, max
}
