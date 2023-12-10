package config

type Config struct {
	Env      string `env:"CONFIG_ENV"`
	ColorLog bool   `env:"CONFIG_COLOR_LOG" envDefault:"true"`
}

var App Config

func Init() error {
	// TODO Load config
	return nil
}

func Load() (*Config, error) {
	return &Config{}, nil
}
