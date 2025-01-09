package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/swaggo/files"          // Alias as swaggerFiles
    "github.com/swaggo/gin-swagger"    // Swagger handler
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

	//_ "citizen_system_back/docs" // Swag generated docs package
	//_ "example.com/citizen_system_back/docs"

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
}

// Logger instance
var logger = logrus.New()

// Database instance
var db *gorm.DB

// initDB initializes the database connection
func initDB() *gorm.DB {
	config := Config{
		Host:     "bedrock-dev-db.cluster-cq6wq7ckjmhj.ap-southeast-1.rds.amazonaws.com",
		Port:     "5432",
		User:     "sts_dev_app",
		Password: "9{Ll&&6{!Cm4d5M#",
		DBName:   "sts_dev",
		SSLMode:  "disable",
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to the database: %v", err)
	}

	logger.Info("Database connection established")
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
	// Setup logger
	setupLogger()

	// Initialize database
	db = initDB()

	// Setup Gin
	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

// Example handlers
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
