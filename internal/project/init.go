package project

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sparkymat/echogen/internal/project/templates"
)

var ErrDirectoryNotEmpty = errors.New("directory is not empty")

var projectTemplate = map[string]any{
	".echogen.json": templates.EchogenJson,
	".gitignore":    templates.Gitignore,
	".env.example":  templates.EnvExample,
	"Makefile":      templates.Makefile,
	"main.go":       templates.MainGo,
	"go.mod":        templates.GoMod,
	"sqlc.yaml":     templates.SqlcYaml,
	"internal": map[string]any{
		"services.go":     templates.ServicesGo,
		"user_service.go": templates.UserServiceGo,
		"auth": map[string]any{
			"jwt.go":                   templates.AuthJWTGo,
			"middleware.go":            templates.AuthMiddlewareGo,
			"proxy_auth_middleware.go": templates.AuthProxyAuthMiddlewareGo,
			"session.go":               templates.AuthSessionGo,
		},
		"config": map[string]any{
			"service.go": templates.ConfigServiceGo,
		},
		"database": map[string]any{
			"queries.sql": templates.DatabaseQueriesSQL,
			"service.go":  templates.DatabaseServiceGo,
		},
		"handler": map[string]any{
			"common.go":   templates.HandlerCommonGo,
			"home.go":     templates.HandlerHomeGo,
			"login.go":    templates.HandlerLoginGo,
			"register.go": templates.HandlerRegisterGo,
		},
		"route": map[string]any{
			"api.go":   templates.RouteAPIGo,
			"setup.go": templates.RouteSetupGo,
			"web.go":   templates.RouteWebGo,
		},
		"service": map[string]any{
			"user": map[string]any{
				"service.go": templates.ServiceUserServiceGo,
			},
		},
		"view": map[string]any{
			"layout.templ":   templates.ViewLayoutTempl,
			"login.templ":    templates.ViewLoginTempl,
			"home.templ":     templates.ViewHomeTempl,
			"register.templ": templates.ViewRegisterTempl,
		},
	},
	"migrations": map[string]any{
		"{{.Timestamp}}_create_users_table.up.sql":   templates.CreateUsersTableUpSQL,
		"{{.Timestamp}}_create_users_table.down.sql": templates.CreateUsersTableDownSQL,
	},
}

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

	err = p.writeMapToFS(path, []string{}, projectTemplate, values)
	if err != nil {
		log.Error().Err(err).Msg("failed to write map to FS")
		return err
	}

	return nil
}

func (p *Project) writeMapToFS(basePath string, pathChain []string, fileMap map[string]any, values map[string]string) error {
	for currentPath, value := range fileMap {
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.String:
			folder := filepath.Join(pathChain...)
			filePath := filepath.Join(folder, currentPath)
			err := p.renderTemplateToFile(filePath, v.String(), basePath, folder, currentPath, values)
			if err != nil {
				return fmt.Errorf("failed to render template to file: %w", err)
			}

		case reflect.Map:
			newPathChain := append(pathChain, currentPath)
			innerMap := value.(map[string]any)
			err := p.writeMapToFS(basePath, newPathChain, innerMap, values)
			if err != nil {
				return fmt.Errorf("failed to write map to FS: %w", err)
			}
		}
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

	actualFilename, err := renderTemplateToString(filename, values)
	if err != nil {
		log.Error().Err(err).Msgf("failed to render %s", filename)
	}

	fullFilePath := filepath.Join(fullFolderPath, actualFilename)
	filePath := filepath.Join(folder, actualFilename)

	if err = os.MkdirAll(fullFolderPath, 0755); err != nil {
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

func renderTemplateToString(in string, values map[string]string) (string, error) {
	tmpl, err := template.New(in).Parse(in)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, values); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
