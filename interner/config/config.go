package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
}
