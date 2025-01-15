package models

import (
)

// PropertySurveyType represents the structure of the property_survey_type table
type PropertySurveyType struct {
	SurveyTypeId int `gorm:"column:survey_type_id;primaryKey" json:"survey_type_id"`
	SurveyTypeDetailName *string `gorm:"column:survey_type_detail_name" json:"survey_type_detail_name"`
}
