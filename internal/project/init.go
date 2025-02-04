package project

import (
	"context"
	"errors"
	"io"
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

	// Create main.go
	fp, err := os.Create(filepath.Join(path, "main.go"))
	if err != nil {
		log.Error().Err(err).Msg("failed to create main.go")
		return err
	}
	defer fp.Close()

	if _, err = io.WriteString(fp, templates.MainGo); err != nil {
		log.Error().Err(err).Msg("failed to write to main.go")
		return err
	}

	// Create go.mod

	return nil
}
