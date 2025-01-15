package asset_services

import (
	"citizen_system_back/database"
	"citizen_system_back/models"
	"fmt"
)

// GetAllAssets retrieves all assets from the database
func GetAllAssets() ([]models.AsLands, error) {
	var db = database.GetDB()
	var assets []models.AsLands
	err := db.Limit(10).Find(&assets).Error
	return assets, err
}

// GetAssetByID retrieves a single asset by its ID
//get geometry with WKB (Well-Known Binary)
// func GetAssetByID(db *gorm.DB, id string) (*models.AsLands, error) {
// 	var asset models.AsLands
// 	if err := db.Where("id = ?", id).First(&asset).Error; err != nil {
// 		return nil, err
// 	}
// 	return &asset, nil
// }

// get geometry with GeoJSON
func GetAssetByID(id string) (*models.AsLands, error) {
	var db = database.GetDB()
	var asset models.AsLands
	if err := db.Select("*, ST_AsGeoJSON(map_geometry) as map_geometry").
		Where("id = ?", id).
		First(&asset).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}

func GetAssetByUserID(userId string, parcelType string, muniCode string) ([]map[string]interface{}, error) {
	var db = database.GetDB()
	var citizens []map[string]interface{}

	query := `
		SELECT c.*, mm.municipality_name_t AS muni_name 
		FROM citizen c
		INNER JOIN master_municipality mm ON c.muni_code = mm.municipality_code
		WHERE c.citizen_id = ?
	`
	if err := db.Raw(query, userId).Scan(&citizens).Error; err != nil {
		return nil, err
	}

	var availableMuniCode []map[string]interface{}
	// var asset []map[string]interface{}
	for _, citizen := range citizens {
		availableMuniCode = append(availableMuniCode, map[string]interface{}{
			"muni_code": citizen["muni_code"],
			"muni_name": citizen["muni_name"],
		})

	}

	fmt.Println(availableMuniCode)

	return availableMuniCode, nil
}

// func findLandData(ownerId string, parcelType string, muniCode string) []map[string]interface{} {

// }

// body
// available_muni_code:[
// 	{
// 		muni_code: "",
// 		muni_name:""
// 	}
// ],
// asset:[
// 	lands:{
// 		...data,
// 		land_used:[
// 			{}
// 		]
// 		asset_image: ""
// 	},
// 	buildings,[
// 		{
// 			...data,
// 			asset_image: ""
// 		}
// 	]
// 	signboards,[
// 		{
// 			...data,
// 			asset_image: ""
// 		}
// 	]
// ]
