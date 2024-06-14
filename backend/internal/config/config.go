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
	Env       string `env:"ENV" envDefault:"local"`
	Dir       string `env:"CONFIG_DIR" envDefault:"/powerplay/config"`
	DebugVars bool   `env:"DEBUG_VARS" envDefault:"false"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	LogColor  bool   `env:"LOG_COLOR" envDefault:"true"`
	JwtSecret string `env:"JWT_SECRET"`
	// VapidPublicKey  string   `env:"VAPID_PUBLIC_KEY"  envDefault:"BMPQhGq2KuP92WTzRK7S5UgLk5v8H0ZoNXXJji0J5wO3ufLm24AgelUfpe0BvasoupYfSagpGFZvwRTSBS-KYzY"`
	// VapidPrivateKey string   `env:"VAPID_PRIVATE_KEY" envDefault:"ZcXYJyrk0kAeC0VkIcJWkwlPvC6CwrVsjTlys1Uu2P8"`
	Port        string   `env:"PORT" envDefault:"8080"`
	Db          Postgres `envPrefix:"DB_"`
	PasswordKey string   `env:"PASSWORD_KEY,required"`
}

type Postgres struct {
	Host     string `env:"HOST" envDefault:"database"`
	Port     string `env:"PORT" envDefault:"5432"`
	Username string `env:"USERNAME" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"password"`
	DbName   string `env:"NAME" envDefault:"powerplay"`
}

var Vars Config

func Init() error {
	opts := env.Options{
		Prefix: "POWERPLAY_",
	}

	envConfig := struct {
		Env string `env:"ENV" envDefault:"local"`
		Dir string `env:"CONFIG_DIR" envDefault:"/powerplay/config"`
	}{}

	if err := env.ParseWithOptions(&envConfig, opts); err != nil {
		log.WithErr(err).Alert("Failed to get env")
		return err
	}

	if envConfig.Env == constants.Local || envConfig.Env == constants.Test {
		join := path.Join(envConfig.Dir, fmt.Sprintf("%s.env", envConfig.Env))
		log.Alert("Loading environment from %s", join)
		err := godotenv.Load(join)
		if err != nil {
			log.WithErr(err).Alert("Failed to load %s", join)
			return err
		}
	}

	if err := env.ParseWithOptions(&Vars, opts); err != nil {
		log.WithErr(err).Alert("Failed to parse env vars")
		return err
	}

	if Vars.DebugVars {
		log.Debug("Env vars: %v", Vars)
	}

	return nil
}
