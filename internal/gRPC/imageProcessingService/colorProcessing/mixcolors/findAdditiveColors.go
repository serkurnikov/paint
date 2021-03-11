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
	P                    = "."
	Slash                = "/"
)

type ColorAdditiveLUV struct {
	hex        string
	difference []float64
}

type SimilarColor struct {
	hex        string
	difference float64
}

func FindAdditiveColorsLUVInterColorFabric(mainColorS, colorFabric string) []ColorAdditiveLUV {
	sortAdditiveColors := make(map[int][]ColorAdditiveLUV)
	colorsHexValues := make([]string, 0)

	for _, colorAsset := range InitColors().mapOfAllColors[colorFabric] {
		colorsHexValues = append(colorsHexValues, colorAsset.hex)
	}

	for i := 0; i < DefaultChannelsLUV; i++ {
		sortAdditiveColors[i] = getSortAdditiveColorsByChanelLUV(mainColorS, colorsHexValues, i, DefaultCountOfColors)
	}

	getSimilarColorsLUV(mainColorS, sortAdditiveColors)

	return nil
}

func getSimilarColorsLUV(mainColorS string, additiveColors map[int][]ColorAdditiveLUV) []string {
	colors := make([]string, 0)
	resultColors := make([]string, 0)

	for k, value := range additiveColors {
		for i := 0; i < len(value); i++ {
			if !utils.Contains(colors, value[i].hex) {
				colors = append(colors, fmt.Sprintf(value[i].hex+P+strconv.Itoa(k)))
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

			for _, child := range blendResult.S(strings.Split(combinations[i][0], P)[0]).Children() {
				for _, value := range child.Children() {
					var blendColors map[string]interface{}
					json.Unmarshal([]byte(value.String()), &blendColors)

					for shade, v := range blendColors {
						if !utils.Contains(resultColors, v.(string)) {
							//#213f24/0.5/#123f22
							resultColors = append(resultColors, combinations[i][0]+Slash+shade+Slash+v.(string))
						}
					}
				}
			}
		}
	}
	//test(mainColorS, resultColors)
	return nil
}

func getSortAdditiveColorsByChanelLUV(mainColorS string, additiveColors []string, channel, count int) []ColorAdditiveLUV {
	params := make([]ColorAdditiveLUV, 0)

	cMain, _ := colorful.Hex(mainColorS)
	l, u, v := cMain.Luv()

	for i := 0; i < len(additiveColors); i++ {
		cAdditive, _ := colorful.Hex(additiveColors[i])
		li, ui, vi := cAdditive.Luv()

		params = append(params, ColorAdditiveLUV{
			hex:        cAdditive.Hex(),
			difference: []float64{math.Abs(l - li), math.Abs(u - ui), math.Abs(v - vi)},
		})
	}

	sort.Slice(params, func(i, j int) bool { return params[i].difference[channel] < params[j].difference[channel] })
	return params[:count]
}

func test(mainColorS string, colors []string) []SimilarColor {
	result := make([]SimilarColor, 0)

	for i := 0; i < len(colors); i++ {
		result = append(result, SimilarColor{
			hex:        colors[i],
			difference: DistanceLuv(mainColorS, colors[i]),
		})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].difference < result[j].difference })
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
