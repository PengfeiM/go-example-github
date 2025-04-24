package texttemp

import "text/template"

var createTemp = func(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}
