package cmd

import (
	"filecheck/internal/report"
	"filecheck/internal/scanner"
	"fmt"

	"github.com/spf13/cobra"
)

var extendedReport bool

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for file issues",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		scanned, err := scanner.Scan(path)
		if err != nil {
			fmt.Printf("Scan failed: %v\n", err)
			return
		}

		fmt.Printf("Scanned %d files/directories\n", len(scanned.Files))

		var reportMessage string

		if extendedReport {
			reported := report.CreateExtendedReport(scanned)
			reportMessage = report.PrintExtendedReport(reported)
		} else {
			reported := report.CreateReport(scanned)
			reportMessage = report.PrintReport(reported)
		}

		fmt.Print(reportMessage)
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&extendedReport, "extended", "e", false, "Print an extended version of the report")
	rootCmd.AddCommand(scanCmd)
}
