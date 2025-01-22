package users_services

import (
	"citizen_system_back/database"
	"citizen_system_back/models"
	"errors"
	"time"

	// "citizen_system_back/models"
	asset_types "citizen_system_back/modules/users/types"
	"fmt"

	"github.com/google/uuid"
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
	StatusResponse StatusResponse      `json:"statusResponse"`
	Data           asset_types.Profile `json:"data"`
}

// func getUserByEmail() (*AuthServiceResponse, error) {
func GetUserByEmail() (*models.CitizenUser, error) {
	// 8fc7264f-e610-4541-934e-37bddcbaacf2

	var db = database.GetDB()
	var citizenUser *models.CitizenUser
	err := db.Find(&citizenUser).Where("id = ?", "8fc7264f-e610-4541-934e-37bddcbaacf2").Error
	fmt.Println("citizenUser : ", citizenUser)
	return citizenUser, err

}

func MockAuthUserData() *AuthServiceResponse {
	// Simulating successful response
	// id, err := uuid.Parse("8fc7264f-e610-4541-934e-37bddcbaacf2")
	// if err != nil {
	// 	panic("Invalid UUID format")
	// }
	mockResponse := &AuthServiceResponse{
		StatusResponse: StatusResponse{
			StatusCode:  "200",
			StatusDesc:  "OK",
			MessageCode: "00000",
			MessageTH:   "สำเร็จ",
			MessageEN:   "Success",
		},
		Data: asset_types.Profile{
			ID:           "8fc7264f-e610-4541-934e-37bddcbaacf2",
			Title:        "นาย",
			FirstName:    "user_first_name",
			LastName:     "user_last_name",
			PhoneNumber:  "0987654321",
			Picture:      nil,
			DocumentType: "nationalId",
			NationalID:   "5546751829457",
			PassportNo:   nil,
			Status:       "active",
			// Accounts: []Account{
			// 	{
			// 		Provider: "email",
			// 		Type:     "email",
			// 		Account:  "user_test@gmail.com",
			// 		Status:   "active",
			// 	},
			// 	{
			// 		Provider: "email",
			// 		Type:     "phone_number",
			// 		Account:  "0987654321",
			// 		Status:   "active",
			// 	},
			// },
		},
	}

	// Return the mock response
	return mockResponse
}

func UpSertProfileCitizenUser(profile asset_types.Profile) (models.CitizenUser, error) {

	fmt.Println(" **************************** UpSertProfileCitizenUser profile", profile)
	fmt.Println(" **************************** UpSertProfileCitizenUser profile.id", profile.ID)

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!! mock data !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	response := MockAuthUserData()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return asset_types.Profile{} , err
	// }

	if response.StatusResponse.StatusCode != "200" {
		return models.CitizenUser{}, fmt.Errorf("failed to upsert user: %w", errors.New("error test"))
	}

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!! mock data !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	// return response.Data, nil

	fmt.Println("response:", response)

	var db = database.GetDB()
	// var citizen []models.Citizen
	var citizenUser []models.CitizenUser

	// todo: find user in CitizenUser
	errCitizenUser := db.Find(&citizenUser).Where("id = ?", response.Data.ID).Error
	if errCitizenUser != nil {
		fmt.Println("errCitizenUser : ", errCitizenUser)
		return models.CitizenUser{}, fmt.Errorf("failed to upsert user: %w", errors.New("error test"))
	}

	fmt.Println("citizenUser :  ", citizenUser)

	// todo: build model data CitizenUser
	newCitizenUser := models.CitizenUser{}

	active := true
	CreatedBy := "system"
	now := time.Now()
	id, err := uuid.Parse(response.Data.ID)
	if err != nil {
		panic("Invalid UUID format")
	}
	newCitizenUser.Id = &id
	newCitizenUser.UserTel = &response.Data.PhoneNumber
	newCitizenUser.Active = &active
	newCitizenUser.UserName = &response.Data.FirstName
	newCitizenUser.UserLastName = &response.Data.LastName
	newCitizenUser.CreatedAt = &now
	newCitizenUser.CreatedBy = &CreatedBy

	fmt.Println("newCitizenUser :  ", &newCitizenUser)

	// todo: not found citizen user
	if len(citizenUser) == 0 {
		fmt.Println(" *********** citizenUser not found user data *********** ")

		// todo: open transaction
		tx := db.Begin()
		if tx.Error != nil {
			return models.CitizenUser{}, fmt.Errorf("failed to begin transaction: %w", tx.Error)
		}

		// todo: insert data to CitizenUser
		errCreateCitizenUser := tx.Create(&newCitizenUser).Error

		// todo: when can't create citizen user and error then rollback transaction and return error
		if errCreateCitizenUser != nil {
			fmt.Println("errCreateCitizenUser : ", errCreateCitizenUser)
			tx.Rollback() // Rollback the transaction on error
			return models.CitizenUser{}, fmt.Errorf("failed to upsert user: %w", errCreateCitizenUser)
		}

		// todo: Commit the transaction
		if errCommit := tx.Rollback().Error; errCommit != nil {
			return models.CitizenUser{}, fmt.Errorf("failed to commit transaction: %w", errCommit)
		}

		// return asset_types.Profile{}, nil
	} else {
		// found citizen user
		fmt.Println(" *********** citizenUser found user data *********** ")

		// errCitizen := tx.Find(&citizen).Where("tax_id = ?", response.Data.NationalID).Error
		// if errCitizen != nil {
		// 	fmt.Println("errCitizen : ", errCitizen)
		// 	tx.Rollback() // Rollback the transaction on error
		// 	return asset_types.Profile{}, fmt.Errorf("failed to upsert user: %w", errCitizen)
		// }

		// // db.First(&user)
		// // user.Name = "jinzhu 2"
		// // user.Age = 100
		// // db.Save(&user)

		// // todo: Commit the transaction
		// if errCommit := tx.Commit().Error; errCommit != nil {
		// 	return asset_types.Profile{}, fmt.Errorf("failed to commit transaction: %w", errCommit)
		// }

		// fmt.Println("citizen:", citizen)
	}
	fmt.Println(" 213 ******* newCitizenUser ******* :", &newCitizenUser)
	fmt.Println(" 214 ******* newCitizenUser.UserName ******* :", *newCitizenUser.UserName)

	// todo: update citizen (Field : citizen_id)
	go UpdateCitizenIdInCitizen(&newCitizenUser)

	return newCitizenUser, nil
}

func UpdateCitizenIdInCitizen(CitizenUser *models.CitizenUser) {

	// // todo: find user citizen
	// errFindCitizen := db.Find(&citizen).Where("citizen_id = ? ", response.Data.ID).Error
	// if errFindCitizen != nil {
	// 	fmt.Println("errFindCitizen : ", errFindCitizen)
	// 	return asset_types.Profile{}, fmt.Errorf("failed to upsert user: %w", errors.New("error test"))
	// }

	fmt.Println("data CitizenUser:", CitizenUser)
}

func GetProfile() string {

	// // todo: find user citizen
	// errFindCitizen := db.Find(&citizen).Where("citizen_id = ? ", response.Data.ID).Error
	// if errFindCitizen != nil {
	// 	fmt.Println("errFindCitizen : ", errFindCitizen)
	// 	return asset_types.Profile{}, fmt.Errorf("failed to upsert user: %w", errors.New("error test"))
	// }

	fmt.Println("data CitizenUser:")

	return "test"
}
