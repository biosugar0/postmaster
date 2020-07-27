package config

const defaultRefreshRate = 2

type P8r struct {
	RefreshRate    int    `yaml:"refreshRate"`
	ReadOnly       bool   `yaml:"readOnly"`
	CurrentProfile string `yaml:"currentContext"`
}

// NewP8r create a new P8r configuration.
func NewP8r() *P8r {
	return &P8r{
		RefreshRate: defaultRefreshRate,
	}
}
