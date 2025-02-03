/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
			panic(err)
		}

		forceFlag, err := cmd.Flags().GetBool("force")
		if err != nil {
			panic(err)
		}

		fmt.Printf("initializing project with name=%s\n", name)

		p := project.New(name)

		if err := p.Init(cmd.Context(), forceFlag); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	folderName := filepath.Base(wd)

	initCmd.Flags().BoolP("force", "f", false, "Force initialize even if the directory is not empty")
	initCmd.Flags().StringP("name", "n", folderName, "Name of the project")
}
