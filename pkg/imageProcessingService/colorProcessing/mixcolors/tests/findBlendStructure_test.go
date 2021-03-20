package tests

import (
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"paint/internal/utils"
	"paint/pkg/imageProcessingService/colorProcessing/mixcolors"
	"testing"
)

const (
	testFindBlendStructure = "\\testBlendStructure.jpg"
)

func TestBlendStructure(t *testing.T) {
	portion := 0.5
	c1, _ := colorful.Hex("#11317a")
	c2, _ := colorful.Hex("#a14e23")
	c3, _ := colorful.Hex("#e3b9a1")

	mainColorS := c1.BlendLuv(c2, portion).BlendLuv(c3, portion).Hex()
	blendStructures := mixcolors.BlendStructureAmongFabricColors(mainColorS, mixcolors.MasterColors,
		mixcolors.TypeLABCIEDE2000)

	m, _ := colorful.Hex(mainColorS)
	r, _ := colorful.Hex(blendStructures[0].ResultHex)

	//Additive colors
	c21, _ := colorful.Hex(blendStructures[0].C1Hex)
	c22, _ := colorful.Hex(blendStructures[0].C2Hex)
	c23, _ := colorful.Hex(blendStructures[0].C3Hex)

	colors := []colorful.Color{m, r, c21, c22, c23}

	diff := mixcolors.DifferenceByType(m, r, mixcolors.TypeLABCIEDE2000)

	if diff > mixcolors.DefaultDiff {
		t.Errorf("Incorrect distance got %f, wait < %v,", diff, mixcolors.DefaultDiff)
	} else {
		t.Logf("MainColors %s, BlendResultColor %s, Distance %f", mainColorS, blendStructures[0].ResultHex, diff)
		t.Logf("BlendResult: \nС1DIF = %v, \nP2 = %v, \nС2DIF = %v, \nP3 = %v, \nС3DIF = %v",
			blendStructures[0].C1Hex,
			blendStructures[0].C2Portion,
			blendStructures[0].C2Hex,
			blendStructures[0].C3Portion,
			blendStructures[0].C3Hex)
	}

	result := image.NewRGBA(image.Rect(0, 0, mixcolors.WEIGHT, mixcolors.HEIGHT))

	h := result.Bounds().Max.X / len(colors)
	w := result.Bounds().Max.Y

	for i := 0; i < len(colors); i++ {
		r, g, b, a := colors[i].RGBA()
		currentColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
		draw.Draw(result, image.Rect(0, i*h, w, i*h+h), &image.Uniform{C: currentColor}, image.Point{
			X: 0,
			Y: 0,
		}, draw.Src)
	}

	currentDir, _ := utils.GetCurrentDir()
	file, _ := os.Create(currentDir + testFindBlendStructure)
	png.Encode(file, result)
}

/*Latest Test
BenchmarkSample
BenchmarkSample-6   	       3	 470548300 ns/op
*/
func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mixcolors.BlendStructureAmongFabricColors("#85739b", mixcolors.MasterColors, mixcolors.TypeLABCIEDE2000)
	}
}
