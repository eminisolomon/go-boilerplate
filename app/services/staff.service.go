package services

import (
	"errors"
	"fmt"
	"inventory/app/database"
	"inventory/app/models"
	"inventory/app/utils"
)

func GetStaffs() (*[]models.User, error) {
	var err error
	var staff []models.User
	tx := database.PgDB.Where("role = ?", "staff").Find(&staff)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if len(staff) == 0 {
		err = errors.New("Staff not found")
	}

	return &staff, err
}

func GetStaff(userId string) (*models.User, error) {
	var err error
	var staff models.User
	tx := database.PgDB.Where("role = ?", "staff").First(&staff, userId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if staff.Cod == 0 {
		err = errors.New("Staff not found")
	}

	return &staff, err
}

func AddStaff(staff *models.User) (*models.User, error) {
	var err error

	existingUser, err := GetUserByEmail(staff.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", staff.Email)
	}

	existingUser, err = GetUserByUsername(staff.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with username %s already exists", staff.Username)
	}

	existingUser, err = GetUserByPhone(staff.Phone)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with phone %s already exists", staff.Phone)
	}

	staff.Role = "staff"

	hashedPassword, err := utils.HashPassword(staff.Password)
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, err
	}

	staff.Password = hashedPassword

	tx := database.PgDB.Create(&staff)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return staff, err
}

func UpdateStaff(newUser *models.User, staffId string) (*models.User, error) {
	var err error

	existingUser, err := GetUserByEmail(newUser.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil && existingUser.Cod != newUser.Cod {
		return nil, fmt.Errorf("user with email %s already exists", newUser.Email)
	}

	existingUser, err = GetUserByUsername(newUser.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil && existingUser.Cod != newUser.Cod {
		return nil, fmt.Errorf("user with username %s already exists", newUser.Username)
	}

	existingUser, err = GetUserByPhone(newUser.Phone)
	if err != nil {
		return nil, err
	}
	if existingUser != nil && existingUser.Cod != newUser.Cod {
		return nil, fmt.Errorf("user with phone %s already exists", newUser.Phone)
	}

	if newUser.Password != "" {
		hashedPassword, err := utils.HashPassword(newUser.Password)
		if err != nil {
			utils.WarningLog.Println(err.Error())
			return nil, err
		}

		newUser.Password = hashedPassword
	}

	staff, err := GetStaff(staffId)
	if err == nil {
		tx := database.PgDB.Model(staff).Updates(newUser)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}

	return staff, err
}

func DeleteStaff(userId string) error {
	var err error
	tx := database.PgDB.Where("role = ?", "staff").Delete(&models.User{}, userId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return err
}
