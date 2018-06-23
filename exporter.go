package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

var httpFunctions template.FuncMap

func init() {

	httpFunctions = template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"VERSION": func() string {
			return VERSION
		},
		"underscore": func(html string) string {
			return UnderScoreString(html)
		},
	}

}

func ExportIndexHTML() string {
	out := index{*core, services}
	nav, _ := tmplBox.String("nav.html")
	footer, _ := tmplBox.String("footer.html")
	render, err := tmplBox.String("index.html")
	if err != nil {
		panic(err)
	}
	t := template.New("message")
	t.Funcs(httpFunctions)
	t, _ = t.Parse(nav)
	t, _ = t.Parse(footer)
	t.Parse(render)

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, out); err != nil {
		log.Println(err)
	}
	result := tpl.String()

	fmt.Println(result)
	return result
}
