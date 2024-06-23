package render

import (
	"strings"
	"text/template"
)

var funcs = template.FuncMap{
	"ReplaceAll": strings.ReplaceAll,
	"ToUpper":    strings.ToUpper,
	"ToLower":    strings.ToLower,
	"Capitalize": capitalize,
}

func capitalize(v string) string {

	if len(v) < 2 {
		return v
	}

	ls := strings.Split(v, "")
	return strings.ToUpper(ls[0]) + strings.Join(ls[1:], "")
}
