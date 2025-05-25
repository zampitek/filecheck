package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/checks"
	"github.com/zampitek/filecheck/internal/config"
	"github.com/zampitek/filecheck/internal/report"

	"github.com/spf13/cobra"
)

// requireCheckForFlag ensures that a dependent flag is only used if its associated check is enabled.
// For example, if the "age" check is not present in --checks but --age-top is set, it returns an error.
func requireCheckForFlag(cmd *cobra.Command, check string, flags ...string) error {
	checkStr, _ := cmd.Flags().GetString("checks")
	checks := strings.Split(checkStr, ",")
	found := false

	for _, c := range checks {
		if strings.TrimSpace(c) == check {
			found = true
			break
		}
	}

	if !found {
		for _, flag := range flags {
			if cmd.Flags().Changed(flag) {
				return fmt.Errorf("flag --%s can only be used with the '%s' check", flag, check)
			}
		}
	}

	return nil
}

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for file issues",
	Args:  cobra.ExactArgs(1),

	// Executes the scan command by collecting file metadata and applying selected checks.
	Run: func(cmd *cobra.Command, args []string) {

		// Parse and normalize enabled checks
		checkStr, _ := cmd.Flags().GetString("checks")
		checksVars := strings.Split(checkStr, ",")
		checkSet := make(map[string]bool)
		for _, check := range checksVars {
			checkSet[strings.TrimSpace(check)] = true
		}

		ageTop, _ := cmd.Flags().GetInt("age-top")
		sizeTop, _ := cmd.Flags().GetInt("size-top")
		rulesPath, _ := cmd.Flags().GetString("rules")

		var rules config.Rules
		if rulesPath != "" {
			loadedRules, err := config.LoadConfig(rulesPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			rules = *loadedRules
		} else {
			rules = config.LoadDefaultConfig()
		}

		if err := requireCheckForFlag(cmd, "age", "age-top"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := requireCheckForFlag(cmd, "size", "size-top"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		files, _ := internal.Scan(args[0])
		var reportResult string

		reportResult += report.Header()

		if checkSet["age"] {
			lowAge, mediumAge, highAge := checks.CheckAge(files, rules)
			reportResult += report.AgeReport(lowAge, mediumAge, highAge, ageTop, rules)
		}
		if checkSet["size"] {
			lowSize, mediumSize, highSize := checks.CheckSize(files, rules)
			reportResult += report.SizeReport(lowSize, mediumSize, highSize, sizeTop, rules)
		}
		if checkSet["empty"] {
			emptyFiles := checks.GetEmptyFiles(files)
			reportResult += report.EmptyFilesReport(emptyFiles)
		}

		fmt.Print(reportResult)
	},
}

func init() {
	scanCmd.Flags().String("checks", "", "Comma-separated list of checks to run (e.g. age,size)")
	scanCmd.Flags().Int("age-top", 0, "Show top N files per age group (only used with 'age' check)")
	scanCmd.Flags().Int("size-top", 0, "Show top N files per size group (only used with 'size' check)")
	scanCmd.Flags().String("rules", "", "Path to a YAML file with custom rules")
}
