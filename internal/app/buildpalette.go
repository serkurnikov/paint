package app

import "gocv.io/x/gocv"

func buildPalette(in, out string) {
	gocv.BuildPalette(in, out)
}