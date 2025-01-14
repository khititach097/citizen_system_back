package database

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

var logger = logrus.New()

func setupLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

// InitDB initializes and returns a database connection
func InitDB() (*gorm.DB, error) {
	setupLogger()

	// config := Config{
	// 	Host:     "bedrock-dev-db.cluster-cq6wq7ckjmhj.ap-southeast-1.rds.amazonaws.com",
	// 	Port:     "5432",
	// 	User:     "sts_dev_app",
	// 	Password: "9{Ll&&6{!Cm4d5M#",
	// 	DBName:   "sts_dev",
	// 	SSLMode:  "verify-full",
	// 	SSLCert:  os.Getenv("DATABASE_SSL_CERT"),
	// }
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
		logger.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	certBytes, err := base64.StdEncoding.DecodeString(config.SSLCert)
	if err != nil {
		logger.Fatalf("Failed to decode SSL certificate: %v", err)
	}
	if _, err := tmpFile.Write(certBytes); err != nil {
		logger.Fatalf("Failed to write certificate to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		logger.Fatalf("Failed to close temp file: %v", err)
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

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	logger.Info("Database connection established with SSL")
	return db, nil
}
