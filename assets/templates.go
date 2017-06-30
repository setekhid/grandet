package assets

import (
	"bytes"
	"text/template"
)

var asset_go_tmpl, grandet_go_tmpl *assetTmpl

func getTemplate(name string) *assetTmpl {
	tmpl_str := Grandet.Asset(name)
	tmpl := template.New(name)
	tmpl = template.Must(tmpl.Parse(string(tmpl_str)))
	return (*assetTmpl)(tmpl)
}

type assetTmpl template.Template

func (t *assetTmpl) parse(data interface{}) []byte {
	buff := bytes.Buffer{}
	err := (*template.Template)(t).Execute(&buff, data)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}
