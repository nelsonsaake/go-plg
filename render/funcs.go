package render

import (
	"strings"
	"text/template"

	"github.com/nelsonsaake/go/str"
)

var funcs = template.FuncMap{
	"ReplaceAll": strings.ReplaceAll,
	"ToUpper":    strings.ToUpper,
	"ToLower":    strings.ToLower,
	"Capitalize": str.Capitalize,
}
