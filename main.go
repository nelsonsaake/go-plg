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

	//output
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

		if str.IsInt(v) {
			shade = v
		} else if strings.HasPrefix(v, "#") {
			colorCode = v
		} else {
			name = v
		}

		if str.IsEmpty(name) || str.IsEmpty(shade) || str.IsEmpty(colorCode) {
			continue
		}

		output += fmt.Sprintf("const %s%s = %s;\n", strings.ToLower(name), shade, colorCode)
	}

	err := ufs.WriteFile("colors.dart", output)
	do.Die(err)
}
