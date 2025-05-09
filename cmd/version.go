package cmd

import (
	"filecheck/version"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of filecheck",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("filecheck version %s\n", version.Version)

		if version.Commit != "" {
			fmt.Printf("Commit: %s\n", version.Commit)
		}
		if version.BuildDate != "" {
			fmt.Printf("Built at: %s\n", version.BuildDate)
		}
	},
}
