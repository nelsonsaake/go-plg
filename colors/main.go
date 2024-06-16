package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

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

func toDart(name, shade, colorCode string) string {
	return fmt.Sprintf("static final Color %s%s = strToColor(\"%s\");\n", strings.ToLower(name), shade, colorCode)
}

func toExlize(name, shade, colorCode string) string {
	return fmt.Sprintf("static final Color %s%s = strToColor(\"%s\");\n", strings.ToLower(name), shade, colorCode)
}

func render(templatefile string, data any) (string, error) {

	t := template.New("main")

	t, err := t.ParseFiles(templatefile)
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}

	err = t.ExecuteTemplate(&buf, "main", data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

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

		// data will be used as input to template
		data = []Color{}
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

		data = append(data, Color{Name: name, Shade: shade, Code: colorCode})
	}

	output, err := render("xlscolors", data)
	do.Die(err)

	err = ufs.DelFile(outputfile)
	do.Die(err)

	err = ufs.WriteFile(outputfile, output)
	do.Die(err)
}
