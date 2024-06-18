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

	ppseeders "github.com/jak103/powerplay/internal/db/seeders"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var dbConnection *gorm.DB

type session struct {
	*gorm.DB
}

func Init() error {

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Vars.Db.Host,
		config.Vars.Db.Username,
		config.Vars.Db.Password,
		config.Vars.Db.DbName,
		config.Vars.Db.Port)

	dsnRedacted := fmt.Sprintf("host=%s user=%s password=<REDACTED> dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Vars.Db.Host,
		config.Vars.Db.Username,
		config.Vars.Db.DbName,
		config.Vars.Db.Port)

	log.Info("DB DSN: %s", dsnRedacted)

	// TODO replace GORM logger with our logger
	newLogger := gormlogger.New(
		stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags), // io writer
		gormlogger.Config{
			SlowThreshold:             time.Second,       // Slow SQL threshold
			LogLevel:                  gormlogger.Silent, // Log level
			IgnoreRecordNotFoundError: true,              // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,              // Don't include params in the SQL log
			Colorful:                  false,             // Disable color
		},
	)

	dbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.WithErr(err).Alert("Failed to connect to DB")
		return err
	}
	log.Info("Database connected")

	return nil
}

func Migrate() error {
	logger := log.TheLogger

	s := dbConnection.Session(&gorm.Session{
		Logger: &dbLogger{
			theLogger: &logger,
		},
	})

	return migrations.Run(s)
}

func GetDB() *gorm.DB {
	return dbConnection
}

func GetSession(c *fiber.Ctx) session {

	logger := log.TheLogger
	if c != nil {
		logger = locals.Logger(c)
	}

	return session{
		dbConnection.Session(&gorm.Session{
			Logger: &dbLogger{
				theLogger: &logger,
			},
		}),
	}
}

func RunSeeders(seeders []ppseeders.Seeder, args ...interface{}) error {
	s := GetSession(nil)
	for _, seeder := range seeders {
		if _, err := seeder.Seed(s.DB, args...); err != nil {
			return err
		}
	}
	return nil
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
