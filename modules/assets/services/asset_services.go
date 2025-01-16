package asset_services

import (
	"citizen_system_back/database"
	"citizen_system_back/models"
	asset_types "citizen_system_back/modules/assets/types"
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

func GetAssetByUserID(query asset_types.QueryParams) (interface{}, error) {
	var db = database.GetDB()
	page := 1
	pageSize := 10

	if query.Page != "" {
		fmt.Sscanf(query.Page, "%d", &page)
	}
	if query.PageSize != "" {
		fmt.Sscanf(query.PageSize, "%d", &pageSize)
	}

	offset := (page - 1) * pageSize

	// Get citizen information
	citizenSQL := `
		SELECT c.*, mm.municipality_name_t AS muni_name 
		FROM citizen c
		INNER JOIN master_municipality mm ON c.muni_code = mm.municipality_code
		WHERE c.citizen_id = ?`

	var citizens []map[string]interface{}
	if err := db.Raw(citizenSQL, query.UserID).Scan(&citizens).Error; err != nil {
		return nil, fmt.Errorf("error querying citizens: %w", err)
	}

	// fmt.Println("citizens ***>>>", citizens)

	// Create a direct map to hold the results
	result := map[string]interface{}{
		"available_muni_code": []map[string]string{},
		"data":                []interface{}{},
		"total":               0,
	}

	// Process the citizens data
	for _, citizen := range citizens {
		muniCode, ok := citizen["muni_code"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid muni_code")
		}
		muniName, ok := citizen["muni_name"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid muni_name")
		}

		// Append municipality data
		result["available_muni_code"] = append(result["available_muni_code"].([]map[string]string), map[string]string{
			"muni_code": muniCode,
			"muni_name": muniName,
		})

		// Build the base query for land IDs
		landIDsSQL := `
			SELECT id as land_id, al.muni_code 
			FROM as_lands al 
			WHERE id IN (
				SELECT DISTINCT al.id AS land_id
				FROM as_lands al
				LEFT JOIN as_land_owners alo ON al.id = alo.land_id
				WHERE alo.owner_id = ? 
				UNION
				SELECT DISTINCT al.id AS land_id
				FROM as_lands al
				LEFT JOIN as_buildings ab ON al.id = ab.land_id
				WHERE ab.id IN (
					SELECT building_id
					FROM as_building_owners abo
					WHERE abo.owner_id = ? 
				)
				UNION
				SELECT DISTINCT al.id AS land_id
				FROM as_lands al
				LEFT JOIN as_signboards as2 ON al.id = as2.land_id
				WHERE as2.id IN (
					SELECT signboard_id
					FROM as_signboard_owners aso
					WHERE aso.owner_id = ? 
				)
			)`

		fmt.Println("citizen id ***>>>", citizen["id"])
		args := []interface{}{citizen["id"], citizen["id"], citizen["id"]}

		if query.MuniCode != "" {
			landIDsSQL += " AND al.muni_code = ?"
			args = append(args, query.MuniCode)
		}
		if query.ParcelType != "" {
			landIDsSQL += " AND al.parcel_type = ?"
			args = append(args, query.ParcelType)
		}

		landIDsSQL += " LIMIT ? OFFSET ?"
		args = append(args, pageSize, offset)

		// fmt.Println("all land id ***>>>", landIDs)

		var landIDs []map[string]interface{}
		if err := db.Raw(landIDsSQL, args...).Scan(&landIDs).Error; err != nil {
			return nil, fmt.Errorf("error querying land IDs: %w", err)
		}

		fmt.Println("all land id ***>>>", landIDs)
		// Process land IDs and get details
		for _, landID := range landIDs {
			landIDVal, ok := landID["land_id"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid land_id")
			}
			muniCodeVal, ok := landID["muni_code"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid muni_code")
			}

			assetData, err := getLandDetails(landIDVal, muniCodeVal)
			if err != nil {
				return nil, fmt.Errorf("error getting land details: %w", err)
			}

			// Append land details to the data
			result["data"] = append(result["data"].([]interface{}), assetData)
		}

		// Get total count using similar query without pagination
		countSQL := `
			SELECT COUNT(DISTINCT al.id) 
			FROM as_lands al
			WHERE al.id IN (
				SELECT DISTINCT al.id AS land_id
				FROM as_lands al
				LEFT JOIN as_land_owners alo ON al.id = alo.land_id
				WHERE alo.owner_id = ? 
				UNION
				SELECT DISTINCT al.id AS land_id
				FROM as_lands al
				LEFT JOIN as_buildings ab ON al.id = ab.land_id
				WHERE ab.id IN (
					SELECT building_id
					FROM as_building_owners abo
					WHERE abo.owner_id = ? 
				)
				UNION
				SELECT DISTINCT al.id AS land_id
				FROM as_lands al
				LEFT JOIN as_signboards as2 ON al.id = as2.land_id
				WHERE as2.id IN (
					SELECT signboard_id
					FROM as_signboard_owners aso
					WHERE aso.owner_id = ? 
				)
			)`

		countArgs := []interface{}{citizen["id"], citizen["id"], citizen["id"]}

		if query.MuniCode != "" {
			countSQL += " AND al.muni_code = ?"
			countArgs = append(countArgs, query.MuniCode)
		}
		if query.ParcelType != "" {
			countSQL += " AND al.parcel_type = ?"
			countArgs = append(countArgs, query.ParcelType)
		}

		// fmt.Println("countSQL ***>>>", countSQL)

		var total int64
		if err := db.Raw(countSQL, countArgs...).Count(&total).Error; err != nil {
			return nil, fmt.Errorf("error getting total count: %w", err)
		}

		// fmt.Println("total ***>>>", total)

		result["total"] = int(total)
	}

	return result, nil
}

func getLandDetails(landID string, muniCode string) (interface{}, error) {
	var db = database.GetDB()

	// Get land details
	landSQL := `
		SELECT 
			al.id,
			al.deed_no,
			pdt.doc_type_name,
			al.land_space_rai,
			al.land_space_ngan,
			al.land_space_wa,
			al.estimated_price,
			al.estimated_price_by_treasury,
			al.muni_code,
			md.province_name_t AS province_name,
			md.district_name_t AS district_name,
			ms.subdistrict_name_t AS subdistrict_name,
			mp.postcodemain
		FROM as_lands al
		LEFT JOIN property_doc_type pdt 
			ON al.parcel_type = pdt.doc_type_id
		LEFT JOIN master_district md 
			ON md.district_code::numeric = al.land_district_id
		LEFT JOIN master_subdistrict ms 
			ON ms.subdistrict_code::numeric = al.land_sub_district_id
		LEFT JOIN (
			SELECT DISTINCT ON (districtid) districtid, postcodemain 
			FROM master_postcode
		) mp
			ON md.district_code = mp.districtid
		WHERE al.id = ?`

	var landData map[string]interface{}
	if err := db.Raw(landSQL, landID).Scan(&landData).Error; err != nil {
		return nil, fmt.Errorf("error getting land details: %w", err)
	}

	// Get land used
	landUsedSQL := `
		SELECT put.using_type_detail, total_space_rai, total_space_ngan, total_space_square_wa 
		FROM as_land_used alu 
		LEFT JOIN property_using_type put ON alu.using_type_id = put.using_type_id 
		WHERE land_id = ?`
	var landUsed []map[string]interface{}
	if err := db.Raw(landUsedSQL, landID).Scan(&landUsed).Error; err != nil {
		return nil, fmt.Errorf("error getting land used: %w", err)
	}
	landData["land_used"] = landUsed

	// Get buildings
	buildingsSQL := `
		SELECT
			ab.id,
			pbmt.building_type_name,
			ab.building_width_meter,
			ab.building_length_meter,
			ab.property_building_main_type_id,
			property_building_sub_type_id,
			(EXTRACT(YEAR FROM CURRENT_DATE) + 543 - CAST(ab.build_year AS INTEGER) + 1) AS building_year_total
		FROM as_buildings ab
		LEFT JOIN property_building_main_type pbmt ON ab.property_building_main_type_id = pbmt.building_type_id
		WHERE land_id = ? AND pbmt.muni_code = ?`

	var buildings []map[string]interface{}
	if err := db.Raw(buildingsSQL, landID, muniCode).Scan(&buildings).Error; err != nil {
		return nil, fmt.Errorf("error getting buildings: %w", err)
	}

	// Get building details
	for i := range buildings {
		buildingUsedSQL := `
			SELECT
				put.using_type_detail,
				abu.building_using_type_id,
				abu.building_household_type_id,
				abu.building_cultivation_space_square_meter,
				abu.building_self_using_square_meter,
				abu.building_for_rent_square_meter,
				abu.building_empty_space_square_meter,
				abu.building_etc_square_meter
			FROM as_building_used abu
			LEFT JOIN property_using_type put ON abu.building_using_type_id = put.using_type_id
			WHERE building_id = ?`

		var buildingUsed []map[string]interface{}
		if err := db.Raw(buildingUsedSQL, buildings[i]["id"]).Scan(&buildingUsed).Error; err != nil {
			return nil, fmt.Errorf("error getting building used: %w", err)
		}
		buildings[i]["building_used"] = buildingUsed

		assetImages, err := getAssetAttachmentLatestSurveyByAssetId(buildings[i]["id"].(string), landID, map[string]string{"muni_code": muniCode})
		if err != nil {
			return nil, fmt.Errorf("error getting building asset images: %w", err)
		}
		buildings[i]["asset_images"] = assetImages
	}

	// Get signboards
	signboardsSQL := `
		SELECT
			id,
			signboard_width,
			signboard_height,
			signboard_side,
			as2.property_signboard_display_type_id,
			psdt.property_signboard_display_type,
			psdt.property_signboard_display_name,
			psdt.property_signboard_display_desc
		FROM as_signboards as2
		LEFT JOIN property_signboard_display_type psdt 
			ON psdt.property_signboard_display_type = as2.property_signboard_display_type_id
		WHERE land_id = ?`

	var signboards []map[string]interface{}
	if err := db.Raw(signboardsSQL, landID).Scan(&signboards).Error; err != nil {
		return nil, fmt.Errorf("error getting signboards: %w", err)
	}

	// Get asset images for each signboard
	for i := range signboards {
		assetImages, err := getAssetAttachmentLatestSurveyByAssetId(signboards[i]["id"].(string), landID, map[string]string{"muni_code": muniCode})
		if err != nil {
			return nil, fmt.Errorf("error getting signboard asset images: %w", err)
		}
		signboards[i]["asset_images"] = assetImages
	}

	// Get asset images for land
	landAssetImages, err := getAssetAttachmentLatestSurveyByAssetId(landID, landID, map[string]string{"muni_code": muniCode})
	if err != nil {
		return nil, fmt.Errorf("error getting land asset images: %w", err)
	}
	landData["asset_images"] = landAssetImages

	// Return the combined data directly as a map
	return map[string]interface{}{
		"land_data":  landData,
		"buildings":  buildings,
		"signboards": signboards,
	}, nil
}

func getAssetAttachmentLatestSurveyByAssetId(assetID string, landID string, params map[string]string) ([]asset_types.AssetAttachment, error) {
	var db = database.GetDB()
	query := `
		SELECT
				aa.img_id,
				aa.key,
				aa.image_from,
				aa.survey_request_id
		FROM asset_attachment aa
		WHERE aa.asset_id = ? 
				AND aa.muni_code = ? 
				AND (
						CASE 
								WHEN NOT EXISTS (
										SELECT 1 
										FROM survey_request sr 
										WHERE sr.land_id = ? 
												AND sr.muni_code = ? 
										ORDER BY sr.created_at DESC LIMIT 1
								) THEN aa.image_from = 3
								ELSE aa.survey_request_id = (
										SELECT 
												sr.id 
										FROM survey_request sr
										WHERE sr.land_id = ? 
												AND sr.muni_code = ? 
										ORDER BY sr.created_at DESC
										LIMIT 1
								)
						END
				)
		ORDER BY aa.created_date`

	var attachments []asset_types.AssetAttachment

	err := db.Raw(
		query,
		assetID,
		params["muni_code"],
		landID,
		params["muni_code"],
		landID,
		params["muni_code"],
	).Scan(&attachments).Error

	if err != nil {
		// Log the error but return an empty slice as per original function
		fmt.Printf("Error fetching asset attachments: %v\n", err)
		return []asset_types.AssetAttachment{}, nil
	}

	if len(attachments) == 0 {
		return []asset_types.AssetAttachment{}, nil
	}

	return attachments, nil
}
