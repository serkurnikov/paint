package mixcolors
//https://www.imgonline.com.ua/

import (
	"encoding/json"
	"image"
	"io/ioutil"
	"log"
	"paint/internal/gRPC/imageProcessingService/colorProcessing/prominentcolor"
	"paint/internal/utils"
	"strings"
	"sync"
)
import "paint/assets"

const (
	root = "assets/structure.json"
	Path = "Path"
	MasterColors = "masters_colors"
)

type ColorAsset struct {
	name string
	hex  string
	path []string
}

type singleton struct {
	mapOfAllColors map[string][]ColorAsset
}

var instance *singleton
var once sync.Once

func InitColors() *singleton {
	once.Do(func() {
		instance = &singleton{}
		instance.mapOfAllColors = createMapOfAllColors()
	})
	return instance
}

func GetMastersColors() map[string]string {
	return map[string]string{
		"ararat_green":                "#322c26",
		"belila_titan":                "#e3e4e8",
		"belila_zink":                 "#e9e8e9",
		"berlin_azur":                 "#253455",
		"black_travertine":            "#3a3635",
		"blue":                        "#11317a",
		"burnt_bone":                  "#3d3938",
		"burnt_sienna":                "#673023",
		"burnt_umber":                 "#462c26",
		"cadmin_lemon":                "#fdf90b",
		"cadmin_orange":               "#fe8e01",
		"cadmin_red_dark":             "#fe060a",
		"cadmin_red_light":            "#fe1202",
		"cadmin_yellow_dark":          "#febf32",
		"cadmin_yellow_light":         "#fef406",
		"cadmin_yellow_medium":        "#fedf26",
		"cadmium_yellow_waik":         "#a75419",
		"caput_mortuum_dark":          "#602828",
		"carbon_black_gas":            "#38312f",
		"ceruleum_blue":               "#0091f6",
		"cherry_marcar":               "#9b2d19",
		"cherry_tavush":               "#73291b",
		"chrome_cobalt_blue_green":    "#026fab",
		"chrome_cobalt_green_blue":    "#076d95",
		"cinnabar":                    "#fe0600",
		"cinnabar_yellow_green":       "#dbda4c",
		"cobalt_blue_medium":          "#033887",
		"cobalt_blue_spectral":        "#0e1b67",
		"cobalt_green_dark":           "#276f5e",
		"cobalt_green_light":          "#1ba582",
		"cobalt_violet_dark":          "#6d1951",
		"cobalt_violet_light":         "#a91173",
		"coral_pink":                  "#fe9495",
		"dark_ochre_kotayk":           "#9e4f18",
		"emerald":                     "#17453f",
		"english_green_dark":          "#266032",
		"english_green_light":         "#317b2a",
		"english_red":                 "#b6230b",
		"geranium_red":                "#fd0a23",
		"glauconite":                  "#293123",
		"golden_arzakan":              "#a14e23",
		"golden_yellow":               "#fe9901",
		"grass_green":                 "#2c3c1b",
		"green_bzhni":                 "#725a39",
		"green_FC":                    "#0e4848",
		"green_tavush":                "#425438",
		"gutangira_purple":            "#58211c",
		"gutangira_raspberry":         "#792a1a",
		"indanthrene_blue":            "#210c4d",
		"indian_red":                  "#7f2f29",
		"indian_yellow":               "#bc3d2a",
		"indigo":                      "#32313b",
		"irgazin_yellow":              "#e9bd00",
		"kraplak_pink_solid":          "#fd002a",
		"kraplak_purple_durable":      "#be0824",
		"light_malachite":             "#61e297",
		"lilac_hinacridon":            "#500220",
		"manganese_violet_light":      "#7d264b",
		"mars_black":                  "#322d2c",
		"mars_black_warm":             "#35302f",
		"mars_brown_dark":             "#4c2f26",
		"mars_brown_dark_transparent": "#363130",
		"mars_brown_light":            "#4e332b",
		"mars_orange_transparent":     "#3b251f",
		"mars_yellow_transparent":     "#452619",
		"may_green":                   "#17aa20",
		"natural_umber":               "#302b28",
		"neapolitan_corporeal":        "#f6d9bc",
		"neapolitan_pink":             "#e3b9a1",
		"neapolitan_yellow_fawn":      "#e1d0b5",
		"neopolitan_yellow":           "#fea217",
		"ochre_golden":                "#c1672b",
		"ochre_light":                 "#c5680d",
		"ochre_red":                   "#a03928",
		"ochre_yellow":                "#fea316",
		"ochre_yellow_marcar":         "#c26a26",
		"olive":                       "#424a16",
		"olive_arzakan":               "#4a4a3b",
		"orange":                      "#fe600d",
		"orange_travertine":           "#e44c1d",
		"oxide_of_chromium":           "#446e31",
		"payne_grey":                  "#2b2f3c",
		"pink_hinacridon":             "#d7054c",
		"purple_brown_sevan":          "#6a453f",
		"purple_dark_hinacridon":      "#ae010b",
		"purple_gray_lori":            "#5a4340",
		"red_brown_sevan":             "#963416",
		"red_brown_waik":              "#7c331b",
		"red_hinacridon":              "#fd0102",
		"royal_blue":                  "#76a9fb",
		"shakhnazarskaya_red":         "#982610",
		"sienna_natural":              "#9a6344",
		"sky_blue":                    "#00a6fe",
		"st_petersburg_gray":          "#c4c1b9",
		"st_petersburg_pink":          "#fe8aad",
		"st_petersburg_purple":        "#e878bf",
		"strontium_yellow":            "#f8ef69",
		"tiondigo_pink":               "#fe053a",
		"turkish_blue":                "#75f1f3",
		"turkish_green":               "#95f1ec",
		"turquoise":                   "#057d9c",
		"ultramarine_dark":            "#0c166d",
		"ultramarine_light":           "#0217a3",
		"ultramarine_pink":            "#8f2b4a",
		"ultramarine_purple":          "#311e4f",
		"van_dik_brown":               "#2d2726",
		"venetian_purple":             "#5e051a",
		"venetian_red":                "#9f0103",
		"viridian_green":              "#233f28",
		"volchonskoite":               "#344434",
		"yellow_travertine":           "#b77111"}
}

func createMapOfAllColors() map[string][]ColorAsset {

	mapOfAllColors := make(map[string][]ColorAsset)
	colors := make([]ColorAsset, 0)
	data := getAssetsData()

	mastersColors, _ := utils.StructToMap(data.Assets[0].ColorsFabric.MastersColors)
	for colorName, images := range mastersColors {

		var imagesPath = make([]string, 0)
		for _, image := range images.([]interface{}) {
			for key, value := range image.(map[string]interface{}) {
				if strings.Compare(key, Path) == 0 {
					imagesPath = append(imagesPath, value.(string))
				}
			}
		}

		colors = append(colors, ColorAsset{
			name: colorName,
			hex:  GetMastersColors()[colorName],
			path: imagesPath,
		})
	}

	mapOfAllColors[MasterColors] = colors
	return mapOfAllColors
}

func calculateDominantColorByImageAssets(imagesPath []string) string {

	imagesPalette := make(map[image.Image][]string)
	for _, path := range imagesPath {
		img, cols := prominentcolor.BuildP(path, prominentcolor.DefaultK)
		hexColors := make([]string, 0)
		for i := 0; i < len(cols); i++ {
			hexColors = append(hexColors, "#"+cols[i].AsString())
		}
		imagesPalette[img] = hexColors
	}

	println(imagesPath[0])
	for _, hexes := range imagesPalette {
		for _, hex := range hexes {
			//TODO calculate color hex by images
			println(hex)
		}
	}

	return "#nil"
}

func getAssetsData() assets.Assets {
	data, err := ioutil.ReadFile(root)
	assets := assets.Assets{}
	err = json.Unmarshal(data, &assets)
	if err != nil {
		log.Fatalln(err)
	}
	return assets
}
