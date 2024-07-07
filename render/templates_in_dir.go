package render

import (
	"path/filepath"
	"strings"

	"github.com/nelsonsaake/go/dir"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ufs"
)

// TemplatesInDir: for each template in `indir`, combine with data and write the results to a file in `outdir`
// it will pass each file in the `indir` through the `Template` func and write the results to an output file
// the entry template for each file is `main`
// output filename is the input file without the prefix `.tmpl` if it ends with `.tmpl`
func TemplatesInDir(indir, outdir string, data any) {

	// tmplDir contains different tempaltes, for different outputs
	tmplDir, err := dir.New(indir)
	do.Die(err)

	// load all those templates
	templates, err := tmplDir.Files()
	do.Die(err)

	// build output for each template and write it out
	for _, template := range templates {

		// generate output
		output, err := Template(template, data)
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
