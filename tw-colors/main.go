package main

import (
	_ "embed"
	"go-plg/render"
	"strings"

	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/str"
	"github.com/nelsonsaake/go/ufs"
)

var (
	//go:embed data
	data string

	// output file name
	outputfile string = "output"

	// output
	output string
)

func main() {

	type Color struct {
		Name  string
		Shade string
		Code  string
	}

	var (
		name      string
		shade     string
		colorCode string
	)

	var (
		// parse the content, i.e `data` of a file into line
		ls = do.Lines(data)

		// colors will be used as input to template
		colors = []Color{}
	)

	for _, v := range ls {

		if strings.HasPrefix(v, "#") {
			colorCode = v
		} else if str.IsInt(v) {
			shade = v
			continue
		} else {
			name = v
			output += "\n"
			continue
		}

		colors = append(colors, Color{Name: name, Shade: shade, Code: colorCode})
	}

	output, err := render.Main("templates/text_colors.dart", colors)
	do.Die(err)

	err = ufs.DelFile(outputfile)
	do.Die(err)

	err = ufs.WriteFile(outputfile, output)
	do.Die(err)
}
