package main

import (
	"citizen_system_back/database"
	"citizen_system_back/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "citizen_system_back/docs"
)

var logger = logrus.New()

func setupLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

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
	if err := router.Run(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
