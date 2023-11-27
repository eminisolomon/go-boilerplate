package services

import (
	"errors"
	"fmt"
	"inventory/app/database"
	"inventory/app/dtos"
	"inventory/app/models"
	"inventory/app/utils"
)

func GetUser(userId string) (*models.User, error) {
	var err error
	var user models.User
	tx := database.PgDB.First(&user, userId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if user.Cod == 0 {
		err = errors.New("User not found")
	}
	return &user, err
}

func GetUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	tx := database.PgDB.Find(&users)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if len(users) == 0 {
		err = errors.New("users not found")
	}

	return &users, err
}

func ChangePassword(userID string, changePasswordDto *dtos.ChangePasswordDto) error {
	if changePasswordDto.NewPassword != changePasswordDto.ConfirmPassword {
		return fmt.Errorf("new password and confirm password do not match")
	}

	if err := utils.ValidatePassword(changePasswordDto.NewPassword); err != nil {
		return err
	}

	user, err := GetUser(userID)
	if err != nil {
		return err
	}

	err = utils.VerifyPassword(user.Password, changePasswordDto.Password)
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return err
	}

	hashedPassword, err := utils.HashPassword(changePasswordDto.NewPassword)
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return err
	}

	user.Password = hashedPassword

	tx := database.PgDB.Save(&user)
	if tx.Error != nil {
		err := tx.Error
		utils.WarningLog.Println(err.Error())
		return err
	}

	return nil
}
