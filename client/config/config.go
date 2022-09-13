package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RedisAddress string `envconfig:"REDIS_ADDRESS" default:""`
}

func Get() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
