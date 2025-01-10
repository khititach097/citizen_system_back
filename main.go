package main
import (
	_ "citizen_system_back/docs"
	"fmt"
	"os"
	"log"
	"encoding/base64"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)


// @title Asset Tax API
// @version 1.0
// @description API to manage assets and calculate taxes.
// @host localhost:8080
// @BasePath /api/v1

// Config holds database configuration

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

// Database instance
var db *gorm.DB

// initDB initializes the database connection with SSL
// initDB initializes the database connection with SSL
func initDB() *gorm.DB {
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

// setupLogger configures the logger
func setupLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

// main is the entry point for the application
func main() {
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}
	// Setup logger
	setupLogger()

	// Initialize database
	db = initDB()

	// Setup Gin
	router := gin.Default()

	// Swagger endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//to test go to http://localhost:8080/docs/swagger/index.html

	// Define API routes
	api := router.Group("/api/v1")
	{
		api.GET("/assets", getAssets)          // List all assets
		api.POST("/assets", createAsset)      // Create a new asset
		api.PUT("/assets/:id", updateAsset)   // Update an asset
		api.DELETE("/assets/:id", deleteAsset) // Delete an asset
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Infof("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

// @Summary Get all assets
// @Description Fetches a list of all assets
// @Tags assets
// @Success 200 {object} map[string]string "List of assets"
// @Failure 500 {string} string "Internal Server Error"
// @Router /assets [get]
func getAssets(c *gin.Context) {
	logger.Info("Fetching all assets")
	c.JSON(200, gin.H{"message": "List all assets"})
}

func createAsset(c *gin.Context) {
	logger.Info("Creating a new asset")
	c.JSON(201, gin.H{"message": "Create a new asset"})
}

func updateAsset(c *gin.Context) {
	logger.Info("Updating an asset")
	c.JSON(200, gin.H{"message": "Update an asset"})
}

func deleteAsset(c *gin.Context) {
	logger.Info("Deleting an asset")
	c.JSON(200, gin.H{"message": "Delete an asset"})
}
