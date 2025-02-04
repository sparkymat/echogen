package project

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sparkymat/echogen/internal/project/templates"
)

var ErrDirectoryNotEmpty = errors.New("directory is not empty")

func (p *Project) Init(ctx context.Context, forceCreate bool) error {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Error().Err(err).Msg("failed to read directory")
		return err
	}

	if len(entries) > 0 && !forceCreate {
		log.Error().Err(ErrDirectoryNotEmpty).Msg("directory is not empty")
		return ErrDirectoryNotEmpty
	}

	// Create main.go
	fp, err := os.Create("main.go")
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
