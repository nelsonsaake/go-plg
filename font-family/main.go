package main

import (
	_ "embed"
	"encoding/json"
	"go-plg/render"
	"strings"

	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/pretty"
	"github.com/nelsonsaake/go/str"
	"github.com/nelsonsaake/go/ufs"
)

var (
	templateDir = "resources/template"
	outputDir   = "output"

	//go:embed resources/data/fontfamilies.json
	rawString string
)

func main() {

	type Data struct {
		Name  string
		Value string
	}

	var (
		data  = []Data{}
		fonts = []string{}
	)

	err := json.Unmarshal([]byte(rawString), &fonts)
	do.Die(err)

	for _, v := range fonts {

		// value
		value := v

		// name
		name := strings.ReplaceAll(v, " ", "")
		name = str.Capitalize(name)

		data = append(data, Data{Name: name, Value: value})
	}

	render.TemplatesInDir(templateDir, outputDir, data)

	ufs.WriteFile("ff.json", pretty.JSON(data))
}
