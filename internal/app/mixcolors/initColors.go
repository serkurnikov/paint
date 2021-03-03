package mixcolors

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"paint/internal/utils/filesmanager"
)
import "paint/assets"

const root = "assets/structure.json"

type ColorAsset struct {
	name string
	hex  string
	path string
}

func InitColors() []ColorAsset {

	colors := make([]ColorAsset, 0)
	data := GetAssetsData()

	for _, a := range data.Assets {
		mastersColors, _ := filesmanager.StructToMap(a.ColorsFabric.MastersColors)

		for colorName, images := range mastersColors {

			for _, image := range images.([]interface{}) {
				for key, value := range image.(map[string]interface{}) {
					println(key + value.(string))
				}

				colors = append(colors, ColorAsset{
					name: colorName,
					hex:  "getHex(name)",
					path: "getPath(name)",
				})
			}
		}
	}
	return colors
}

func GetAssetsData() assets.Assets {
	data, err := ioutil.ReadFile(root)
	assets := assets.Assets{}
	err = json.Unmarshal(data, &assets)
	if err != nil {
		log.Fatalln(err)
	}
	return assets
}
