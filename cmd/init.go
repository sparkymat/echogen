/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/sparkymat/echogen/internal/project"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init command bootstraps a new Echo-based web app in the current folder",
	Long:  `init command bootstraps a new Echo-based web app in the current folder. It errors out if the folder is not empty. `,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read name")
		}

		forceFlag, err := cmd.Flags().GetBool("force")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read force flag")
		}

		log.Info().Str("name", name).Msg("initializing project")

		p := project.New(name)

		if err := p.Init(cmd.Context(), forceFlag); err != nil {
			log.Fatal().Msg("failed to initialize project")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read cwd")
	}

	folderName := filepath.Base(wd)

	initCmd.Flags().BoolP("force", "f", false, "Force initialize even if the directory is not empty")
	initCmd.Flags().StringP("name", "n", folderName, "Name of the project")
}
