package s3_controllers

import (
	s3_services "citizen_system_back/modules/s3/services"
	"citizen_system_back/utils/response"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Preview image
// @Description Preview image by img id
// @Tags s3
// // @Security BearerAuth
// @Accept json
// @Produce json
// @Param img_id path string true "The UUID (PK) of image id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/s3/preview/{img_id} [get]
func PreviewImageByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		imgIDStr := c.Param("img_id")
		imgID, err := uuid.Parse(imgIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				"Invalid image ID format",
				err,
			))
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
				"Failed to initialize S3 client",
				err,
			))
			return
		}
		fmt.Println("img_id ***>>>" + imgIDStr)

		if imgIDStr == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid img id", errors.New("img id cannot be empty")))
			return
		}

		s3_services.PreviewImageByID(c, imgID)

		// c.JSON(http.StatusOK, response.NewResponse("Gen model successfully", result))
	}
}
