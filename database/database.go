// database/db.go
package database

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	once   sync.Once
	logger = logrus.New()
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	SSLCert  string
}

func setupLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return db
}

// InitDB initializes the database connection
func InitDB() error {
	var initError error

	once.Do(func() {
		setupLogger()

		config := Config{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			DBName:   os.Getenv("DATABASE_NAME"),
			SSLMode:  "verify-full",
			SSLCert:  os.Getenv("DATABASE_SSL_CERT"),
		}

		// Create temporary file for SSL certificate
		tmpFile, err := ioutil.TempFile("", "postgres-cert-*.pem")
		if err != nil {
			initError = fmt.Errorf("failed to create temporary file: %w", err)
			logger.Error(initError)
			return
		}
		defer os.Remove(tmpFile.Name())

		certBytes, err := base64.StdEncoding.DecodeString(config.SSLCert)
		if err != nil {
			initError = fmt.Errorf("failed to decode SSL certificate: %w", err)
			logger.Error(initError)
			return
		}

		if _, err := tmpFile.Write(certBytes); err != nil {
			initError = fmt.Errorf("failed to write certificate to temp file: %w", err)
			logger.Error(initError)
			return
		}

		if err := tmpFile.Close(); err != nil {
			initError = fmt.Errorf("failed to close temp file: %w", err)
			logger.Error(initError)
			return
		}

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s sslrootcert=%s",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBName,
			config.SSLMode,
			tmpFile.Name(),
		)

		db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
		if err != nil {
			initError = fmt.Errorf("failed to connect to the database: %w", err)
			logger.Error(initError)
			return
		}

		logger.Info("Database connection established with SSL")
	})

	return initError
}
