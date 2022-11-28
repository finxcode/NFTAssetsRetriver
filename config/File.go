package config

type File struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}
