package db

import (
	"errors"
	"fmt"
	stdLog "log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db/migrations"

	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

var db *gorm.DB

type Session struct {
	connection *gorm.DB
}

func Init() error {

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Vars.Db.Host,
		config.Vars.Db.Username,
		config.Vars.Db.Password,
		config.Vars.Db.DbName,
		config.Vars.Db.Port)

	dsn_redacted := fmt.Sprintf("host=%s user=%s password=<REDACTED> dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Vars.Db.Host,
		config.Vars.Db.Username,
		config.Vars.Db.DbName,
		config.Vars.Db.Port)

	log.Info("DB DSN: %s", dsn_redacted)

	// TODO replace GORM logger with our logger
	newLogger := gorm_logger.New(
		stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags), // io writer
		gorm_logger.Config{
			SlowThreshold:             time.Second,        // Slow SQL threshold
			LogLevel:                  gorm_logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,               // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,               // Don't include params in the SQL log
			Colorful:                  false,              // Disable color
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.WithErr(err).Alert("Failed to connect to DB")
		return err
	}
	log.Info("Databse connected")

	return nil
}

func Migrate() error {
	s := GetSession(nil)
	return migrations.Run(s.connection)
}

func GetSession(c *fiber.Ctx) Session {
	logger := log.TheLogger
	if c != nil {
		logger = locals.Logger(c)
	}

	s := Session{
		connection: db.Session(&gorm.Session{
			Logger: &dbLogger{
				theLogger: &logger,
			},
		}),
	}
	return s
}

func resultOrError[T any](t *T, result *gorm.DB) (*T, error) {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return t, nil
}

func resultsOrError[S ~[]E, E any](s S, result *gorm.DB) (S, error) {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return s, nil
}
