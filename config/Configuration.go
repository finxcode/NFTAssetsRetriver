package config

type Configuration struct {
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	File     File     `mapstructure:"file" json:"file" yaml:"file"`
	Excel    Excel    `mapstructure:"excel" json:"excel" yaml:"excel"`
}
