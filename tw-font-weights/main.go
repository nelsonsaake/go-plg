package main

import (
	_ "embed"
	"encoding/json"
	"go-plg/render"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nelsonsaake/go/dir"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ufs"
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

// writeOut: for each template in `indir`, combine with data and write the results to a file in `outdir`
func writeOut(indir, outdir string, data []FontSize) {

	// tmplDir contains different tempaltes, for different outputs
	tmplDir, err := dir.New(indir)
	do.Die(err)

	// load all those templates
	templates, err := tmplDir.Files()
	do.Die(err)

	// build output for each template and write it out
	for _, template := range templates {

		// generate output
		output, err := render.Main(template, data)
		do.Die(err)

		// generate an output file base on the template file
		outputfile := filepath.Base(template)
		outputfile = strings.TrimSuffix(outputfile, ".tmpl")
		outputfile = filepath.Join(outdir, outputfile)

		// delete old output file
		// we don't want any situation of old data, sort of overlapping the new one
		err = ufs.DelFile(outputfile)
		do.Die(err)

		// write new output file
		err = ufs.WriteFile(outputfile, output)
		do.Die(err)
	}
}

func main() {

	var (
		fontsizes = getFontSizes()
	)

	writeOut(templateDirPath, outputDirPath, fontsizes)
}
