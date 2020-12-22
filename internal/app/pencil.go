package app

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"path/filepath"
)

func pencil(in, out string) {
	path := filepath.Join(in)
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", path)
		os.Exit(1)
	}
	defer img.Close()

	gocv.BitwiseNot(img, &img)

	if ok := gocv.IMWrite(out, img); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
