package templates

import _ "embed"

var (
	//go:embed main.go.tmpl
	MainGo string

	//go:embed go.mod.tmpl
	GoMod string

	//go:embed internal_config_service.go.tmpl
	ConfigServiceGo string
)
