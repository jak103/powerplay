package db

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jak103/powerplay/internal/db/migrations"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbTestingSuite struct {
	suite.Suite
	session session
	pool    *dockertest.Pool
	db      *dockertest.Resource
}

func TestDatabase(t *testing.T) {
	suite.Run(t, new(dbTestingSuite))
}

func (suite *dbTestingSuite) SetupSuite() {
	log.Info("Starting db container")
	start := time.Now()
	dbConnection = suite.startDb()
	if dbConnection == nil {
		suite.FailNow("Failed to start db")
	}

	suite.session = session{
		DB: dbConnection.Session(&gorm.Session{
			Logger: logger.Default.LogMode(logger.Silent),
		}),
	}
	log.Info("Done! (%v)", time.Since(start))
}

func (suite *dbTestingSuite) TearDownSuite() {
	if err := suite.pool.Purge(suite.db); err != nil {
		log.Error("Could not purge resource: %s", err)
	}
}

func (s *dbTestingSuite) startDb() *gorm.DB {
	var db *gorm.DB
	var err error
	pwd, err := os.Getwd()
	if err != nil {
		log.Error("failed to get working directory: %s", err)
	}

	dbDataDir, err := filepath.Abs(fmt.Sprintf("%s/../../../db", pwd))
	if err != nil {
		log.WithErr(err).Error("Failed to get db data directory")
		return nil
	}

	log.Info("Mounting %s", dbDataDir)

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	s.pool, err = dockertest.NewPool("")
	if err != nil {
		log.Error("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = s.pool.Client.Ping()
	if err != nil {
		log.Error("Could not connect to Docker: %s", err)
	}
	log.Info("Connected to docker engine")

	s.pool.Client.Logs(docker.LogsOptions{
		RawTerminal:  false,
		Stdout:       false,
		Stderr:       false,
		Follow:       false,
		ErrorStream:  io.Discard,
		OutputStream: io.Discard,
	})

	// pulls an image, creates a container based on it and runs it
	s.db, err = s.pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16",
		Env: []string{
			"PGUSER=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=password",
			"POSTGRES_DB=powerplay",
			"listen_addresses = '*'",
		},

		Mounts: []string{
			fmt.Sprintf("%s:/docker-entrypoint-initdb.d", dbDataDir),
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	log.Info("Started db docker instance")
	if err != nil {
		log.Error("Could not start resource: %s", err)
	}
	time.Sleep(10 * time.Second) // Let db start
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := s.pool.Retry(func() error {
		dsn := fmt.Sprintf("host=localhost user=postgres password=password dbname=powerplay port=%s sslmode=disable", s.db.GetPort(("5432/tcp")))
		// // dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		// 	"postgres",
		// 	"password",
		// 	"localhost",
		// 	s.db.GetPort("5432/tcp"),
		// 	"powerplay",
		// )
		log.Info("Trying to ping db at %s", dsn)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			return err
		}
		d, _ := db.DB()
		return d.Ping()
	}); err != nil {
		log.Error("Could not connect to database: %s", err)
		return nil
	}

	migrations.Run(db)

	return db
}

func TestResultOrErrorNominal(t *testing.T) {
	result := &gorm.DB{
		Error: nil,
	}

	var record *models.KeyRecord = &models.KeyRecord{
		UserId: 99,
	}

	r, err := resultOrError(record, result)

	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestResultOrErrorNoResult(t *testing.T) {
	result := &gorm.DB{
		Error: gorm.ErrRecordNotFound,
	}

	var record *models.KeyRecord = &models.KeyRecord{
		UserId: 99,
	}

	r, err := resultOrError(record, result)

	assert.Nil(t, r)
	assert.Nil(t, err)
}

func TestResultOrErrorError(t *testing.T) {
	result := &gorm.DB{
		Error: gorm.ErrDuplicatedKey,
	}

	var record *models.KeyRecord = &models.KeyRecord{
		UserId: 99,
	}

	r, err := resultOrError(record, result)

	assert.Nil(t, r)
	assert.NotNil(t, err)
}
