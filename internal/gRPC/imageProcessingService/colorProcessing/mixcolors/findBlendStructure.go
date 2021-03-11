package mixcolors

import (
	"encoding/json"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"math"
	"paint/internal/utils"
	"sort"
	"strconv"
	"strings"
)

const (
	DefaultCountOfColors = 5
	DefaultChannelsLUV   = 3
	DefaultDifferenceLUV = 0.05
	P                    = "."
	Slash                = "/"
	Lattice              = "#"
)

type ColorAdditiveLUV struct {
	Hex        string
	Difference []float64
}

type SimilarColor struct {
	Structure  string
	Difference float64
}

type BlendStructure struct {
	C1Hex string
	C2Hex string
	C3Hex string

	C2Portion string
	C3Portion string

	ResultHex string
}

func FindBlendStructureAmongFabricColorsLUV(mainColorS, colorFabric string) []BlendStructure {
	sortAdditiveColors := make(map[int][]ColorAdditiveLUV)
	colorsHexValues := make([]string, 0)

	for _, colorAsset := range InitColors().mapOfAllColors[colorFabric] {
		colorsHexValues = append(colorsHexValues, colorAsset.hex)
	}

	for i := 0; i < DefaultChannelsLUV; i++ {
		sortAdditiveColors[i] = getSortAdditiveColorsByChanelLUV(mainColorS, colorsHexValues, i, DefaultCountOfColors)
	}

	blendStructures := make([]BlendStructure, 0)

	cs := getSimilarColorsLUV(mainColorS, sortAdditiveColors)
	for _, similarColor := range cs {
		values := strings.Split(similarColor.Structure, Slash)

		blendStructures = append(blendStructures, BlendStructure{
			C1Hex:     values[0],
			C2Hex:     "#" + strings.Split(values[1], Lattice)[1],
			C3Hex:     "#" + strings.Split(values[2], Lattice)[1],
			C2Portion: strings.Split(values[1], Lattice)[0],
			C3Portion: strings.Split(values[2], Lattice)[0],
			ResultHex: values[len(values)-1],
		})
	}

	if DistanceLuv(mainColorS, blendStructures[0].ResultHex) > 0.1 {
		FindBlendStructureAmongFabricColorsLUV(mainColorS, colorFabric)
	}
	return blendStructures
}

func getSimilarColorsLUV(mainColorS string, additiveColors map[int][]ColorAdditiveLUV) []SimilarColor {
	colors := make([]string, 0)
	resultColors := make([]string, 0)

	for k, value := range additiveColors {
		for i := 0; i < len(value); i++ {
			if !utils.Contains(colors, value[i].Hex) {
				colors = append(colors, fmt.Sprintf(value[i].Hex+P+strconv.Itoa(k)))
			}
		}
	}

	combinations := Combinations(colors, DefaultChannelsLUV)

	for i := 0; i < len(combinations); i++ {
		if isCombinationRight(combinations[i]) {
			blendResult := Blend3Colors(Color3{
				C1: strings.Split(combinations[i][0], P)[0],
				C2: strings.Split(combinations[i][1], P)[0],
				C3: strings.Split(combinations[i][2], P)[0]}, DefaultNumberOfShades)

			for C2, child := range blendResult.S(strings.Split(combinations[i][0], P)[0]).ChildrenMap() {
				for C3, value := range child.ChildrenMap() {
					var blendColors map[string]interface{}
					json.Unmarshal([]byte(value.String()), &blendColors)

					for shade, v := range blendColors {
						if !utils.Contains(resultColors, v.(string)) {
							result := fmt.Sprintf("%v/%v/%v%v/%v", strings.Split(combinations[i][0], P)[0],
								C2, shade, C3, v.(string))

							resultColors = append(resultColors, result)
						}
					}
				}
			}
		}
	}
	return sortResultColorsByDifferenceLUV(mainColorS, resultColors)
}

func getSortAdditiveColorsByChanelLUV(mainColorS string, additiveColors []string, channel, count int) []ColorAdditiveLUV {
	params := make([]ColorAdditiveLUV, 0)

	cMain, _ := colorful.Hex(mainColorS)
	l, u, v := cMain.Luv()

	for i := 0; i < len(additiveColors); i++ {
		cAdditive, _ := colorful.Hex(additiveColors[i])
		li, ui, vi := cAdditive.Luv()

		params = append(params, ColorAdditiveLUV{
			Hex:        cAdditive.Hex(),
			Difference: []float64{math.Abs(l - li), math.Abs(u - ui), math.Abs(v - vi)},
		})
	}

	sort.Slice(params, func(i, j int) bool {
		return lessFunction(i, j, channel, params)
	})

	return params[:count]
}

func lessFunction(i, j, channel int, params []ColorAdditiveLUV) bool {
	return params[i].Difference[channel] < params[j].Difference[channel] &&
		params[i].Difference[getOtherChannels(channel)[0]] < DefaultDifferenceLUV &&
		params[i].Difference[getOtherChannels(channel)[1]] < DefaultDifferenceLUV
}

func getOtherChannels(channel int) []int {
	otherChannels := make([]int, 2)

	for i := 0; i < DefaultChannelsLUV; i++ {
		if i != channel {
			otherChannels = append(otherChannels, i)
		}
	}
	return otherChannels
}

func sortResultColorsByDifferenceLUV(mainColorS string, colors []string) []SimilarColor {
	result := make([]SimilarColor, 0)

	for i := 0; i < len(colors); i++ {
		values := strings.Split(colors[i], Slash)
		result = append(result, SimilarColor{
			Structure:  colors[i],
			Difference: DistanceLuv(mainColorS, values[len(values)-1]),
		})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Difference < result[j].Difference })
	return result[:5]
}

func DistanceLuv(colorS1, colorS2 string) float64 {
	c1, _ := colorful.Hex(colorS1)
	c2, _ := colorful.Hex(colorS2)

	l1, u1, v1 := c1.Luv()
	l2, u2, v2 := c2.Luv()

	return math.Sqrt(utils.Sq(l1-l2) + utils.Sq(u1-u2) + utils.Sq(v1-v2))
}

func isCombinationRight(elements []string) bool {
	channels := make([]string, 0)
	for i := 0; i < len(elements); i++ {
		channels = append(channels, strings.Split(elements[i], P)[1])
	}

	for i := 0; i < len(channels); i++ {
		if !utils.Contains(channels, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}
