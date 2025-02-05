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
		filepath.Join(path, "main.go"),
		values,
	); err != nil {
		log.Error().Err(err).Msg("failed to write to main.go")
		return err
	}
	log.Info().
		Str("name", "main.go").
		Msg("generated file")

	// Create go.mod
	if err := p.renderTemplateToFile(
		"gomod",
		templates.GoMod,
		filepath.Join(path, "go.mod"),
		values,
	); err != nil {
		log.Error().Err(err).Msg("failed to render go.mod")
		return err
	}
	log.Info().
		Str("name", "go.mod").
		Msg("generated file")

	return nil
}

func (p *Project) renderTemplateToFile(templateName string, templateString string, path string, values map[string]string) error {
	tmpl, err := template.New(templateName).Parse(templateString)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	fp, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer fp.Close()

	if err = tmpl.Execute(fp, values); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
