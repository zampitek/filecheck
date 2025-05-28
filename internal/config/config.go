package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadConfig returns a Rules struct with the data parsed from
// the provided YAML file.
func LoadConfig(path string) (*Rules, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rules Rules
	if err := yaml.Unmarshal(file, &rules); err != nil {
		return nil, err
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
