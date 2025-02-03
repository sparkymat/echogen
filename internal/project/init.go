package project

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
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

	return nil
}
