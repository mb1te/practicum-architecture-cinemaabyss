package internal

import "github.com/caarlos0/env/v11"

type Config struct {
	Port         string   `env:"PORT"`
	KafkaBrokers []string `env:"KAFKA_BROKERS"`
	KafkaTopic   string   `env:"KAFKA_TOPIC"`
}

func NewConfigFromEnv() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return cfg, err
}
