package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/checks"
	"github.com/zampitek/filecheck/internal/config"
	"github.com/zampitek/filecheck/internal/report"
)

func init() {
	scanCmd.Flags().String("checks", "", "Comma-separated list of checks to run (e.g. age,size)")
	scanCmd.Flags().Int("age-top", 0, "Show top N files per age group (only used with 'age' check)")
	scanCmd.Flags().Int("size-top", 0, "Show top N files per size group (only used with 'size' check)")
	scanCmd.Flags().String("rules", "", "Path to a YAML file with custom rules")
}

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for file issues",
	Args:  cobra.ExactArgs(1),
	Run:   runScan,
}

func runScan(cmd *cobra.Command, args []string) {
	checkSet := parseChecks(cmd)
	rules := loadRules(cmd)

	validateFlagDependencies(cmd, checkSet)

	files, _ := internal.Scan(args[0])
	result := buildReport(files, checkSet, cmd, rules)

	fmt.Print(result)
}

// parseChecks parses and normalizes the --checks flag into a map.
func parseChecks(cmd *cobra.Command) map[string]bool {
	checkStr, _ := cmd.Flags().GetString("checks")
	checkSet := make(map[string]bool)
	for _, check := range strings.Split(checkStr, ",") {
		checkSet[strings.TrimSpace(check)] = true
	}
	return checkSet
}

// loadRules loads custom or default rules.
func loadRules(cmd *cobra.Command) config.Rules {
	rulesPath, _ := cmd.Flags().GetString("rules")
	if rulesPath == "" {
		return config.LoadDefaultConfig()
	}
	rules, err := config.LoadConfig(rulesPath)
	if err != nil {
		exitWithError(err)
	}
	if rules.Age == nil {
		rules.Age = config.LoadDefaultConfig().Age
	}
	if rules.Size == nil {
		rules.Size = config.LoadDefaultConfig().Size
	}

	return *rules
}

// validateFlagDependencies ensures flag-check consistency.
func validateFlagDependencies(cmd *cobra.Command, checkSet map[string]bool) {
	enforceCheckFlagDependency(cmd, checkSet, "age", "age-top")
	enforceCheckFlagDependency(cmd, checkSet, "size", "size-top")
}

// enforceCheckFlagDependency exits if a flag is used without its required check.
func enforceCheckFlagDependency(cmd *cobra.Command, checkSet map[string]bool, check string, flags ...string) {
	if checkSet[check] {
		return
	}
	for _, flag := range flags {
		if cmd.Flags().Changed(flag) {
			exitWithError(fmt.Errorf("flag --%s can only be used with the '%s' check", flag, check))
		}
	}
}

// buildReport generates the full report string.
func buildReport(files []internal.FileInfo, checkSet map[string]bool, cmd *cobra.Command, rules config.Rules) string {
	var output strings.Builder
	output.WriteString(report.Header())

	if checkSet["age"] {
		ageTop, _ := cmd.Flags().GetInt("age-top")
		low, mid, high := checks.CheckAge(files, rules)
		output.WriteString(report.AgeReport(low, mid, high, ageTop, rules))
	}
	if checkSet["size"] {
		sizeTop, _ := cmd.Flags().GetInt("size-top")
		low, mid, high := checks.CheckSize(files, rules)
		output.WriteString(report.SizeReport(low, mid, high, sizeTop, rules))
	}
	if checkSet["empty"] {
		empty := checks.GetEmptyFiles(files)
		output.WriteString(report.EmptyFilesReport(empty))
	}

	return output.String()
}

// exitWithError prints an error and exits the program.
func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
