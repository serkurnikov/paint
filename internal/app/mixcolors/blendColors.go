package mixcolors

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/lucasb-eyer/go-colorful"
)

const P = "."
const DefaultNumberOfShades = 5

func BlendColors(colorS1, colorS2 string, numberOfShades int) map[float64]colorful.Color {

	colors := make(map[float64]colorful.Color)

	c1, _ := colorful.Hex(colorS1)
	c2, _ := colorful.Hex(colorS2)

	for i := 0; i < numberOfShades; i++ {
		portion := float64(i) / float64(numberOfShades-1)
		colors[portion] = c1.BlendLuv(c2, portion)
	}

	/*{1, "#124"}, {2, "#124"}, {3, "#124"}, {4, "#124"}, {5, "#124"}*/

	return colors
}

func BlendMainColorsWithShades(colorMainS string, arrayOfSecondaryColorsShades []string) *gabs.Container {
	jsonObj := gabs.New()
	jsonObj.ArrayOfSize(len(arrayOfSecondaryColorsShades), colorMainS)

	for i := 0; i < len(arrayOfSecondaryColorsShades); i++ {
		result := BlendColors(colorMainS, arrayOfSecondaryColorsShades[i], DefaultNumberOfShades)
		jsonObj.SetP(result, colorMainS+P+arrayOfSecondaryColorsShades[i])
		/* colorMainS {
				arrayOfSecondaryColorsShades[1] {
					{1, "#124"}, {2, "#124"}, {3, "#124"}, {4, "#124"}, {5, "#124"}
				}

				arrayOfSecondaryColorsShades[2] {
					{1, "#124"}, {2, "#124"}, {3, "#124"}, {4, "#124"}, {5, "#124"}
				}
		*/
	}
	return jsonObj
}

func BlendCombination(combinationElements []string) *gabs.Container {
	jsonObj := gabs.New()

	for i := 0; i < len(combinationElements)-1; i++ {
		if i == 0 {
			/*
			result := BlendColors(combinationElements[0], combinationElements[1], DefaultNumberOfShades)
			combinationElements[0] / combinationElements[1]
			result {1, "#124"}, {2, "#124"}, {3, "#124"}, {4, "#124"}, {5, "#124"}
			for (< result/Count) {
				BlendMainColorsWithShades()
			}
			*/
		} else {

		}
	}
	return jsonObj
}

func mixColors(colorsDataS []string) {
	subsets := All(colorsDataS)
	for i := 0; i < len(subsets); i++ {
		BlendCombination(subsets[i])
	}
}
