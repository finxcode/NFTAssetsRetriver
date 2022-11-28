package config

type Excel struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}
