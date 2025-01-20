package users_services

import (
	"citizen_system_back/database"
	"citizen_system_back/models"

	// "citizen_system_back/models"
	asset_types "citizen_system_back/modules/users/types"
	"fmt"
)

// Mock response structure
type StatusResponse struct {
	StatusCode  string `json:"statusCode"`
	StatusDesc  string `json:"statusDesc"`
	MessageCode string `json:"messageCode"`
	MessageTH   string `json:"messageTH"`
	MessageEN   string `json:"messageEN"`
}

type Account struct {
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Account  string `json:"account"`
	Status   string `json:"status"`
}

type Data struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	PhoneNumber  string    `json:"phoneNumber"`
	Picture      *string   `json:"picture"`
	DocumentType string    `json:"documentType"`
	NationalID   string    `json:"nationalId"`
	PassportNo   *string   `json:"passportNo"`
	Status       string    `json:"status"`
	Accounts     []Account `json:"accounts"`
}

type AuthServiceResponse struct {
	StatusResponse StatusResponse `json:"statusResponse"`
	Data           Data           `json:"data"`
}

// func getUserByEmail() (*AuthServiceResponse, error) {
func GetUserByEmail() (*models.CitizenUser, error) {
	// 8fc7264f-e610-4541-934e-37bddcbaacf2

	var db = database.GetDB()
	var citizenUser *models.CitizenUser
	err := db.Find(&citizenUser).Where("id = ?", "8fc7264f-e610-4541-934e-37bddcbaacf2").Error
	fmt.Println("citizenUser : ", citizenUser)
	return citizenUser, err

	// Simulating successful response
	// mockResponse := &AuthServiceResponse{
	// 	StatusResponse: StatusResponse{
	// 		StatusCode:  "200",
	// 		StatusDesc:  "OK",
	// 		MessageCode: "00000",
	// 		MessageTH:   "สำเร็จ",
	// 		MessageEN:   "Success",
	// 	},
	// 	Data: Data{
	// 		ID:           "18Av6F8xTQZ1CqBJg6NoMMqkAZe2X",
	// 		Title:        "คุณหญ",
	// 		FirstName:    "First name",
	// 		LastName:     "Last name",
	// 		PhoneNumber:  "09000000000",
	// 		Picture:      nil,
	// 		DocumentType: "nationalId",
	// 		NationalID:   "1111111111111",
	// 		PassportNo:   nil,
	// 		Status:       "active",
	// 		Accounts: []Account{
	// 			{
	// 				Provider: "email",
	// 				Type:     "email",
	// 				Account:  "chanasuek.chk@gmail.com",
	// 				Status:   "active",
	// 			},
	// 			{
	// 				Provider: "email",
	// 				Type:     "phone_number",
	// 				Account:  "09000000000",
	// 				Status:   "active",
	// 			},
	// 		},
	// 	},
	// }

	// Return the mock response
	// return mockResponse, nil
}

func UpSertProfileCitizenUser(profile asset_types.Profile) {
	fmt.Println(" **************************** UpSertProfileCitizenUser profile", profile)
	fmt.Println(" **************************** UpSertProfileCitizenUser profile.id", profile.ID)

	// mock data
	// response, err := getUserByEmail()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("response:", response)

	// var db = database.GetDB()
	// var citizen []models.Citizen
	// // var citizenUser []models.CitizenUser
	// err_citizen := db.Where("tax_id = ?", response.Data.NationalID).Find(&citizen).Error
	// if err_citizen != nil {
	// 	fmt.Println("err_citizen : ", err_citizen)
	// 	return
	// }
	// fmt.Println("citizen:", citizen)

	// Business logic to handle found citizens (e.g., update or insert)
	// if len(citizen) == 0 {
	// 	fmt.Println("No citizen found with NationalID:", response.Data.NationalID)

	// 	// Convert string to *string
	// 	nationalIDPtr := &response.Data.NationalID
	// 	FirstNamePtr  := &response.Data.FirstName
	// 	LastNamePtr  := &response.Data.LastName

	// 	// Insert logic for new citizen
	// 	newCitizen := models.Citizen{
	// 		TaxId:  nationalIDPtr,
	// 		FirstName:  FirstNamePtr,
	// 		LastName:  LastNamePtr,
	// 	}
	// 	// errInsert := db.Create(&newCitizen).Error
	// 	// if errInsert != nil {
	// 	// 	fmt.Println("Error inserting new citizen:", errInsert)
	// 	// 	return
	// 	// }
	// 	fmt.Println("Inserted new citizen:", newCitizen)
	// } else {
	// 	fmt.Println("Citizen already exists, updating records...")
	// 	now := time.Now()
	// 	data := map[string]interface{}{
	// 		"UpdatedAt": now,
	// 		"UpdatedBy": "test",
	// 	}
	// 	fmt.Println(" update data " , data)

	// 	for _, citizen := range citizen {
	// 		now := time.Now()
	// 		citizenUUID = &updatedByUUID
	// 		citizen.UpdatedAt = &now
	// 		citizen.UpdatedBy =citizenUUID
	// 		errUpdate := db.Save(&citizen).Error
	// 		if errUpdate != nil {
	// 			fmt.Println("Error updating citizen:", errUpdate)
	// 		} else {
	// 			fmt.Println("Updated citizen:", citizen)
	// 		}
	// 	}
	// }

	fmt.Println("----- LOG -----")

	// err_citizenUser := db.Limit(10).Find(&citizenUser).Error
	// if err_citizenUser != nil {
	// 	fmt.Println("err_citizenUser : ", err_citizenUser)

	// 	return
	// }

	// return assets, err

}
