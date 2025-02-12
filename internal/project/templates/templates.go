package templates

import _ "embed"

var (
	//go:embed main.go.tmpl
	MainGo string

	//go:embed go.mod.tmpl
	GoMod string

	//go:embed internal_config_service.go.tmpl
	ConfigServiceGo string

	//go:embed internal_database_service.go.tmpl
	DatabaseServiceGo string

	//go:embed internal_route_setup.go.tmpl
	RouteSetupGo string

	//go:embed internal_route_web.go.tmpl
	RouteWebGo string

	//go:embed internal_route_api.go.tmpl
	RouteAPIGo string
)
