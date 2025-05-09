/*
Copyright © 2025 RICCARDO ZAMPIERI riccardo.zampieri28@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filecheck",
	Short: "A program to help you organize your files and folders",
	Long: `Filecheck is a free, open-source program that helps you keeping your file system organized. 
It creates a report of all the problems a given directory has.
Its functionalities can be extended through rules.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(versionCmd)
}
