package project

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

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

	// Create internal/config/service.go
	if err := p.renderTemplateToFile(
		"gomod",
		templates.ConfigServiceGo,
		path,
		filepath.Join("internal", "config"),
		"service.go",
		values,
	); err != nil {
		return err
	}

	// Create internal/config/service.go
	if err := p.renderTemplateToFile(
		"gomod",
		templates.ConfigServiceGo,
		path,
		filepath.Join("internal", "database"),
		"service.go",
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
