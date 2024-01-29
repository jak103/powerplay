package db

import (
	"fmt"

	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db/migrations"
	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

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

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithErr(err).Alert("Failed to connect to DB")
		return err
	}
	log.Info("Databse connected")

	return nil
}

func Migrate() error {
	s := GetSession()
	return migrations.Run(s)
}

func GetSession() *gorm.DB {
	// TODO set the logger in here so that it logs DB stuff correctly
	return db.Session(&gorm.Session{})
}
