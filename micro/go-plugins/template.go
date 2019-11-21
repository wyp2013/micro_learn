package plugin

var (
	tmpl = `
package main

import (
	"{{.Path}}"
	"micro_learn/micro/go-plugins"
)

var Plugin = plugin.Plugin{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
