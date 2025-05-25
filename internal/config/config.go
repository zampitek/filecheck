package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Rules defines the syntax the YAML files will follow.
type Rules struct {
	Age struct {
		Low    int `yaml:"low"`
		Medium int `yaml:"medium"`
	} `yaml:"age"`

	Size struct {
		Low    int64 `yaml:"low"`
		Medium int64 `yaml:"medium"`
	} `yaml:"size"`
}

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
