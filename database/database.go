package database

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"encoding/base64"
	"io/ioutil"
	"github.com/sirupsen/logrus"
)

// @title Asset Tax API
// @version 1.0
// @description API to manage assets and calculate taxes.
// @host localhost:8080
// @BasePath /api/v1

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	SSLCert  string // New field for SSL certificate
}

// Logger instance
var logger = logrus.New()

// setupLogger configures the logger
func setupLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

// initDB initializes the database connection with SSL
func InitDB() *gorm.DB {
	// Initialize the logger
	setupLogger() // Ensure logger is set up before usage

	config := Config{
		Host:     "bedrock-dev-db.cluster-cq6wq7ckjmhj.ap-southeast-1.rds.amazonaws.com",
		Port:     "5432",
		User:     "sts_dev_app",
		Password: "9{Ll&&6{!Cm4d5M#",
		DBName:   "sts_dev",
		SSLMode:  "verify-full",
		SSLCert:  os.Getenv("DATABASE_SSL_CERT"),
	}

	// Create a temporary file to store the certificate
	tmpFile, err := ioutil.TempFile("", "postgres-cert-*.pem")
	if err != nil {
		logger.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Decode base64 SSL certificate and write to temp file
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

	// Create the connection string with SSL configuration
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s sslrootcert=%s",
		config.Host, 
		config.Port, 
		config.User, 
		config.Password, 
		config.DBName, 
		config.SSLMode,
		tmpFile.Name(),
	)

	// Initialize the database connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		logger.Fatalf("Failed to connect to the database: %v", err)
	}

	logger.Info("Database connection established with SSL")
	return db
}
