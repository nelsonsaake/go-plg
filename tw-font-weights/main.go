package main

import (
	_ "embed"
	"encoding/json"
	"go-plg/render"
	"sort"

	"github.com/nelsonsaake/go/do"
)

var (
	//go:embed data/data
	rawData string

	// template dir path
	templateDirPath = "./templates"

	// output dir prth
	outputDirPath = "./output"
)

// FontSize eg. bold: 800
type FontSize struct {
	Name  string
	Value string
}

// getFontSizes: breakdown `rawData` and convert the data into `[]FontSize`
func getFontSizes() []FontSize {

	var (
		data    = []byte(rawData)
		dataMap = map[string]string{}
		res     = []FontSize{}
	)

	err := json.Unmarshal(data, &dataMap)
	do.Die(err)

	for name, value := range dataMap {
		res = append(res, FontSize{Name: name, Value: value})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})

	return res
}

func main() {

	var (
		fontsizes = getFontSizes()
	)

	render.TemplatesInDir(templateDirPath, outputDirPath, fontsizes)
}
