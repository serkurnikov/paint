package mixcolors

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/lucasb-eyer/go-colorful"
	"strconv"
)

const MixedColors = "MixedColors"
const DefaultNumberOfShades = 5

func BlendColors(colorS1, colorS2 string, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()

	c1, _ := colorful.Hex(colorS1)
	c2, _ := colorful.Hex(colorS2)

	for i := 0; i < numberOfShades; i++ {
		portion := float64(i) / float64(numberOfShades-1)
		jsonObj.Set(c1.BlendLuv(c2, portion).Hex(), strconv.FormatFloat(portion, 'f', 6, 64))
	}
	return jsonObj
}

func BlendMainColorsWithArrayOfColors(colorMainS string, arrayOfColors []string, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()
	for i := 0; i < len(arrayOfColors); i++ {
		if !jsonObj.Exists(colorMainS+arrayOfColors[i]) {
			result := BlendColors(colorMainS, arrayOfColors[i], numberOfShades)
			jsonObj.Set(result, colorMainS+arrayOfColors[i])
		}
	}
	return jsonObj
}

func BlendCombination(combinationElements []string, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()
	result := BlendMainColorsWithArrayOfColors(combinationElements[0], combinationElements[1:], numberOfShades)
	jsonObj.Set(result, combinationElements[0])
	println(jsonObj.String())
	return jsonObj
}

func GetAllMixColors(colorsDataS []string, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()
	jsonObj.Array(MixedColors)

	subsets := All(colorsDataS)
	for i := 0; i < len(subsets); i++ {
		result := BlendCombination(subsets[i], numberOfShades)
		jsonObj.ArrayAppend(result, MixedColors)
	}
	return jsonObj
}