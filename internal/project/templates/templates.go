package templates

import _ "embed"

var (
	//go:embed main.go.tmpl
	MainGo string

	//go:embed go.mod.tmpl
	GoMod string
)
