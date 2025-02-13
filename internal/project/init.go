package project

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sparkymat/echogen/internal/project/templates"
)

var ErrDirectoryNotEmpty = errors.New("directory is not empty")

func (p *Project) Init(ctx context.Context, path string, forceCreate bool) error {
	if _, err := os.Stat(path); err != nil && errors.Is(err, os.ErrNotExist) && forceCreate {
		if err = os.MkdirAll(path, 0755); err != nil {
			log.Error().Err(err).Msg("failed to create directory")
			return err
		}
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Error().Err(err).Msg("failed to read directory")
		return err
	}

	if len(entries) > 0 && !forceCreate {
		log.Error().Err(ErrDirectoryNotEmpty).Msg("directory is not empty")
		return ErrDirectoryNotEmpty
	}

	values := map[string]string{
		"Project":    p.Name,
		"ProjectURL": p.URL + "/" + p.Name,
		"Timestamp":  time.Now().Format("20060102150405"),
	}

	// Create Makefile
	if err := p.renderTemplateToFile(
		"makefile",
		templates.Makefile,
		path,
		"",
		"Makefile",
		values,
	); err != nil {
		return err
	}

	// Create main.go
	if err := p.renderTemplateToFile(
		"maingo",
		templates.MainGo,
		path,
		"",
		"main.go",
		values,
	); err != nil {
		return err
	}

	// Create .golangci.yml
	if err := p.renderTemplateToFile(
		"golangciyml",
		templates.GolangciYml,
		path,
		"",
		".golangci.yml",
		values,
	); err != nil {
		return err
	}

	// Create go.mod
	if err := p.renderTemplateToFile(
		"gomod",
		templates.GoMod,
		path,
		"",
		"go.mod",
		values,
	); err != nil {
		return err
	}

	// Create sqlc.yaml
	if err := p.renderTemplateToFile(
		"sqlcyaml",
		templates.SqlcYaml,
		path,
		filepath.Join(""),
		"sqlc.yaml",
		values,
	); err != nil {
		return err
	}

	// Create internal/config/service.go
	if err := p.renderTemplateToFile(
		"configservicego",
		templates.ConfigServiceGo,
		path,
		filepath.Join("internal", "config"),
		"service.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/database/service.go
	if err := p.renderTemplateToFile(
		"databaseservicego",
		templates.DatabaseServiceGo,
		path,
		filepath.Join("internal", "database"),
		"service.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/route/setup.go
	if err := p.renderTemplateToFile(
		"routesetupgo",
		templates.RouteSetupGo,
		path,
		filepath.Join("internal", "route"),
		"setup.go",
		values,
	); err != nil {
		return err
	}

	// Create internal//service.go
	if err := p.renderTemplateToFile(
		"routewebgo",
		templates.RouteWebGo,
		path,
		filepath.Join("internal", "route"),
		"web.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/route/api.go
	if err := p.renderTemplateToFile(
		"routeapigo",
		templates.RouteAPIGo,
		path,
		filepath.Join("internal", "route"),
		"api.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/handler/common.go
	if err := p.renderTemplateToFile(
		"handlercommongo",
		templates.HandlerCommonGo,
		path,
		filepath.Join("internal", "handler"),
		"common.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/handler/register.go
	if err := p.renderTemplateToFile(
		"handlerregistergo",
		templates.HandlerRegisterGo,
		path,
		filepath.Join("internal", "handler"),
		"register.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/handler/login.go
	if err := p.renderTemplateToFile(
		"handlerlogingo",
		templates.HandlerLoginGo,
		path,
		filepath.Join("internal", "handler"),
		"login.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/handler/home.go
	if err := p.renderTemplateToFile(
		"handlerhomego",
		templates.HandlerHomeGo,
		path,
		filepath.Join("internal", "handler"),
		"home.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/view/layout.templ
	if err := p.renderTemplateToFile(
		"viewlayouttempl",
		templates.ViewLayoutTempl,
		path,
		filepath.Join("internal", "view"),
		"layout.templ",
		values,
	); err != nil {
		return err
	}

	// Create internal/view/register.templ
	if err := p.renderTemplateToFile(
		"viewregistertempl",
		templates.ViewRegisterTempl,
		path,
		filepath.Join("internal", "view"),
		"register.templ",
		values,
	); err != nil {
		return err
	}

	// Create internal/view/login.templ
	if err := p.renderTemplateToFile(
		"viewlogintempl",
		templates.ViewLoginTempl,
		path,
		filepath.Join("internal", "view"),
		"login.templ",
		values,
	); err != nil {
		return err
	}

	// Create internal/view/home.templ
	if err := p.renderTemplateToFile(
		"viewhometempl",
		templates.ViewHomeTempl,
		path,
		filepath.Join("internal", "view"),
		"home.templ",
		values,
	); err != nil {
		return err
	}

	// Create internal/services.go
	if err := p.renderTemplateToFile(
		"servicesgo",
		templates.ServicesGo,
		path,
		filepath.Join("internal"),
		"services.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/user_service.go
	if err := p.renderTemplateToFile(
		"userservicego",
		templates.UserServiceGo,
		path,
		filepath.Join("internal"),
		"user_service.go",
		values,
	); err != nil {
		return err
	}

	// Create migrations/timestamp_create_users_table.up.sql
	if err := p.renderTemplateToFile(
		"createuserstableupsql",
		templates.CreateUsersTableUpSQL,
		path,
		filepath.Join("migrations"),
		values["Timestamp"]+"_create_users_table.up.sql",
		values,
	); err != nil {
		return err
	}

	// Create migrations/timestamp_create_users_table.down.sql
	if err := p.renderTemplateToFile(
		"createuserstabledownsql",
		templates.CreateUsersTableDownSQL,
		path,
		filepath.Join("migrations"),
		values["Timestamp"]+"_create_users_table.down.sql",
		values,
	); err != nil {
		return err
	}

	// Create internal/database/queries.sql
	if err := p.renderTemplateToFile(
		"databasequeriessql",
		templates.DatabaseQueriesSQL,
		path,
		filepath.Join("internal", "database"),
		"queries.sql",
		values,
	); err != nil {
		return err
	}

	return nil
}

func (p *Project) renderTemplateToFile(
	templateName string,
	templateString string,
	path string,
	folder string,
	filename string,
	values map[string]string,
) error {
	fullFolderPath := filepath.Join(path, folder)
	fullFilePath := filepath.Join(fullFolderPath, filename)
	filePath := filepath.Join(folder, filename)

	if err := os.MkdirAll(fullFolderPath, 0755); err != nil {
		log.Error().Err(err).Msgf("failed to render %s", filePath)
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpl, err := template.New(templateName).Parse(templateString)
	if err != nil {
		log.Error().Err(err).Msgf("failed to render %s", filePath)
		return fmt.Errorf("failed to parse template: %w", err)
	}

	fp, err := os.Create(fullFilePath)
	if err != nil {
		log.Error().Err(err).Msgf("failed to render %s", filePath)
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer fp.Close()

	if err = tmpl.Execute(fp, values); err != nil {
		log.Error().Err(err).Msgf("failed to render %s", filePath)
		return fmt.Errorf("failed to execute template: %w", err)
	}

	log.Info().
		Str("name", filePath).
		Msg("generated file")

	return nil
}
