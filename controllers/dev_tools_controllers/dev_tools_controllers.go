package dev_tools_controllers

import (
	"citizen_system_back/services/dev_tools_services"
	"fmt"
	"strings"

	"errors"
	// "fmt"
	"net/http"

	"citizen_system_back/utils/response"

	"github.com/gin-gonic/gin"
)

// @Summary Generate model
// @Description Generate model
// @Tags dev_tools
// //@Security CookieAuth
// @Accept json
// @Produce json
// @Param table_name path string true "The name of the table to generate model for"  // Add description here
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/dev_tools/gen_model/{table_name} [get] // Update the URL path to include {table_name}
func GenModel() gin.HandlerFunc {
	return func(c *gin.Context) {
		table_name := c.Param("table_name")
		url := c.Request.URL.String()
		domain := c.Request.Host
		fmt.Println("table_name ***>>>" + table_name)
		fmt.Println("url ***>>>" + url)
		fmt.Println("domain ***>>>" + domain)

		// Check if the domain does not include "localhost"
		if !strings.Contains(domain, "localhost") {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access not allowed from this domain",
			})
			return
		}

		if table_name == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid table name", errors.New("table name cannot be empty")))
			return
		}

		result := dev_tools_services.GenModel(table_name)

		c.JSON(http.StatusOK, response.NewResponse("Gen model successfully", result))
	}
}
