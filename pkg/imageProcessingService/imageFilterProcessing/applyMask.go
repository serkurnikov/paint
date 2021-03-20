package imageFilterProcessing

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
)

func applyMask(in, out, mask string) {
	img := gocv.IMRead(in, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}
	defer img.Close()


	m := gocv.IMRead(mask, gocv.IMReadColor)
	if m.Empty() {
		fmt.Printf("Failed to img image: %s\n", in)
		os.Exit(1)
	}
	defer m.Close()

	gocv.Resize(img, &img, image.Pt(m.Cols(), m.Rows()), 0, 0, gocv.InterpolationLinear)

	dest := gocv.NewMat()
	defer dest.Close()

	img.CopyToWithMask(&dest, m)

	if ok := gocv.IMWrite(out, dest); !ok {
		fmt.Printf("Failed to write image\n")
		os.Exit(1)
	}
}
