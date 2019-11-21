package template

var (
	Plugin = `package main
{{if .Plugins}}
import ({{range .Plugins}}
	_ "micro_learn/micro/go-plugins/{{.}}"{{end}}
){{end}}
`
)
