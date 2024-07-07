package main

import (
	_ "embed"
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
)

func main() {

	type Data struct {
		Name  string
		Value string
	}

	data := []Data{}

	ff, err := getGoogleFonts()
	do.Die(err)

	err = ufs.WriteFile("resources/data/fontfamilies.json", pretty.JSON(ff))
	do.Die(err)

	for _, v := range ff {

		// value
		value := v

		// name
		name := strings.ReplaceAll(v, " ", "")
		name = str.Capitalize(name)

		data = append(data, Data{Name: name, Value: value})
	}

	render.TemplatesInDir(templateDir, outputDir, data)
}
