package config

type Config struct {
}

func Load() (*Config, error) {
	return &Config{}, nil
}
