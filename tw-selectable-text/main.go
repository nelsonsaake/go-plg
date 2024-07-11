package main

import (
	"encoding/json"
	"fmt"
	"go-plg/render"
	"go-plg/tw-selectable-text/models"
	"path/filepath"
	"strings"

	"github.com/nelsonsaake/go/ufs"
)

func renderTemplateAndWriteToDir(templateFilepath string, data any, outputDir string) error {

	output, err := render.Template(templateFilepath, data)
	if err != nil {
		return fmt.Errorf("error executing data: %v", err)
	}

	// generate an output file base on the template file
	outputfile := filepath.Base(templateFilepath)
	outputfile = strings.TrimSuffix(outputfile, ".tmpl")
	outputfile = filepath.Join(outputDir, outputfile)

	// delete old output file
	// we don't want any situation of old data, sort of overlapping the new one
	err = ufs.DelFile(outputfile)
	if err != nil {
		return fmt.Errorf("error deleting old output file: %v", err)
	}

	// write new output file
	err = ufs.WriteFile(outputfile, output)
	if err != nil {
		return fmt.Errorf("error writing output to file: %v", err)
	}

	return nil
}

func main() {

	type Config struct {
		Tmpl     string
		DataFile string
		Data     any
	}

	var (
		tmplDir     = "resources/template"
		dataDir     = "resources/data"
		outputDir   = "output"
		configArray = []Config{
			{
				Tmpl:     "selectable_text_colors.dart.tmpl",
				DataFile: "tailwind_colors.json",
				Data:     &[]models.Color{},
			},
			{
				Tmpl:     "selectable_text_font_family.dart.tmpl",
				DataFile: "google_font_family.json",
				Data:     &[]string{},
			},
			{
				Tmpl:     "selectable_text_font_size.dart.tmpl",
				DataFile: "font_size.json",
				Data:     &[]models.FontSize{},
			},
			{
				Tmpl:     "selectable_text_font_weight.dart.tmpl",
				DataFile: "font_weight.json",
				Data:     &[]models.FontWeight{},
			},
			{
				Tmpl:     "selectable_text_letter_spacing.dart.tmpl",
				DataFile: "letter_spacing.json",
				Data:     &[]models.LetterSpacing{},
			},
			{
				Tmpl:     "selectable_text_line_height.dart.tmpl",
				DataFile: "line_height.json",
				Data:     &[]models.LineHeight{},
			},
		}
	)

	for _, config := range configArray {

		var (
			tmpl     = filepath.Join(tmplDir, config.Tmpl)
			datafile = filepath.Join(dataDir, config.DataFile)
		)

		raw, err := ufs.ReadFileAsBytes(datafile)
		if err != nil {
			panic(fmt.Errorf("error reading %s: %v", datafile, err))
		}

		err = json.Unmarshal(raw, config.Data)
		if err != nil {
			panic(fmt.Errorf("error unmarshalling bytes from %s to %t: %v", datafile, config.Data, err))
		}

		renderTemplateAndWriteToDir(tmpl, config.Data, outputDir)
	}
}
