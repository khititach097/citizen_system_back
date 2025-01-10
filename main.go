package main
import (
	_ "citizen_system_back/docs"
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
	"citizen_system_back/database"
)



// Logger instance
var logger = logrus.New()

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
	database.InitDB()

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
