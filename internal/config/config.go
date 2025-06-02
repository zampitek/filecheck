package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/zampitek/filecheck/internal/err"
)

// LoadConfig returns a Rules struct with the data parsed from
// the provided YAML file.
func LoadConfig(path string) (*Rules, error) {
	file, e := os.ReadFile(path)
	if e != nil {
		return nil, err.Wrap("opening config file", e)
	}

	var rules Rules
	if e := yaml.Unmarshal(file, &rules); e != nil {
		return nil, err.Wrap("parsing config file", e)
	}

	return &rules, nil
}

// LoadDefaultConfig returns a Rules struct containing the default rules:
//   - 90 and 180 days for age categorization
//   - 100 MB and 1 GB for size categorization
func LoadDefaultConfig() Rules {
	return Rules{
		Age: &AgeRules{
			Low:    90,
			Medium: 180,
		},
		Size: &SizeRules{
			Low:    100 * 1024 * 1024,
			Medium: 1024 * 1024 * 1024,
		},
	}
}
