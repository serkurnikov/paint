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
	TypeLuv = iota
	TypeRGB
	TypeLinearRGB
	TypeLAB
	TypeLABCIE94
	TypeLABCIEDE2000
	TypeHSV
	TypeHCL

	DefaultCountOfColors = 5
	DefaultChannels      = 3
	DefaultDiff          = 0.2

	P       = "."
	Slash   = "/"
	Lattice = "#"
)

type BlendStructure struct {
	C1Hex string
	C2Hex string
	C3Hex string

	C2Portion string
	C3Portion string

	ResultHex string
}

type ColorAdditive struct {
	Hex  string
	Diff []float64
}

func BlendStructureAmongFabricColors(mainColorS, colorFabric string, spaceType int) []BlendStructure {
	for {
		additiveColors := make(map[int][]ColorAdditive)

		for ch := 0; ch < DefaultChannels; ch++ {
			additiveColors[ch] = sortAdditiveColorsByChannel(mainColorS,
				ColorsHex(colorFabric), ch, spaceType, DefaultCountOfColors)
		}

		colors := SimilarColors(mainColorS, additiveColors, spaceType)
		blendStructures := fillBlendStructure(colors)

		c1, _ := colorful.Hex(mainColorS)
		c2, _ := colorful.Hex(blendStructures[0].ResultHex)

		if DifferenceByType(c1, c2, spaceType) < DefaultDiff {
			return blendStructures
		}
	}
}

func ColorsHex(colorFabric string) []string {
	colorsHexValues := make([]string, 0)

	for _, colorAsset := range InitColors().mapOfAllColors[colorFabric] {
		colorsHexValues = append(colorsHexValues, colorAsset.hex)
	}

	return colorsHexValues
}

func CompareByType(c1, c2 string, t int) []float64 {
	m1, m2, m3 := getColorArgsByType(c1, t)
	a1, a2, a3 := getColorArgsByType(c2, t)

	return []float64{math.Abs(m1 - a1), math.Abs(m2 - a2), math.Abs(m3 - a3)}
}

func DifferenceByType(c1, c2 colorful.Color, spaceType int) float64 {
	switch spaceType {
	case TypeLuv:
		return c1.DistanceLuv(c2)
	case TypeLAB:
		return c1.DistanceLab(c2)
	case TypeHSV:
		return c1.DistanceHSLuv(c2)
	case TypeHCL:
		return c1.DistanceHPLuv(c2)
	case TypeLABCIE94:
		return c1.DistanceCIE94(c2)
	case TypeLABCIEDE2000:
		return c1.DistanceCIEDE2000(c2)
	case TypeRGB:
		return c1.DistanceRgb(c2)
	case TypeLinearRGB:
		return c1.DistanceLinearRGB(c2)
	default:
		return 0
	}
}

func OtherChls(channel int) []int {
	otherChannels := make([]int, DefaultChannels-1)

	for i := 0; i < DefaultChannels; i++ {
		if i != channel {
			otherChannels = append(otherChannels, i)
		}
	}
	return otherChannels
}

func SimilarColors(mainColorS string, additiveColors map[int][]ColorAdditive, spaceType int) []ColorAdditive {
	colors := make([]string, 0)

	for k, value := range additiveColors {
		for i := 0; i < len(value); i++ {
			if !utils.Contains(colors, value[i].Hex) {
				colors = append(colors, fmt.Sprintf(value[i].Hex+P+strconv.Itoa(k)))
			}
		}
	}

	combinations := Combinations(colors, DefaultChannels)

	blendingResultColors := getBlendingColorsResult(combinations)

	return sortResultColorsByDifference(mainColorS, blendingResultColors, spaceType)
}


func isCombinationRight(elements []string) bool {
	chls := make([]string, 0)
	for i := 0; i < len(elements); i++ {
		chls = append(chls, strings.Split(elements[i], P)[1])
	}

	for i := 0; i < len(chls); i++ {
		if !utils.Contains(chls, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

func fillBlendStructure(colors []ColorAdditive) []BlendStructure {
	blendStructures := make([]BlendStructure, 0)

	for _, similarColor := range colors {
		values := strings.Split(similarColor.Hex, Slash)
		blendStructures = append(blendStructures, BlendStructure{
			C1Hex:     values[0],
			C2Hex:     Lattice + strings.Split(values[1], Lattice)[1],
			C3Hex:     Lattice + strings.Split(values[2], Lattice)[1],
			C2Portion: strings.Split(values[1], Lattice)[0],
			C3Portion: strings.Split(values[2], Lattice)[0],
			ResultHex: values[len(values)-1],
		})
	}

	return blendStructures
}

func getBlendingColorsResult(c [][]string) []string {
	resultColors := make([]string, 0)

	for i := 0; i < len(c); i++ {
		if isCombinationRight(c[i]) {
			blendResult := Blend3Colors(Color3{
				C1: strings.Split(c[i][0], P)[0],
				C2: strings.Split(c[i][1], P)[0],
				C3: strings.Split(c[i][2], P)[0]}, DefaultNumberOfShades)

			for C2, child := range blendResult.S(strings.Split(c[i][0], P)[0]).ChildrenMap() {
				for C3, value := range child.ChildrenMap() {
					var blendColors map[string]interface{}
					_ = json.Unmarshal([]byte(value.String()), &blendColors)

					for shade, v := range blendColors {
						if !utils.Contains(resultColors, v.(string)) {
							result := fmt.Sprintf("%v/%v/%v%v/%v", strings.Split(c[i][0], P)[0],
								C2, shade, C3, v.(string))

							resultColors = append(resultColors, result)
						}
					}
				}
			}
		}
	}

	return resultColors
}

func getColorArgsByType(color string, spaceType int) (float64, float64, float64) {
	c, _ := colorful.Hex(color)
	switch spaceType {
	case TypeLuv:
		return c.Luv()
	case TypeLAB, TypeLABCIE94, TypeLABCIEDE2000:
		return c.Lab()
	case TypeHSV:
		return c.Hsv()
	case TypeHCL:
		return c.Hcl()
	case TypeRGB, TypeLinearRGB:
		return c.R, c.G, c.B
	default:
		return 0, 0, 0
	}
}

func lessFunction(i, j, ch int, c []ColorAdditive) bool {
	chls := OtherChls(ch)
	return c[i].Diff[ch] < c[j].Diff[ch] &&
		c[i].Diff[chls[0]] < DefaultDiff &&
		c[i].Diff[chls[1]] < DefaultDiff
}

func sortAdditiveColorsByChannel(mainColorS string, additiveColors []string, channel, spaceType int, count int) []ColorAdditive {
	additives := make([]ColorAdditive, 0)

	for i := 0; i < len(additiveColors); i++ {

		additives = append(additives, ColorAdditive{
			Hex:  additiveColors[i],
			Diff: CompareByType(mainColorS, additiveColors[i], spaceType),
		})
	}

	sort.Slice(additives, func(i, j int) bool {
		return lessFunction(i, j, channel, additives)
	})

	return additives[:count]
}

func sortResultColorsByDifference(mainColorS string, colors []string, spaceType int) []ColorAdditive {
	result := make([]ColorAdditive, 0)

	for i := 0; i < len(colors); i++ {
		values := strings.Split(colors[i], Slash)
		c1, _ := colorful.Hex(mainColorS)
		c2, _ := colorful.Hex(values[len(values)-1])

		result = append(result, ColorAdditive{
			Hex:  colors[i],
			Diff: []float64{DifferenceByType(c1, c2, spaceType), 0.0, 0.0},
		})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Diff[0] < result[j].Diff[0] })
	return result[:5]
}
