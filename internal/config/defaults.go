package config

func LoadDefaultConfig() Rules {
	return Rules{
		Age: struct {
			Low    int `yaml:"low"`
			Medium int `yaml:"medium"`
		}{
			Low:    90,
			Medium: 180,
		},
		Size: struct {
			Low    int64 `yaml:"low"`
			Medium int64 `yaml:"medium"`
		}{
			Low:    100 * 1024 * 1024,
			Medium: 1024 * 1024 * 1024,
		},
	}
}
