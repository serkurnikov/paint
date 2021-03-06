package mixcolors

import (
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
	"github.com/lucasb-eyer/go-colorful"
	"strconv"
	"strings"
)

const DefaultNumberOfShades = 5

type Color3 struct {
	C1, C2, C3 string
}

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
		if !jsonObj.Exists(arrayOfColors[i]) {
			result := BlendColors(colorMainS, arrayOfColors[i], numberOfShades)
			jsonObj.Set(result, arrayOfColors[i])
		}
	}
	return jsonObj
}

//:TODO change function (recursive method)
func BlendCombination(combinationElements []string, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()
	for i := 0; i < len(combinationElements)-1; i++ {
		result := BlendMainColorsWithArrayOfColors(combinationElements[i], combinationElements[i+1:], numberOfShades)
		jsonObj.Set(result, combinationElements[i])
	}
	return jsonObj
}

func Blend3Colors(color Color3, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()
	var colors map[string]interface{}
	json.Unmarshal([]byte(BlendColors(color.C1, color.C2, numberOfShades).String()), &colors)

	for k, v := range colors {
		blending := BlendColors(color.C2, v.(string), numberOfShades)
		jsonObj.Set(blending, color.C1, k+color.C2, v.(string))
	}
	return jsonObj
}

func GetAllMixColors(colorsDataS []string, numberOfShades int) *gabs.Container {
	jsonObj := gabs.New()
	subsets := All(colorsDataS)
	for i := 0; i < len(subsets); i++ {
		result := BlendCombination(subsets[i], numberOfShades)
		jsonObj.Set(result, strings.Join(subsets[i], ""))
	}
	return jsonObj
}
