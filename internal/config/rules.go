package config

// TODO: implement the logic to manage all this
// Rules defines the syntax the YAML files will follow.
type Rules struct {
	Age  *AgeRules  `yaml:"age"`
	Size *SizeRules `yaml:"size"`
	Rule *RuleSpec  `yaml:"rule"`
}

type AgeRules struct {
	Low    int `yaml:"low"`
	Medium int `yaml:"medium"`
}

type SizeRules struct {
	Low    int64 `yaml:"low"`
	Medium int64 `yaml:"medium"`
}

type RuleSpec struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Filters		 *Filters `yaml:"filters"`
	Action       string `yaml:"action"`
	Confirmation bool   `yaml:"confirmation"`
}

type Filters struct {
	Age          int `yaml:"age"`
	AgeCategory  int `yaml:"age_category"`
	Size         int `yaml:"size"`
	SizeCategory int `yaml:"size_category"`
}
