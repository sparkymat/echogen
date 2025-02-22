package templates

import _ "embed"

var (
	//go:embed main.go.template
	MainGo string

	//go:embed go.mod.template
	GoMod string

	//go:embed sqlc.yaml.template
	SqlcYaml string

	//go:embed golangci.yml.template
	GolangciYml string

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

	//go:embed internal_view_layout.templ.template
	ViewLayoutTempl string

	//go:embed internal_view_register.templ.template
	ViewRegisterTempl string

	//go:embed internal_view_login.templ.template
	ViewLoginTempl string

	//go:embed internal_view_home.templ.template
	ViewHomeTempl string

	//go:embed internal_services.go.template
	ServicesGo string

	//go:embed internal_user_service.go.template
	UserServiceGo string

	//go:embed makefile.template
	Makefile string

	//go:embed create_users_table.up.sql.template
	CreateUsersTableUpSQL string

	//go:embed create_users_table.down.sql.template
	CreateUsersTableDownSQL string

	//go:embed internal_database_queries.sql.template
	DatabaseQueriesSQL string

	//go:embed internal_auth_middleware.go.template
	AuthMiddlewareGo string

	//go:embed internal_auth_session.go.template
	AuthSessionGo string

	//go:embed internal_auth_jwt.go.template
	AuthJWTGo string

	//go:embed internal_auth_proxy_auth_middleware.go.template
	AuthProxyAuthMiddlewareGo string

	//go:embed env_example.template
	EnvExample string

	//go:embed gitignore.template
	Gitignore string
)
