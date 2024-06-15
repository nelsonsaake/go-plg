package main

import (
	_ "embed"
	"fmt"
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
	ls := do.Lines(data)

	var (
		name      string
		shade     string
		colorCode string
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

		output += fmt.Sprintf("static final Color %s%s = strToColor(\"%s\");\n", strings.ToLower(name), shade, colorCode)
	}

	err := ufs.DelFile(outputfile)
	do.Die(err)

	err = ufs.WriteFile(outputfile, output)
	do.Die(err)
}
