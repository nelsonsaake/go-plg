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

	//go:embed resources/data/data
	raw string

	// template file
	tempatefile = "resources/template/asset_illustrations_values.dart.tmpl"

	// output file
	outputfile = "output/asset_illustrations_values.dart"
)

type Data struct {
	Name  string
	Value string
}

func main() {

	var (
		lines = do.Lines(raw)
		data  = []Data{}
	)

	for _, value := range lines {

		name := value

		// remove some junk
		name = strings.TrimPrefix(name, "undraw_")

		// remove some more junk
		name = strings.TrimSuffix(name, ".svg")

		// remove some more junk
		name = strings.ReplaceAll(name, "-", " ")

		// remove even more junk
		i := strings.LastIndex(name, "_re")
		if i != -1 {
			name = name[:i]
		}

		// convert pascal case to words
		name = strings.ReplaceAll(name, "_", " ")

		// convert words to title case
		name = str.Capitalize(name)

		// remove all spaces
		name = strings.ReplaceAll(name, " ", "")

		// compile data for output generation
		data = append(data, Data{name, value})
	}

	output, err := render.Template(tempatefile, data)
	do.Die(err)

	err = ufs.DelFile(outputfile)
	do.Die(err)

	err = ufs.WriteFile(outputfile, output)
	do.Die(err)
}
