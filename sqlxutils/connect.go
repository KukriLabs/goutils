package sqlxutils

import (
	"log"
	"math"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPostgresConfig is a simple helper to build *ConnectionConfig with some sane defaults
func NewPostgresConfig(uri string) *ConnectionConfig {
	return &ConnectionConfig{
		URI:                      uri,
		Database:                 "postgres",
		MaxConnectionLifetime:    time.Duration(1) * time.Hour,
		MaxIdleConns:             20,
		MaxOpenConns:             50,
		ConnectSleepStartSeconds: 1,
		MaxDBConnectAttempts:     5,
		DBConnectAttempts:        0,
	}
}

type ConnectionConfig struct {
	URI                      string
	Database                 string
	MaxConnectionLifetime    time.Duration
	MaxIdleConns             int
	MaxOpenConns             int
	ConnectSleepStartSeconds int
	MaxDBConnectAttempts     int
	DBConnectAttempts        int
}

// MustConnect is an advanced connector that also includes exponential backoff for
// multiple connection attempts. Can be useful when waiting for a database to start
func MustConnect(config *ConnectionConfig) *sqlx.DB {
	db, err := sqlx.Connect(config.Database, config.URI)

	if err != nil {
		if config.DBConnectAttempts == config.MaxDBConnectAttempts {
			log.Fatalln("Unable to connect to database", err)
		}
		// Exponential backoff so as not to swamp server
		sleepTime := time.Duration(config.ConnectSleepStartSeconds*int(math.Pow(2, float64(config.DBConnectAttempts)))) * time.Second
		config.DBConnectAttempts++
		time.Sleep(sleepTime)
		return MustConnect(config)
	}

	db.SetConnMaxLifetime(config.MaxConnectionLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
