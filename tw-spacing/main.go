package main

import (
	_ "embed"
	"encoding/json"
	"go-plg/render"
	"path/filepath"
	"strings"

	"github.com/nelsonsaake/go/dir"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/str"
	"github.com/nelsonsaake/go/ufs"
)

var (
	//go:embed data/default-scale
	defaultScale string

	//go:embed data/responsive-scale
	responsiveScale string

	// template dir path
	templateDirPath = "./templates"

	// output dir prth
	outputDirPath = "./output"

	// reponsive scale only template dir path
	responsiveScaleTemplateDirPath = "./templates/responsive_scale"

	// reponsive scale only output dir path
	responsiveScaleOutputDirPath = "./output/responsive_scale"
)

// Dimen name, css rem, in px
type Dimen struct {
	Name string
	Rem  string
	Px   string
}

// getDimensFromDefaultScale: breakdown `defaultScale` and convert the data into `[]Dimen`
func getDimensFromDefaultScale() []Dimen {

	var (
		res = []Dimen{}
	)

	// parse content the file `data` into lines
	ls := do.Lines(defaultScale)

	// process the data from the file, line by line
	for _, line := range ls {

		// eg. of line:  0.5	0.125rem	2px
		vs := strings.Split(line, "\t") // `\t` separated
		if len(vs) != 3 {
			continue
		}

		// extract data
		var name, rem, px = vs[0], vs[1], vs[2]

		// in programming identifiers can't `.`
		name = strings.ReplaceAll(name, ".5", "a")
		name = strings.TrimSuffix(name, "px")

		// continue
		if str.IsEmpty(name) {
			continue
		}

		// remove units from re
		rem = strings.TrimSuffix(rem, "rem")
		rem = strings.TrimSuffix(rem, "px")

		// remove px units
		px = strings.TrimSuffix(px, "px")

		// make space and add it to list
		space := Dimen{Name: name, Rem: rem, Px: px}
		res = append(res, space)
	}

	return res
}

// getDimensFromDefaultScale: breakdown `responsiveScale` and convert the data into `[]Dimen`
func getDimensFromResponsiveScale() []Dimen {

	var (
		data    = []byte(responsiveScale)
		dataMap = map[string]string{}
		res     = []Dimen{}
	)

	err := json.Unmarshal(data, &dataMap)
	do.Die(err)

	for name, px := range dataMap {

		// capitalize, so it ends up as pascal case
		name = strings.ToTitle(name)

		// clean px, so we only actually have a double value
		px = strings.TrimSuffix(px, "px")

		dimen := Dimen{Name: name, Px: px}

		res = append(res, dimen)
	}

	return res
}

// writeOut: for each template in `indir`, combine with data and write the results to a file in `outdir`
func writeOut(indir, outdir string, data []Dimen) {

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
		defaultDimens    = getDimensFromDefaultScale()
		responsiveDimens = getDimensFromResponsiveScale()
		dimens           = append(defaultDimens, responsiveDimens...)
	)

	writeOut(templateDirPath, outputDirPath, dimens)
	writeOut(responsiveScaleTemplateDirPath, responsiveScaleOutputDirPath, responsiveDimens)
}
