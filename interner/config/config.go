package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" yaml:"redis"`
	Jwt   Jwt   `mapstructure:"jwt" yaml:"jwt"`
}
