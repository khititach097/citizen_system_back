package dev_tools_routes

import (
	dev_tools_controllers "citizen_system_back/modules/dev_tools/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterDevToolsRoutes(router *gin.RouterGroup) {
	router.GET("/dev_tools/gen_model/:table_name", dev_tools_controllers.GenModel())
}
