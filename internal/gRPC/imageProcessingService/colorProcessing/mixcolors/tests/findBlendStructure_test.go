package tests

import (
	"github.com/lucasb-eyer/go-colorful"
	"paint/internal/gRPC/imageProcessingService/colorProcessing/mixcolors"
	"testing"
)

func TestFindBlendStructure(t *testing.T) {
	portion := 0.5
	c1, _ := colorful.Hex("#322c26") //ararat_green
	c2, _ := colorful.Hex("#fe8e01") //cadmin_orange
	c3, _ := colorful.Hex("#0091f6") //ceruleum_blue

	mainColorS := c1.BlendLuv(c2, portion).BlendLuv(c3, portion).Hex()
	blendStructure := mixcolors.FindBlendStructureAmongFabricColorsLUV(mainColorS, mixcolors.MasterColors)

	distance := mixcolors.DistanceLuv(mainColorS, blendStructure[0].ResultHex)

	if distance > 0.1 {
		t.Errorf("Incorrect distanceLUV got %f, wait < 0.1", distance)
	} else {
		t.Logf("MainColors %s, BlendResultColor %s, DistanceLUV %f", mainColorS, blendStructure[0].ResultHex, distance)
		t.Logf("BlendResult: \nС1DIF = %v, \nP2 = %v, \nС2DIF = %v, \nP3 = %v, \nС3DIF = %v",
			blendStructure[0].C1Hex,
			blendStructure[0].C2Portion,
			blendStructure[0].C2Hex,
			blendStructure[0].C3Portion,
			blendStructure[0].C3Hex)
	}
}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//TODO optimized function
		mixcolors.FindBlendStructureAmongFabricColorsLUV("#85739b", mixcolors.MasterColors)
	}
}
