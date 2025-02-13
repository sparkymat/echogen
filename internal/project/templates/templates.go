package templates

import _ "embed"

var (
	//go:embed main.go.template
	MainGo string

	//go:embed go.mod.template
	GoMod string

	//go:embed internal_config_service.go.template
	ConfigServiceGo string

	//go:embed internal_database_service.go.template
	DatabaseServiceGo string

	//go:embed internal_route_setup.go.template
	RouteSetupGo string

	//go:embed internal_route_web.go.template
	RouteWebGo string

	//go:embed internal_route_api.go.template
	RouteAPIGo string

	//go:embed internal_handler_common.go.template
	HandlerCommonGo string

	//go:embed internal_handler_register.go.template
	HandlerRegisterGo string

	//go:embed internal_handler_login.go.template
	HandlerLoginGo string

	//go:embed internal_handler_home.go.template
	HandlerHomeGo string
)
