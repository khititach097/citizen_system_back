package dev_tools_routes

import (
	"citizen_system_back/controllers/dev_tools_controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDevToolsRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/dev_tools/gen_model/:table_name", dev_tools_controllers.GenModel(db))
}
