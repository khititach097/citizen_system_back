package main

import (
	"citizen_system_back/database"
	// "citizen_system_back/middleware"
	"citizen_system_back/routes"
	"fmt"
	"os"

	_ "citizen_system_back/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logrus.New()

func setupLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	setupLogger()

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		logger.Fatalf("Database initialization failed: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Swagger endpoint at /api/v1/docs
	router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes with db
	routes.RegisterRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("Starting server on port %s", port)

	// Hyperlink escape sequence
	yellow := "\033[33m"
	reset := "\033[0m"
	link := "http://localhost:8080/api/v1/docs/swagger/index.html"

	// Print the hyperlink
	fmt.Println(yellow + "Swagger: " + link + reset)

	if err := router.Run(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
