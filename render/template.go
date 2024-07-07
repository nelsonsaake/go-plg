package render

import (
	"bytes"
	"text/template"
)

// Template: renders the templatefile given the data. Uses 'main' template as the entry point.
func Template(templatefile string, data any) (string, error) {

	t := template.New("main")
	t = t.Funcs(funcs)

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
