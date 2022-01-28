package main

import (
	"os"
	"strings"
	"text/template"
)

const templateText = `
Output 0 : {{title .Name1}}
Output 1 : {{title .Name2}}
Output 3 : {{.Name3 | title}}
`
func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go-programing-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programing",
		"Name3": "tour",
	}

	_ = tpl.Execute(os.Stdout, data)
}