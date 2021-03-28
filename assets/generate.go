package assets

//go:generate gojson -name=assets -input structure.json -o structure.go -pkg assets

import (
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
	"log"
	"paint/pkg/utils"
	"strings"
)

type File struct {
	Name string
	Path string
}

func Init() {
	jsonObj := gabs.New()
	information, _ := utils.GetListingDirectoryInfo("assets")

	for _, info := range information {
		path := utils.ParsePath(info)

		if !info.IsDir {
			hierarchy := strings.Join(path[:len(path)-1], ".")

			jsonObj.ArrayConcatP(File{
				Name: info.Name,
				Path: info.Path,
			}, hierarchy)
		}
	}

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonObj.String()), &jsonMap)

	d, err := json.Marshal(jsonMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	utils.WriteFile(d, "assets/structure.json")
	println(jsonObj.String())
}
