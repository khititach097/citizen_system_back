package users_services

import (
	// 	"citizen_system_back/database"
	// 	"citizen_system_back/models"
	// "citizen_system_back/models"
	asset_types "citizen_system_back/modules/users/types"
	"fmt"
)

func UpSertProfileCitizenUser(profile asset_types.Profile) {
	fmt.Println(" **************************** UpSertProfileCitizenUser profile", profile)
	fmt.Println(" **************************** UpSertProfileCitizenUser profile.id", profile.ID)

	// var db = database.GetDB()
	// var citizen []models.Citizen
	// var citizenUser []models.CitizenUser
	// err_citizen := db.Where("").Find(&citizen).Error
	// if err_citizen != nil {
	// 	fmt.Println("err_citizen : ", err_citizen)
	// 	return
	// }

	// err_citizenUser := db.Limit(10).Find(&citizenUser).Error
	// if err_citizenUser != nil {
	// 	fmt.Println("err_citizenUser : ", err_citizenUser)

	// 	return
	// }

	// return assets, err

}
