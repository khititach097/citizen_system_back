package asset_services

import (
	"citizen_system_back/database"
	"citizen_system_back/models"
	asset_types "citizen_system_back/modules/assets/types"
	survey_request_survices "citizen_system_back/modules/survey_request/services"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
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
			WITH combined_assets AS (
				SELECT 
						id AS asset_id,
						al.id as land_id,
						al.muni_code, 
						al.parcel_no, 
						al.parcel_type, 
						'land' AS asset_type
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
				)
				UNION
				SELECT 
						ac.id AS asset_id, 
						al.id as land_id,
						ac.muni_code, 
						al.parcel_no, 
						al.parcel_type, 
						'condo' AS asset_type
				FROM as_condos ac
				LEFT JOIN as_lands al ON ac.land_id = al.id
				WHERE ac.id IN (
						SELECT acr.condo_id 
						FROM as_condo_rooms acr 
						WHERE acr.id IN (
								SELECT aco.condo_room_id 
								FROM as_condo_owners aco 
								WHERE aco.owner_id = ?
						)
				)
		)
		SELECT * 
		FROM combined_assets
		ORDER BY muni_code, parcel_no `

		fmt.Println("citizen id ***>>>", citizen["id"])
		args := []interface{}{citizen["id"], citizen["id"], citizen["id"], citizen["id"]}

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

		var allAssets []map[string]interface{}
		if err := db.Raw(landIDsSQL, args...).Scan(&allAssets).Error; err != nil {
			return nil, fmt.Errorf("error querying land IDs: %w", err)
		}

		fmt.Println("all asset id ***>>>", allAssets)
		// Process land IDs and get details
		for _, asset := range allAssets {
			assetIDVal, assetIdOk := asset["asset_id"].(string)
			if !assetIdOk {
				return nil, fmt.Errorf("invalid asset_id")
			}
			muniCodeVal, muniCodeOk := asset["muni_code"].(string)
			if !muniCodeOk {
				return nil, fmt.Errorf("invalid muni_code")
			}
			assetTypeVal, assetTypeOk := asset["asset_type"].(string)
			if !assetTypeOk {
				return nil, fmt.Errorf("invalid asset_type")
			}

			var assetData interface{}
			var err error

			if assetTypeVal == "land" {
				assetData, err = getLandDetails(assetIDVal, muniCodeVal)
				if err != nil {
					return nil, fmt.Errorf("error getting land details: %w", err)
				}
			}
			if assetTypeVal == "condo" {
				assetData, err = getCondoDetails(assetIDVal, muniCodeVal)
				if err != nil {
					return nil, fmt.Errorf("error getting condo details: %w", err)
				}
			}

			// Append land details to the data
			result["data"] = append(result["data"].([]interface{}), assetData)
		}

		// Get total count using similar query without pagination
		countSQL := `
			WITH combined_assets AS (
				SELECT 
						id AS asset_id,
						al.id as land_id,
						al.muni_code, 
						al.parcel_no, 
						al.parcel_type, 
						'land' AS asset_type
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
				)
				UNION
				SELECT 
						ac.id AS asset_id, 
						al.id as land_id,
						ac.muni_code, 
						al.parcel_no, 
						al.parcel_type, 
						'condo' AS asset_type
				FROM as_condos ac
				LEFT JOIN as_lands al ON ac.land_id = al.id
				WHERE ac.id IN (
						SELECT acr.condo_id 
						FROM as_condo_rooms acr 
						WHERE acr.id IN (
								SELECT aco.condo_room_id 
								FROM as_condo_owners aco 
								WHERE aco.owner_id = ?
						)
				)
		)
		SELECT count(*) 
		FROM combined_assets`

		countArgs := []interface{}{citizen["id"], citizen["id"], citizen["id"], citizen["id"]}

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

	var buildings = []map[string]interface{}{}
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

		var buildingUsed = []map[string]interface{}{}
		if err := db.Raw(buildingUsedSQL, buildings[i]["id"]).Scan(&buildingUsed).Error; err != nil {
			return nil, fmt.Errorf("error getting building used: %w", err)
		}
		buildings[i]["building_used"] = buildingUsed

		assetImages := GetAssetAttachmentLatestSurveyByAssetId(buildings[i]["id"].(string), landID)

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

	var signboards = []map[string]interface{}{}
	if err := db.Raw(signboardsSQL, landID).Scan(&signboards).Error; err != nil {
		return nil, fmt.Errorf("error getting signboards: %w", err)
	}

	// Get asset images for each signboard
	for i := range signboards {
		assetImages := GetAssetAttachmentLatestSurveyByAssetId(signboards[i]["id"].(string), landID)
		signboards[i]["asset_images"] = assetImages
	}

	// Get asset images for land
	landAssetImages := GetAssetAttachmentLatestSurveyByAssetId(landID, landID)

	landData["asset_images"] = landAssetImages

	return map[string]interface{}{
		"asset_type": "land",
		"land_data":  landData,
		"buildings":  buildings,
		"signboards": signboards,
	}, nil
}

func getCondoDetails(condoId string, muniCode string) (interface{}, error) {
	var db = database.GetDB()

	condoSQL := `
		SELECT 
				ac.*,
				al.parcel_type,
				pdt.doc_type_name,
				md.province_name_t AS province_name,
				md.district_name_t AS district_name,
				ms.subdistrict_name_t AS subdistrict_name,
				mp.postcodemain
		FROM 
				as_condos ac
		LEFT JOIN 
				as_lands al 
				ON al.id = ac.land_id 
		LEFT JOIN 
				property_doc_type pdt 
				ON al.parcel_type = pdt.doc_type_id
		LEFT JOIN 
				master_district md 
				ON md.district_code = ac.district_id
		LEFT JOIN 
				master_subdistrict ms 
				ON ms.subdistrict_code = ac.subdistrict_id
		LEFT JOIN (
				SELECT 
						DISTINCT ON (districtid) districtid, 
						postcodemain 
				FROM 
						master_postcode
		) mp
				ON md.district_code = mp.districtid
		WHERE 
				ac.id = ?;
		`

	var condoData map[string]interface{}
	if err := db.Raw(condoSQL, condoId).Scan(&condoData).Error; err != nil {
		return nil, fmt.Errorf("error getting condo details: %w", err)
	}

	assetImages := GetAssetAttachmentLatestSurveyByAssetId(condoId, condoId)

	condoData["asset_images"] = assetImages

	return map[string]interface{}{
		"asset_type": "condo",
		"condo_data": condoData,
	}, nil
}

func GetAssetAttachmentLatestSurveyByAssetId(assetID string, landID string) []asset_types.AssetAttachment {
	var db = database.GetDB()
	query := `
		SELECT DISTINCT ON (aa.img_id)
				aa.img_id,
				aa.key,
				aa.image_from,
				aa.survey_request_id
		FROM asset_attachment aa
		WHERE aa.asset_id = ?
			AND (
				-- Check if no matching survey request exists
				(NOT EXISTS (
					SELECT 1
					FROM survey_request sr
					WHERE sr.land_id = ?
				) AND aa.image_from = 3)
				OR
				-- Match survey_request_id to the latest survey_request
				(aa.survey_request_id = (
					SELECT sr.id
					FROM survey_request sr
					WHERE sr.land_id = ?
					ORDER BY sr.created_at DESC
					LIMIT 1
				))
			)
		ORDER BY aa.img_id, aa.created_date;`

	var attachments []asset_types.AssetAttachment

	err := db.Raw(
		query,
		assetID,
		landID,
		landID,
	).Scan(&attachments).Error

	if err != nil {
		// Log the error but return an empty slice as per original function
		fmt.Printf("Error fetching asset attachments: %v\n", err)
		return []asset_types.AssetAttachment{}
	}

	if len(attachments) == 0 {
		return []asset_types.AssetAttachment{}
	}

	return attachments
}

func GetAssetByLandId(landId string) (map[string]interface{}, error) {

	latestSurveyRequest := survey_request_survices.FindLatestSurveyRequestByAssetId(landId, "", "")

	getLandInfoDetail := GetLandInfoDetailByLandId(landId)

	muniCode, ok := getLandInfoDetail["muni_code"].(string)
	if !ok {
		return nil, errors.New("muni_code is missing or invalid in land info details")
	}

	getLandOwners := GetLandOwnersByLandId(landId, muniCode)
	getLandAssets := GetAssetAttachmentLatestSurveyByAssetId(landId, landId)
	getAdjoiningLands := GetAdjoiningLandDetailListByLandID(landId, muniCode)

	getLandUsedsInfo := GetLandUsedsInfo(landId, muniCode)

	getSignboardInfo := GetAllSignboardsOnLandByLandID(landId, muniCode)

	landInfo := map[string]interface{}{
		"land_info_detail": getLandInfoDetail,
		"land_owners":      getLandOwners,
		"asset_images":     getLandAssets,
		"adjoining_lands":  getAdjoiningLands,
	}

	response := map[string]interface{}{
		"latest_survey_request": latestSurveyRequest,
		"land_info":             landInfo,
		"land_used_info":        getLandUsedsInfo,
		"signboard_info":        getSignboardInfo,
	}

	return response, nil
}

func GetAssetByCondoId(condoId string) (map[string]interface{}, error) {

	latestSurveyRequest := survey_request_survices.FindLatestSurveyRequestByAssetId(condoId, "", "")

	response := map[string]interface{}{
		"asset_type":            "condo",
		"latest_survey_request": latestSurveyRequest,
	}

	return response, nil
}

func GetLandInfoDetailByLandId(landId string) map[string]interface{} {
	var db = database.GetDB()
	query := `
		SELECT 
			al.id,
			al.parcel_type,
			al.land_no,
			al.deed_no,
			al.survey_no, 
			al.parcel_no,
			al.utm_map1,
			al.utm_map2,
			al.utm_map3,
			al.utm_map4,
			al.utm_scale,
			al.land_space_rai,
			al.land_space_ngan,
			al.land_space_wa,
			al.land_space_sub_wa,
			al.map_lat,
			al.map_long,
			al.cutax_land_id,
			al.sync_id,
			al.created_at,
			al.created_by,
			al.updated_at,
			al.updated_by,
			al.deleted_at,
			al.deleted_by,
			al.tax_year,
			al.muni_code,
			al.land_zone,
			al.land_district_id,
			al.land_sub_district_id,
			al.estimated_price,
			al.estimated_price_by_treasury,
			al.is_estimated_by_treasury,
			ST_AsGeoJSON(al.map_geometry) as map_geometry,
			pdt.doc_type_name,
			mm.report_dt,
			mm.municipality_name_t,
			mm.province_code,
			mm.province_name_t,
			mm.district_code,
			mm.district_name_t,
			mm.sub_district_code,
			ms.subdistrict_name_t,
			al.note,
			al.special_usetax_type,
			al.road
		FROM as_lands al
		LEFT JOIN master_municipality mm ON mm.municipality_code = al.muni_code 
		LEFT JOIN master_subdistrict ms ON ms.subdistrict_code = mm.sub_district_code 
		LEFT JOIN property_doc_type pdt ON pdt.doc_type_id = al.parcel_type
		WHERE al.id = ? 
		LIMIT 1
	`

	// Prepare to store the result as []map[string]interface{}
	var landInfo = []map[string]interface{}{}

	// Execute the raw SQL query
	result := db.Raw(query, landId).Scan(&landInfo)

	if result.Error != nil {
		fmt.Println("Error executing query:", result.Error)
		return map[string]interface{}{}
	}

	// If no results found, return an empty slice
	if len(landInfo) == 0 {
		return map[string]interface{}{}
	}

	return landInfo[0]
}

func GetLandOwnersByLandId(landID string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()

	query := `
        SELECT
          alo.id,
          c.id owner_id,
          c.sync_id,
          c.person_type_id,
          c.tax_id,
          c.prefix_name_id,
          c.first_name,
          c.last_name,
          c.phone_number,
          c.line_id,
          c.email,
          c.address_house_number,
          c.address_zone,
          c.address_street,
          c.address_alleyway,
          (
            CASE
              WHEN length(c.address_province_code) = 2 THEN concat(mp.region_code, c.address_province_code)
              ELSE c.address_province_code
            END 
          )AS address_province_code,
          (
            CASE
              WHEN length(c.address_sub_district_code) = 6 THEN concat(mp.region_code, c.address_sub_district_code)
              ELSE c.address_sub_district_code
            END 
          )AS address_sub_district_code,
          (
            CASE
              WHEN length(c.address_district_code) = 4 THEN concat(mp.region_code, c.address_district_code)
              ELSE c.address_district_code
            END 
          )AS address_district_code,
          c.address_postcode,
          c.created_at,
          c.updated_at,
          c.corporate_name,
          c.codept4,
          c.fax_no,
          c.muni_code,
          c.date_source,
          c.citizen_id_check,
          c.code_name,
          c.current_address_house_number,
          c.current_address_zone,
          c.current_address_street,
          c.current_address_alleyway,
          (
            CASE
              WHEN length(c.current_address_province_code) = 2 THEN concat(mp.region_code, c.current_address_province_code)
              ELSE c.current_address_province_code
            END 
          )AS current_address_province_code,
          (
            CASE
              WHEN length(c.current_address_sub_district_code) = 6 THEN concat(mp.region_code, c.current_address_sub_district_code)
              ELSE c.current_address_sub_district_code
            END 
          )AS current_address_sub_district_code,
          (
            CASE
              WHEN length(c.current_address_district_code) = 4 THEN concat(mp.region_code, c.current_address_district_code)
              ELSE c.current_address_district_code
            END 
          )AS current_address_district_code,
          c.current_address_postcode,
          c.is_same_address,
          pt.person_type_name,
          pn."name" AS prefix_name,
          ms.province_name_t AS province_name,
          ms.district_name_t AS district_name,
          ms.subdistrict_name_t AS sub_district_name,
          alo.owner_line_no,
          (
          trim(
            CASE
              WHEN pt.person_type_id = '1' THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN pt.person_type_id = '99' THEN c.corporate_name
              ELSE concat(pn."name", c.corporate_name)
            END
          )
          ) AS text_full_name
        FROM citizen c 
        LEFT JOIN as_land_owners alo ON alo.owner_id = c.id AND alo.muni_code = ?
        LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id 
        LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
        LEFT JOIN master_province mp ON
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_province_code) = 2 THEN mp.ad_province = c.address_province_code::NUMERIC 
                ELSE mp.province_code = c.address_province_code
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_province_code) = 2 THEN mp.ad_province = concat(mp.region_code , c.current_address_province_code)::NUMERIC 
                ELSE mp.province_code = c.current_address_province_code
              END
          END
        LEFT JOIN master_district md ON 
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_district_code) = 4 THEN md.district_code = concat(md.region_code, c.address_district_code)
                ELSE md.district_code = c.address_district_code 
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_district_code) = 4 THEN md.district_code = concat(md.region_code , c.current_address_district_code)
                ELSE md.district_code = c.current_address_district_code
              END
          END
        LEFT JOIN master_subdistrict ms ON 
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.address_sub_district_code)
                ELSE ms.subdistrict_code = c.address_sub_district_code
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.current_address_sub_district_code)
                ELSE ms.subdistrict_code = c.current_address_sub_district_code
              END
          END
        WHERE alo.land_id = ?
          AND (alo.deleted_at IS NULL AND alo.deleted_by IS NULL)
          AND alo.muni_code = ?
          AND c.muni_code = ?
        ORDER BY alo.owner_line_no NULLS LAST
    `
	var landOwners = []map[string]interface{}{}
	result := db.Raw(query, muniCode, landID, muniCode, muniCode).Scan(&landOwners)
	if result.Error != nil {
		return []map[string]interface{}{}
	}

	return landOwners
}

func GetLandUsedOwnersByLandUsedId(landUsedId string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()

	query := `
        SELECT
          aluo.id,
          c.id owner_id,
          c.sync_id,
          c.person_type_id,
          c.tax_id,
          c.prefix_name_id,
          c.first_name,
          c.last_name,
          c.phone_number,
          c.line_id,
          c.email,
          c.address_house_number,
          c.address_zone,
          c.address_street,
          c.address_alleyway,
          (
            CASE
              WHEN length(c.address_province_code) = 2 THEN concat(mp.region_code, c.address_province_code)
              ELSE c.address_province_code
            END 
          )AS address_province_code,
          (
            CASE
              WHEN length(c.address_sub_district_code) = 6 THEN concat(mp.region_code, c.address_sub_district_code)
              ELSE c.address_sub_district_code
            END 
          )AS address_sub_district_code,
          (
            CASE
              WHEN length(c.address_district_code) = 4 THEN concat(mp.region_code, c.address_district_code)
              ELSE c.address_district_code
            END 
          )AS address_district_code,
          c.address_postcode,
          c.created_at,
          c.updated_at,
          c.corporate_name,
          c.codept4,
          c.fax_no,
          c.muni_code,
          c.date_source,
          c.citizen_id_check,
          c.code_name,
          c.current_address_house_number,
          c.current_address_zone,
          c.current_address_street,
          c.current_address_alleyway,
          (
            CASE
              WHEN length(c.current_address_province_code) = 2 THEN concat(mp.region_code, c.current_address_province_code)
              ELSE c.current_address_province_code
            END 
          )AS current_address_province_code,
          (
            CASE
              WHEN length(c.current_address_sub_district_code) = 6 THEN concat(mp.region_code, c.current_address_sub_district_code)
              ELSE c.current_address_sub_district_code
            END 
          )AS current_address_sub_district_code,
          (
            CASE
              WHEN length(c.current_address_district_code) = 4 THEN concat(mp.region_code, c.current_address_district_code)
              ELSE c.current_address_district_code
            END 
          )AS current_address_district_code,
          c.current_address_postcode,
          c.is_same_address,
          pt.person_type_name,
          pn."name" AS prefix_name,
          ms.province_name_t AS province_name,
          ms.district_name_t AS district_name,
          ms.subdistrict_name_t AS sub_district_name,
          aluo.owner_line_no,
          (
          trim(
            CASE
              WHEN pt.person_type_id = '1' THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN pt.person_type_id = '99' THEN c.corporate_name
              ELSE concat(pn."name", c.corporate_name)
            END
          )
          ) AS text_full_name
        FROM citizen c 
        LEFT JOIN as_land_used_owners aluo ON aluo.owner_id = c.id AND aluo.muni_code = ?
        LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id 
        LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
        LEFT JOIN master_province mp ON
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_province_code) = 2 THEN mp.ad_province = c.address_province_code::NUMERIC 
                ELSE mp.province_code = c.address_province_code
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_province_code) = 2 THEN mp.ad_province = concat(mp.region_code , c.current_address_province_code)::NUMERIC 
                ELSE mp.province_code = c.current_address_province_code
              END
          END
        LEFT JOIN master_district md ON 
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_district_code) = 4 THEN md.district_code = concat(md.region_code, c.address_district_code)
                ELSE md.district_code = c.address_district_code 
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_district_code) = 4 THEN md.district_code = concat(md.region_code , c.current_address_district_code)
                ELSE md.district_code = c.current_address_district_code
              END
          END
        LEFT JOIN master_subdistrict ms ON 
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.address_sub_district_code)
                ELSE ms.subdistrict_code = c.address_sub_district_code
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.current_address_sub_district_code)
                ELSE ms.subdistrict_code = c.current_address_sub_district_code
              END
          END
        WHERE aluo.land_used_id = ?
          AND (aluo.deleted_at IS NULL AND aluo.deleted_by IS NULL)
          AND aluo.muni_code = ?
          AND c.muni_code = ?
        ORDER BY aluo.owner_line_no NULLS LAST
    `
	var landOwners = []map[string]interface{}{}
	result := db.Raw(query, muniCode, landUsedId, muniCode, muniCode).Scan(&landOwners)
	if result.Error != nil {
		fmt.Println("Error fetching data:", result.Error)
		return []map[string]interface{}{}
	}

	return landOwners
}

func GetBuildingOwnersByBuildingId(buildingId string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()

	query := `
        SELECT
          abo.id,
          c.id owner_id,
          c.sync_id,
          c.person_type_id,
          c.tax_id,
          c.prefix_name_id,
          c.first_name,
          c.last_name,
          c.phone_number,
          c.line_id,
          c.email,
          c.address_house_number,
          c.address_zone,
          c.address_street,
          c.address_alleyway,
          (
            CASE
              WHEN length(c.address_province_code) = 2 THEN concat(mp.region_code, c.address_province_code)
              ELSE c.address_province_code
            END 
          )AS address_province_code,
          (
            CASE
              WHEN length(c.address_sub_district_code) = 6 THEN concat(mp.region_code, c.address_sub_district_code)
              ELSE c.address_sub_district_code
            END 
          )AS address_sub_district_code,
          (
            CASE
              WHEN length(c.address_district_code) = 4 THEN concat(mp.region_code, c.address_district_code)
              ELSE c.address_district_code
            END 
          )AS address_district_code,
          c.address_postcode,
          c.created_at,
          c.updated_at,
          c.corporate_name,
          c.codept4,
          c.fax_no,
          c.muni_code,
          c.date_source,
          c.citizen_id_check,
          c.code_name,
          c.current_address_house_number,
          c.current_address_zone,
          c.current_address_street,
          c.current_address_alleyway,
          (
            CASE
              WHEN length(c.current_address_province_code) = 2 THEN concat(mp.region_code, c.current_address_province_code)
              ELSE c.current_address_province_code
            END 
          )AS current_address_province_code,
          (
            CASE
              WHEN length(c.current_address_sub_district_code) = 6 THEN concat(mp.region_code, c.current_address_sub_district_code)
              ELSE c.current_address_sub_district_code
            END 
          )AS current_address_sub_district_code,
          (
            CASE
              WHEN length(c.current_address_district_code) = 4 THEN concat(mp.region_code, c.current_address_district_code)
              ELSE c.current_address_district_code
            END 
          )AS current_address_district_code,
          c.current_address_postcode,
          c.is_same_address,
          pt.person_type_name,
          pn."name" AS prefix_name,
          ms.province_name_t AS province_name,
          ms.district_name_t AS district_name,
          ms.subdistrict_name_t AS sub_district_name,
          abo.owner_line_no,
          (
          trim(
            CASE
              WHEN pt.person_type_id = '1' THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN pt.person_type_id = '99' THEN c.corporate_name
              ELSE concat(pn."name", c.corporate_name)
            END
          )
          ) AS text_full_name
        FROM citizen c 
        LEFT JOIN as_building_owners abo ON abo.owner_id = c.id AND abo.muni_code = ?
        LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id 
        LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
        LEFT JOIN master_province mp ON
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_province_code) = 2 THEN mp.ad_province = c.address_province_code::NUMERIC 
                ELSE mp.province_code = c.address_province_code
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_province_code) = 2 THEN mp.ad_province = concat(mp.region_code , c.current_address_province_code)::NUMERIC 
                ELSE mp.province_code = c.current_address_province_code
              END
          END
        LEFT JOIN master_district md ON 
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_district_code) = 4 THEN md.district_code = concat(md.region_code, c.address_district_code)
                ELSE md.district_code = c.address_district_code 
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_district_code) = 4 THEN md.district_code = concat(md.region_code , c.current_address_district_code)
                ELSE md.district_code = c.current_address_district_code
              END
          END
        LEFT JOIN master_subdistrict ms ON 
          CASE
            WHEN COALESCE(c.is_same_address, true) THEN
              CASE
                WHEN length(c.address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.address_sub_district_code)
                ELSE ms.subdistrict_code = c.address_sub_district_code
              END
            ELSE 
              CASE 
                WHEN length(c.current_address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.current_address_sub_district_code)
                ELSE ms.subdistrict_code = c.current_address_sub_district_code
              END
          END
        WHERE abo.building_id = ?
          AND (abo.deleted_at IS NULL AND abo.deleted_by IS NULL)
          AND abo.muni_code = ?
          AND c.muni_code = ?
        ORDER BY abo.owner_line_no NULLS LAST
    `
	var landOwners = []map[string]interface{}{}
	result := db.Raw(query, muniCode, buildingId, muniCode, muniCode).Scan(&landOwners)
	if result.Error != nil {
		fmt.Println("Error fetching data:", result.Error)
		return []map[string]interface{}{}
	}

	return landOwners
}
func GetBuildingUsedByBuildingUsedId(buildingUsedId string) []map[string]interface{} {
	var db = database.GetDB()

	query := `
        SELECT 
						abu.*, 
						prt.rent_type_detail, 
						put.using_type_detail, 
						pht.household_type_name
				FROM 
						as_building_used abu
				LEFT JOIN 
						property_using_type put 
						ON abu.building_using_type_id = put.using_type_id
				LEFT JOIN 
						property_rent_type prt 
						ON abu.property_rent_type_id = prt.rent_type_id
				LEFT JOIN 
						property_household_type pht 
						ON abu.building_household_type_id = pht.household_type_id
				WHERE 
						abu.id = ?;
    `
	var buildingUsed = []map[string]interface{}{}
	result := db.Raw(query, buildingUsedId).Scan(&buildingUsed)
	if result.Error != nil {
		fmt.Println("Error fetching data:", result.Error)
		return []map[string]interface{}{}
	}

	return buildingUsed
}

func GetAdjoiningLandDetailListByLandID(landID string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()
	query := `
		WITH adjoining_lands AS (
          SELECT *
          FROM as_adjoining_lands dal
          WHERE dal.land_id = ?
            AND (dal.deleted_at IS NULL AND dal.deleted_by IS NULL)
            AND dal.muni_code = ?
        ),
        main_land_owner AS (
          SELECT
            alo2.owner_id AS main_owner_id
          FROM as_land_owners alo2
          WHERE alo2.land_id = ?
            AND alo2.owner_line_no = '1'
            AND (alo2.deleted_at IS NULL AND alo2.deleted_by IS NULL)
            AND alo2.muni_code = ?
        ),
        adjoining_lands_survey AS (
          SELECT
            sr.status,
            sr.land_id,
            sr.survey_code,
            ROW_NUMBER() OVER (PARTITION BY sr.land_id ORDER BY sr.created_at DESC) AS created_at_order
          FROM survey_request sr
          WHERE sr.land_id IN (SELECT adl.adjoining_land_id FROM adjoining_lands adl)
            AND sr.muni_code = ?
        ),
        merge_adjoining_lands_survey AS (
          SELECT
            adjl.*,
            CASE
              WHEN adjs.status IS NULL OR adjs.status IN ('7', '9') THEN FALSE
              ELSE true
            END AS is_survey
          FROM adjoining_lands adjl
          LEFT JOIN adjoining_lands_survey adjs ON adjs.land_id = adjl.adjoining_land_id
            AND adjs.created_at_order = '1' OR NULL
        ),
        final_result AS (
          SELECT
            madjs.id,
            madjs.land_id,
            madjs.adjoining_land_id,
            madjs.status,
            madjs.is_survey,
            CASE
                WHEN mlo.main_owner_id = COALESCE(dlo.owner_id, alo.owner_id) THEN TRUE
                ELSE FALSE
            END AS is_same_owner,
            COALESCE(dl.parcel_no, al.parcel_no) AS parcel_no,
            COALESCE(dl.deed_no, al.deed_no) AS deed_no,
            COALESCE(dl.land_space_rai, al.land_space_rai) AS land_space_rai,
            COALESCE(dl.land_space_ngan, al.land_space_ngan) AS land_space_ngan,
            COALESCE(dl.land_space_wa, al.land_space_wa) AS land_space_wa,
            COALESCE(dlo.owner_id, alo.owner_id) AS owner_id
          FROM merge_adjoining_lands_survey madjs
          LEFT JOIN df_lands dl ON dl.id = madjs.adjoining_land_id AND madjs.is_survey AND dl.muni_code = ?
          LEFT JOIN df_land_owners dlo ON dlo.land_id = dl.id
            AND madjs.is_survey
            AND dlo.owner_line_no = '1'
            AND (dlo.deleted_at IS NULL AND dlo.deleted_by IS NULL)
            AND dlo.muni_code = ?
          LEFT JOIN as_lands al ON al.id = madjs.adjoining_land_id
          	AND NOT madjs.is_survey
          	AND al.muni_code = ?
          LEFT JOIN as_land_owners alo ON alo.land_id = al.id
            AND NOT madjs.is_survey
            AND alo.owner_line_no = '1'
            AND (alo.deleted_at IS NULL AND alo.deleted_by IS NULL)
            AND alo.muni_code = ?
          CROSS JOIN main_land_owner mlo
        )
        SELECT
          fr.*,
          trim(
            CASE
              WHEN pt.person_type_id = '1' THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN pt.person_type_id = '99' THEN c.corporate_name
              ELSE concat(pn."name", c.corporate_name)
            END
          ) AS full_name,
          c.tax_id,
          c.phone_number,
          pt.person_type_id,
          pt.person_type_name
        FROM final_result fr
        LEFT JOIN citizen c ON c.id = fr.owner_id AND c.muni_code = ?
        LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id
        LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id
	`
	var adJoiningLand = []map[string]interface{}{}
	result := db.Raw(query, landID, muniCode, landID, muniCode, muniCode, muniCode, muniCode, muniCode, muniCode, muniCode).Scan(&adJoiningLand)
	if result.Error != nil {
		return []map[string]interface{}{}
	}

	return adJoiningLand
}

func GetLandUsedsWithoutBuildingOnLandByLandId(landID string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()
	var allLandUsed = []map[string]interface{}{}
	var result = []map[string]interface{}{}

	// Fetch the data using GORM
	err := db.Raw(`
       SELECT
          alu.*, put.using_type_detail,
					case 
          	when alu.tax_deduction_id is null then 'ไม่เลือก'
          	else mtd.description 
          end as tax_deduction_detail 
        FROM as_land_used alu
        LEFT JOIN property_using_type put on alu.using_type_id = put.using_type_id 
				LEFT JOIN master_tax_deduction mtd on alu.tax_deduction_id = mtd.id 
        WHERE alu.land_id =?
          AND (alu.deleted_at IS NULL AND alu.deleted_by IS NULL)
          AND alu.muni_code = ?
        ORDER BY alu.cutax_landused_id ASC
	`, landID, muniCode).Scan(&allLandUsed).Error

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return []map[string]interface{}{}
	}

	if len(allLandUsed) == 0 {
		return []map[string]interface{}{}
	}

	// Transform the results
	for _, record := range allLandUsed {
		var usingDetailID interface{}

		if usingDetailIDStr, ok := record["using_detail_id"].(string); ok && usingDetailIDStr != "" {
			// Try to parse as JSON first
			var jsonData struct {
				ID []string `json:"id"`
			}

			if err := json.Unmarshal([]byte(usingDetailIDStr), &jsonData); err != nil {
				// If JSON parsing fails, use the alternative parsing method
				usingDetailID = []string{}
				idStr := ""

				for _, char := range usingDetailIDStr {
					if char != '[' && char != ']' && char != '\'' {
						idStr += string(char)
					}
				}

				if idStr != "" {
					usingDetailID = strings.Split(strings.TrimSpace(idStr), ",")
				}
			} else {
				usingDetailID = jsonData.ID
			}
		} else {
			usingDetailID = []string{}
		}

		// Create new record with transformed using_detail_id
		newRecord := make(map[string]interface{})
		for key, value := range record {
			newRecord[key] = value
		}
		newRecord["using_detail_id"] = usingDetailID

		result = append(result, newRecord)
	}

	return result
}

func GetLandUsedBuildingByLandUsedId(landUsedID string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()
	var allBuildings = []map[string]interface{}{}

	// Fetch the data using GORM
	err := db.Raw(`
		SELECT
          ab.*,
          abu.id AS building_used_id,
          c.first_name,
          c.last_name,
          c.prefix_name_id,
          c.person_type_id,
          c.corporate_name,
          pn."name" AS prefix_name,
          pt.person_type_name,
          pbmt.building_type_name as building_main_type_name,
          abo.owner_line_no,
          pt.person_type_name,
					pbst.building_type_name as building_sub_type_name,
					pbdt.building_design_type_name,
          trim(
            concat(
              CASE
                WHEN pt.person_type_id IS NOT NULL THEN concat('(', pt.person_type_name, ') ')
                ELSE ''
              END,
              CASE
                WHEN c.person_type_id = 1 THEN concat(pn."name", c.first_name, ' ', c.last_name)
                WHEN c.person_type_id = 99 THEN c.corporate_name
                ELSE concat(pn."name", ' ', c.corporate_name)
              END
          )) AS text_owner_prefix_fullname,
          trim(
            CASE
              WHEN c.person_type_id = 1 THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN c.person_type_id = 99 THEN c.corporate_name
              ELSE concat(pn."name", ' ', c.corporate_name)
            END
          ) AS text_owner_fullname
        FROM as_buildings ab
          LEFT JOIN as_building_owners abo ON abo.building_id = ab.id AND abo.owner_line_no = '1' AND abo.muni_code = ?
          LEFT JOIN citizen c ON c.id = abo.owner_id AND c.muni_code = ?
          LEFT JOIN as_building_used abu ON abu.building_id = ab.id AND abu.muni_code = ?
          LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id
          LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
          LEFT JOIN property_building_main_type pbmt ON pbmt.building_type_id = ab.property_building_main_type_id
          	AND pbmt.muni_code = ?
          	AND pbmt.building_type_year = ab.building_type_year
					LEFT JOIN property_building_sub_type pbst on pbst.building_sub_id = ab.property_building_sub_type_id 
         AND pbst.muni_code = ?
         AND pbst.building_type_year = ab.building_type_year
				 LEFT JOIN property_building_design_type pbdt on pbdt.building_design_type_id = ab.building_design_type_id
        WHERE ab.land_used_id = ?
          AND (ab.deleted_at IS NULL AND ab.deleted_by IS NULL)
          AND ab.muni_code = ?
        ORDER BY ab.cutax_building_id ASC
	`, muniCode, muniCode, muniCode, muniCode, muniCode, landUsedID, muniCode).Scan(&allBuildings).Error

	// fmt.Println("allBuildings ***>>>", allBuildings)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return []map[string]interface{}{}
	}

	if len(allBuildings) == 0 {
		return []map[string]interface{}{}
	}

	// Transform the result
	for i := 0; i < len(allBuildings); i++ {
		buildingId, ok := allBuildings[i]["id"].(string)
		if ok {
			var allBuildOwners = GetBuildingOwnersByBuildingId(buildingId, muniCode)
			// Add building owners to the data
			allBuildings[i]["building_owners"] = allBuildOwners

			var buildingUsed = []map[string]interface{}{}
			var buildingUsedId, ok = allBuildings[i]["building_used_id"].(string)
			if ok {
				buildingUsed = GetBuildingUsedByBuildingUsedId(buildingUsedId)
			}
			if len(buildingUsed) > 0 {
				allBuildings[i]["building_used"] = buildingUsed[0]
			} else {
				allBuildings[i]["building_used"] = []map[string]interface{}{}
			}
		}

	}

	// Return the transformed result
	return allBuildings
}

func GetLandUsedsInfo(landID string, muniCode string) []map[string]interface{} {
	allLandUsed := GetLandUsedsWithoutBuildingOnLandByLandId(landID, muniCode)

	if len(allLandUsed) == 0 {
		return []map[string]interface{}{}
	}

	var wg sync.WaitGroup
	var ownersMutex, imagesMutex sync.Mutex

	// Fetch land used owners and citizens
	fetchLandUsedOwners := func(i int) {
		defer wg.Done()

		landUsedId, ok := allLandUsed[i]["id"].(string)
		if ok {
			findAllAsLandUsedOwners := GetLandUsedOwnersByLandUsedId(landUsedId, muniCode)

			ownersMutex.Lock()
			allLandUsed[i]["land_used_owners"] = findAllAsLandUsedOwners
			ownersMutex.Unlock()
		}
	}

	// Fetch building data first, synchronously for each land use
	for i := range allLandUsed {
		landUsedid, ok := allLandUsed[i]["id"].(string)
		if ok {
			buildings := GetLandUsedBuildingByLandUsedId(landUsedid, muniCode)
			allLandUsed[i]["buildings"] = buildings
		}
	}

	// Fetch images for land used and buildings
	fetchImages := func(i int) {
		defer wg.Done()

		// Fetch land used images
		landUsedId, ok := allLandUsed[i]["id"].(string)
		if ok {
			imagesLandUsed := GetAssetAttachmentLatestSurveyByAssetId(landUsedId, landID)
			imagesMutex.Lock()
			allLandUsed[i]["asset_images"] = imagesLandUsed
			imagesMutex.Unlock()
		}

		// Fetch building images
		if buildings, ok := allLandUsed[i]["buildings"].([]map[string]interface{}); ok {
			for j, building := range buildings {
				if buildingId, ok := building["id"].(string); ok {
					buildingImages := GetAssetAttachmentLatestSurveyByAssetId(buildingId, landID)

					imagesMutex.Lock()
					allLandUsed[i]["buildings"].([]map[string]interface{})[j]["asset_images"] = buildingImages
					imagesMutex.Unlock()
				}
			}
		}
	}

	// Start goroutines for owners and images
	for i := range allLandUsed {
		wg.Add(2) // Two goroutines per land used: owners and images
		go fetchLandUsedOwners(i)
		go fetchImages(i)
	}
	wg.Wait()

	return allLandUsed
}

// Fetch citizens for each land used owner
func FetchCitizens(ownerID interface{}) []map[string]interface{} {
	var db = database.GetDB()
	var citizens []map[string]interface{}
	citizenQuery := `
		SELECT
									c.id AS owner_id,
									c.sync_id,
									c.person_type_id,
									c.tax_id,
									c.prefix_name_id,
									c.first_name,
									c.last_name,
									c.phone_number,
									c.line_id,
									c.email,
									c.address_house_number,
									c.address_zone,
									c.address_street,
									c.address_alleyway,
									(
										CASE
											WHEN length(c.address_province_code) = 2 THEN concat(mp.region_code, c.address_province_code)
											ELSE c.address_province_code
										END
									)AS address_province_code,
									(
										CASE
											WHEN length(c.address_sub_district_code) = 6 THEN concat(mp.region_code, c.address_sub_district_code)
											ELSE c.address_sub_district_code
										END
									)AS address_sub_district_code,
									(
										CASE
											WHEN length(c.address_district_code) = 4 THEN concat(mp.region_code, c.address_district_code)
											ELSE c.address_district_code
										END
									)AS address_district_code,
									c.address_postcode,
									c.created_at,
									c.updated_at,
									c.corporate_name,
									c.codept4,
									c.fax_no,
									c.muni_code,
									c.date_source,
									c.citizen_id_check,
									c.code_name,
									c.current_address_house_number,
									c.current_address_zone,
									c.current_address_street,
									c.current_address_alleyway,
									(
										CASE
											WHEN length(c.current_address_province_code) = 2 THEN concat(mp.region_code, c.current_address_province_code)
											ELSE c.current_address_province_code
										END
									)AS current_address_province_code,
									(
										CASE
											WHEN length(c.current_address_sub_district_code) = 6 THEN concat(mp.region_code, c.current_address_sub_district_code)
											ELSE c.current_address_sub_district_code
										END
									)AS current_address_sub_district_code,
									(
										CASE
											WHEN length(c.current_address_district_code) = 4 THEN concat(mp.region_code, c.current_address_district_code)
											ELSE c.current_address_district_code
										END
									)AS current_address_district_code,
									c.current_address_postcode,
									c.is_same_address,
									pt.person_type_name,
									pn."name" AS prefix_name,
									ms.province_name_t AS province_name,
									ms.district_name_t AS district_name,
									ms.subdistrict_name_t AS sub_district_name,
									(
									trim(
										CASE
											WHEN pt.person_type_id = '1' THEN concat(pn."name", c.first_name, ' ', c.last_name)
											WHEN pt.person_type_id = '99' THEN c.corporate_name
											ELSE concat(pn."name", c.corporate_name)
										END
									)
									) AS text_full_name
								FROM citizen c
								LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id
								LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
								LEFT JOIN master_province mp ON
									CASE
										WHEN c.is_same_address THEN
											CASE
												WHEN length(c.address_province_code) = 2 THEN mp.ad_province = c.address_province_code::NUMERIC
												ELSE mp.province_code = c.address_province_code
											END
										ELSE
											CASE
												WHEN length(c.current_address_province_code) = 2 THEN mp.ad_province = concat(mp.region_code , c.current_address_province_code)::NUMERIC
												ELSE mp.province_code = c.current_address_province_code
											END
									END
								LEFT JOIN master_district md ON
									CASE
										WHEN c.is_same_address THEN
											CASE
												WHEN length(c.address_district_code) = 4 THEN md.district_code = concat(md.region_code, c.address_district_code)
												ELSE md.district_code = c.address_district_code
											END
										ELSE
											CASE
												WHEN length(c.current_address_district_code) = 4 THEN md.district_code = concat(md.region_code , c.current_address_district_code)
												ELSE md.district_code = c.current_address_district_code
											END
									END
								LEFT JOIN master_subdistrict ms ON
									CASE
										WHEN c.is_same_address THEN
											CASE
												WHEN length(c.address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.address_sub_district_code)
												ELSE ms.subdistrict_code = c.address_sub_district_code
											END
										ELSE
											CASE
												WHEN length(c.current_address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.current_address_sub_district_code)
												ELSE ms.subdistrict_code = c.current_address_sub_district_code
											END
									end
								where c.id = ?`
	err := db.Raw(citizenQuery, ownerID).Scan(&citizens).Error
	if err != nil {
		fmt.Println("Error fetching citizen data:", err)
	}
	return citizens
}

func GetSignboardOwnersBySignboardId(signboardID string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()
	query := `
        SELECT
          aso.id,
          c.id AS owner_id,
          c.sync_id,
          c.person_type_id,
          c.tax_id,
          c.prefix_name_id,
          c.first_name,
          c.last_name,
          c.phone_number,
          c.line_id,
          c.email,
          c.address_house_number,
          c.address_zone,
          c.address_street,
          c.address_alleyway,
          (
            CASE
              WHEN length(c.address_province_code) = 2 THEN concat(mp.region_code, c.address_province_code)
              ELSE c.address_province_code
            END 
          )AS address_province_code,
          (
            CASE
              WHEN length(c.address_sub_district_code) = 6 THEN concat(mp.region_code, c.address_sub_district_code)
              ELSE c.address_sub_district_code
            END 
          )AS address_sub_district_code,
          (
            CASE
              WHEN length(c.address_district_code) = 4 THEN concat(mp.region_code, c.address_district_code)
              ELSE c.address_district_code
            END 
          )AS address_district_code,
          c.address_postcode,
          c.created_at,
          c.updated_at,
          c.corporate_name,
          c.codept4,
          c.fax_no,
          c.muni_code,
          c.date_source,
          c.citizen_id_check,
          c.code_name,
          c.current_address_house_number,
          c.current_address_zone,
          c.current_address_street,
          c.current_address_alleyway,
          (
            CASE
              WHEN length(c.current_address_province_code) = 2 THEN concat(mp.region_code, c.current_address_province_code)
              ELSE c.current_address_province_code
            END 
          )AS current_address_province_code,
          (
            CASE
              WHEN length(c.current_address_sub_district_code) = 6 THEN concat(mp.region_code, c.current_address_sub_district_code)
              ELSE c.current_address_sub_district_code
            END 
          )AS current_address_sub_district_code,
          (
            CASE
              WHEN length(c.current_address_district_code) = 4 THEN concat(mp.region_code, c.current_address_district_code)
              ELSE c.current_address_district_code
            END 
          )AS current_address_district_code,
          c.current_address_postcode,
          c.is_same_address,
          pt.person_type_name,
          pn."name" AS prefix_name,
          ms.province_name_t AS province_name,
          ms.district_name_t AS district_name,
          ms.subdistrict_name_t AS sub_district_name,
          aso.owner_line_no,
          (
          trim(
            CASE
              WHEN pt.person_type_id = '1' THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN pt.person_type_id = '99' THEN c.corporate_name
              ELSE concat(pn."name", c.corporate_name)
            END
          )
          ) AS text_full_name
        FROM citizen c 
          LEFT JOIN as_signboard_owners aso ON aso.owner_id = c.id AND aso.muni_code = ?
          LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id 
          LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
          LEFT JOIN master_province mp ON 
            CASE
              WHEN c.is_same_address THEN 
                CASE
                  WHEN length(c.address_province_code) = 2 THEN mp.ad_province = c.address_province_code::NUMERIC 
                  ELSE mp.province_code = c.address_province_code
                END
              ELSE 
                CASE 
                  WHEN length(c.current_address_province_code) = 2 THEN mp.ad_province = concat(mp.region_code , c.current_address_province_code)::NUMERIC 
                  ELSE mp.province_code = c.current_address_province_code
                END
            END
          LEFT JOIN master_district md ON 
            CASE
              WHEN c.is_same_address THEN 
                CASE
                  WHEN length(c.address_district_code) = 4 THEN md.district_code = concat(md.region_code, c.address_district_code)
                  ELSE md.district_code = c.address_district_code 
                END
              ELSE 
                CASE 
                  WHEN length(c.current_address_district_code) = 4 THEN md.district_code = concat(md.region_code , c.current_address_district_code)
                  ELSE md.district_code = c.current_address_district_code
                END
            END
          LEFT JOIN master_subdistrict ms ON 
            CASE
              WHEN c.is_same_address THEN 
                CASE
                  WHEN length(c.address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.address_sub_district_code)
                  ELSE ms.subdistrict_code = c.address_sub_district_code
                END
              ELSE 
                CASE 
                  WHEN length(c.current_address_sub_district_code) = 6 THEN ms.subdistrict_code = concat(md.region_code , c.current_address_sub_district_code)
                  ELSE ms.subdistrict_code = c.current_address_sub_district_code
                END
            END
        WHERE aso.signboard_id = ? 
          AND (aso.deleted_at IS NULL AND aso.deleted_by IS NULL)
          AND aso.muni_code = ?
          AND c.muni_code = ?
        ORDER BY aso.owner_line_no
      `
	var signboardOwners = []map[string]interface{}{}
	err := db.Raw(query, muniCode, signboardID, muniCode, muniCode).Scan(&signboardOwners).Error
	if err != nil {
		fmt.Println("error getting signboard owners: %w", err)
		return []map[string]interface{}{}
	}

	return signboardOwners
}

func GetAllSignboardsOnLandByLandID(landID string, muniCode string) []map[string]interface{} {
	var db = database.GetDB()
	var landSignboards = []map[string]interface{}{}

	query := `
			SELECT
          as2.id,
          as2.land_id,
          as2.signboard_text,
          as2.signboard_width,
          as2.signboard_height,
          as2.signboard_side,
          as2.signboard_total_area,
          as2.signboard_unit_total,
          as2.property_signboard_type_id,
          as2.property_signboard_display_type_id,
          as2.note,
          as2.map_lat,
          as2.map_long,
          as2.signboard_setup_date,
          as2.parcel_code,
          as2.address_place_number,
          as2.address_zone,
          as2.address_alley_way,
          as2.address_street,
          as2.address_sub_district_id,
          as2.muni_code,
          as2.utm_map1,
          as2.utm_map2,
          as2.utm_map3,
          as2.tax_year,
          as2.cutax_label_id,
          as2.sync_id,
          as2.created_at,
          as2.created_by,
          as2.updated_at,
          as2.updated_by,
          as2.deleted_at,
          as2.deleted_by,
          as2.cancel_status,
          as2.cancel_date,
          as2.cancel_by,
          as2.cancel_at,
					as2.address_village_name,
					as2.address_postcode ,
					as2.signboard_name ,
          pst.signboard_type_id,
          pst.signboard_type_name,
          psdt.property_signboard_display_type,
          psdt.property_signboard_display_name,
          psdt.property_signboard_display_desc,
          trim(concat(pst.signboard_type_id, substring(psdt.property_signboard_display_desc, 1, POSITION(' ' IN psdt.property_signboard_display_desc) - 1))) AS text_signboard_type,
          pt.person_type_name,
          trim(
            concat(
              CASE
                WHEN pt.person_type_id IS NOT NULL THEN concat('(', pt.person_type_name, ') ')
                ELSE ''
              END,
              CASE
                WHEN c.person_type_id = 1 THEN concat(pn."name", c.first_name, ' ', c.last_name)
                WHEN c.person_type_id = 99 THEN c.corporate_name
                ELSE concat(pn."name", ' ', c.corporate_name)
              END
          )) AS text_owner_prefix_fullname,
          trim(
            CASE
              WHEN c.person_type_id = 1 THEN concat(pn."name", c.first_name, ' ', c.last_name)
              WHEN c.person_type_id = 99 THEN c.corporate_name
              ELSE concat(pn."name", ' ', c.corporate_name)
            END
          ) AS text_owner_fullname,
          md.province_name_t AS province_name,
					md.district_name_t AS district_name,
					ms.subdistrict_name_t AS subdistrict_name
          FROM as_signboards as2
            LEFT JOIN property_signboard_type pst ON pst.signboard_type_id = as2.property_signboard_type_id
            LEFT JOIN property_signboard_display_type psdt ON psdt.property_signboard_display_type = as2.property_signboard_display_type_id
            LEFT JOIN as_signboard_owners aso ON aso.signboard_id = as2.id
              AND aso.owner_line_no = '1'
              AND aso.deleted_at IS NULL
              AND aso.deleted_by IS NULL
              AND aso.muni_code = ?
            LEFT JOIN citizen c ON c.id = aso.owner_id AND c.muni_code = ?
            LEFT JOIN person_type pt ON pt.person_type_id = c.person_type_id
            LEFT JOIN prefix_name pn ON pn.id = c.prefix_name_id AND pn.person_type = c.person_type_id
            LEFT JOIN master_district md ON md.district_code::numeric = as2.address_district_id 
            LEFT JOIN master_subdistrict ms ON as2.address_sub_district_id::text = ms.subdistrict_code
          WHERE as2.land_id = ?
            AND (as2.deleted_at IS NULL AND as2.deleted_by IS NULL)
            AND as2.muni_code = ?
          ORDER BY as2.cutax_label_id ASC`

	// Execute the query
	if err := db.Raw(query, muniCode, muniCode, landID, muniCode).Scan(&landSignboards).Error; err != nil {
		fmt.Println("failed to fetch land signboards: %w", err)
		return []map[string]interface{}{}
	}

	if len(landSignboards) == 0 {
		return []map[string]interface{}{}
	}

	// Fetch and merge images and owners
	for i, signboard := range landSignboards {
		signboardId := signboard["id"].(string)

		// Get images
		images := GetAssetAttachmentLatestSurveyByAssetId(signboardId, landID)
		landSignboards[i]["asset_images"] = images

		// Get owners
		owners := GetSignboardOwnersBySignboardId(signboardId, muniCode)
		landSignboards[i]["signboard_owners"] = owners
	}

	return landSignboards
}
