package survey_request_survices

import (
	"citizen_system_back/database"
	"citizen_system_back/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindLatestSurveyRequestByAssetId(
	assetID string,
	surveyType string,
	condoRoomID string,
) interface{} {
	var db = database.GetDB()

	whereOption := map[string]interface{}{
		"land_id": assetID,
	}

	// Default to "LAND" if surveyType is not provided
	if surveyType == "" {
		surveyType = "LAND"
	}

	var result models.SurveyRequest

	if surveyType == "LAND" {
		if assetID == "" {
			return nil
		}

		// Query for "LAND"
		err := db.Where(whereOption).Order("created_at DESC").First(&result).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil
			}
			fmt.Println("Error querying for LAND survey request:", err)
			return nil
		}
		return result
	} else if surveyType == "CONDO" {
		// Query for "CONDO"
		err := db.Where("land_id = ? AND survey_type_id = ? AND muni_code = ?", assetID, "YEARLY_OR_CHANGE_CONDO").
			Order("created_at DESC").
			First(&result).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil
			}
			fmt.Println("Error querying for CONDO survey request:", err)
			return nil
		}
		return result
	} else { // "ROOM"
		// Query for "ROOM"
		query := `
			SELECT sr.id, sr.created_at
			FROM survey_request sr
			WHERE sr.id IN (
				SELECT as2.survey_request_id
				FROM asset_survey as2
				WHERE as2.asset_id = ?
			)
			ORDER BY sr.created_at DESC
			LIMIT 1`
		err := db.Raw(query, condoRoomID).Scan(&result).Error
		if err != nil {
			fmt.Println("Error querying for ROOM survey request:", err)
			return nil
		}
		if result.Id != uuid.Nil {
			return result
		}
		return nil
	}
}
