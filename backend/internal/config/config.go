package config

import (
	"fmt"
	"path"

	"github.com/caarlos0/env/v10"
	"github.com/jak103/powerplay/internal/utils/constants"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/joho/godotenv"
)

type Config struct {
	Env       string `env:"POWERPLAY_ENV"`
	Dir       string `env:"POWERPLAY_CONFIG_DIR" envDefault:"/app/config"`
	DebugVars bool   `env:"POWERPLAY_DEBUG_VARS" envDefault:"false"`
	LogLevel  string `env:"POWERPLAY_LOG_LEVEL" envDefault:"INFO"`
	ColorLog  bool   `env:"POWERPLAY_COLOR_LOG" envDefault:"true"`
	Db        Postgres
}

type Postgres struct {
	Host     string `env:"DB_HOST" envDefault:"database"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Username string `env:"DB_USERNAME" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	DbName   string `env:"DB_NAME" envDefault:"powerplay"`
}

var Vars Config

func Init() error {
	envConfig := struct {
		Env string `env:"POWERPLAY_ENV"`
		Dir string `env:"POWERPLAY_CONFIG_DIR"`
	}{}

	if err := env.Parse(&envConfig); err != nil {
		log.WithErr(err).Alert("Failed to get env")
		return err
	}

	if envConfig.Env == constants.Local || envConfig.Env == constants.Test {
		path := path.Join(envConfig.Dir, fmt.Sprintf("%s.env", envConfig.Env))
		err := godotenv.Load(path)
		if err != nil {
			log.WithErr(err).Alert("Failed to load %s", path)
			return err
		}
	}

	if err := env.Parse(&Vars); err != nil {
		log.WithErr(err).Alert("Failed to parse env vars")
		return err
	}

	if Vars.DebugVars {
		log.Debug("Env vars: %v", Vars)
	}

	return nil
}
