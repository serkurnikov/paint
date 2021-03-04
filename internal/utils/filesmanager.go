package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Info struct {
	Name    string
	Path    string
	IsDir   bool
	ModTime time.Time
	Size    int64
}

func GetListingDirectoryInfo(root string) ([]Info, error) {
	information := make([]Info, 0)
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			information = append(information, Info{
				Name:    info.Name(),
				Path:    path,
				IsDir:   info.IsDir(),
				ModTime: info.ModTime(),
				Size:    info.Size(),
			})
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return information, err
}

func WriteFile(data []byte, path string) {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ParsePath(info Info) []string {
	result := strings.Split(info.Path, "\\")
	return result
}
