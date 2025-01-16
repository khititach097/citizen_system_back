package asset_types

type QueryParams struct {
	Page       string
	PageSize   string
	MuniCode   string
	ParcelType string
	UserID     string
}

type AssetAttachment struct {
	ImgID           string `json:"img_id" gorm:"column:img_id"`
	Key             string `json:"key" gorm:"column:key"`
	ImageFrom       int    `json:"image_from" gorm:"column:image_from"`
	SurveyRequestID string `json:"survey_request_id" gorm:"column:survey_request_id"`
}
