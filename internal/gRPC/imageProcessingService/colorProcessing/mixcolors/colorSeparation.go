package mixcolors

import (
	"github.com/lucasb-eyer/go-colorful"
	"sort"
)

func ColorSeparation(mainColorS, colorFabric string, typeS int) []BlendStructure {
	similarColors := sortByDiff(mainColorS, ColorsHex(colorFabric), typeS)
	compensationChls(mainColorS, similarColors, typeS)

	return additionShades(nil)
}

func sortByDiff(mainColorS string, colors []string, typeS int) []ColorAdditive {
	result := calculateDiff(mainColorS, colors, typeS)
	sort.Slice(result, func(i, j int) bool { return lessDiff(i, j, result) })
	return result[:5]
}

func calculateDiff(mainColorS string, colors []string, typeS int) []ColorAdditive {
	result := make([]ColorAdditive, 0)

	for i := 0; i < len(colors); i++ {

		result = append(result, ColorAdditive{
			Hex:  colors[i],
			Diff: CompareByType(mainColorS, colors[i], typeS)},
		)
	}

	return result
}

func lessDiff(i, j int, result []ColorAdditive) bool {
	return result[i].Diff[0] < result[j].Diff[0] &&
		result[i].Diff[1] < result[j].Diff[1] &&
		result[i].Diff[2] < result[j].Diff[2]
}

func compensationChls(mainC string, additives []ColorAdditive, typeS int) (result []string) {

	result = make([]string, len(additives))
	for _, additive := range additives {
		hexResult := additive.Hex
		for ch := 0; ch < DefaultChannels; ch++ {
			hexResult = addChl(mainC, hexResult, ch, typeS)
		}
		result = append(result, hexResult)
	}

	return result
}

func addChl(hex1, hex2 string, ch, typeS int) string {
	c1, _ := colorful.Hex(hex1)
	c2, _ := colorful.Hex(hex2)

	portion := calculatePortion(c1, c2, ch, typeS)
	return Blending(c2, getChlColor(ch), portion, typeS).Hex()
}

func getChlColor(ch int) colorful.Color {
	sortColorsByChls := InitColors().sortColorsByChls
	result, _ := colorful.Hex(sortColorsByChls[ch][0].Hex)
	return result
}

func Blending(c1, c2 colorful.Color, p float64, typeS int) colorful.Color {
	switch typeS {
	case TypeLAB:
		return c1.BlendLab(c2, p)
	default:
		return colorful.Color{}
	}
}

func calculatePortion(c1, c2 colorful.Color, ch, typeS int) float64 {
	switch typeS {
	case TypeLAB:
		return calculatePortionLab(c1, c2, ch)
	default:
		return 0.0
	}
}

func calculatePortionLab(c0, c1 colorful.Color, ch int) float64 {
	l0, a0, b0 := c0.Lab()
	l1, a1, b1 := c1.Lab()

	l2, a2, b2 := getChlColor(ch).Lab()

	switch ch {
	case 0:
		return (l0 - l1)/(l2-l1)
	case 1:
		return (a0 - a1)/(a2-a1)
	case 2:
		return (b0 - b1)/(b2-b1)
	default:
		return 0.0
	}
}

func additionShades([]ColorAdditive) []BlendStructure {
	return nil
}
