package app

import "github.com/kelseyhightower/envconfig"

func New() (cfg Config) {
	envconfig.MustProcess("disputes", &cfg)
	return cfg
}

type Config struct {
	Name    string
	Version string
	Host    string
}
