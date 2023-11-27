package services

import (
	"fmt"
	"inventory/app/database"
	"inventory/app/dtos"
	"inventory/app/models"
	"inventory/app/utils"
)

func Signup(user *models.User) (*models.User, error) {
	var err error

	existingUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	existingUser, err = GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with username %s already exists", user.Username)
	}

	existingUser, err = GetUserByPhone(user.Phone)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with phone %s already exists", user.Phone)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, err
	}

	user.Password = hashedPassword

	tx := database.PgDB.Create(&user)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return user, err
}

func Signin(loginDto *dtos.LoginDto) (*models.User, *utils.TokenDetails, *utils.TokenDetails, error) {
	var user models.User
	err := database.PgDB.Where("username = ?", loginDto.Username).First(&user).Error
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, nil, nil, err
	}

	err = utils.VerifyPassword(user.Password, loginDto.Password)
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, nil, nil, err
	}

	// Issue Tokens
	accessToken, err := utils.IssueAccessToken(user)
	if err != nil {
		return nil, nil, nil, err
	}

	refreshToken, err := utils.IssueRefreshToken(user)
	if err != nil {
		return nil, nil, nil, err
	}

	return &user, accessToken, refreshToken, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.PgDB.Where("email = ?", email).First(&user).Error
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.PgDB.Where("username = ?", username).First(&user).Error
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := database.PgDB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		utils.WarningLog.Println(err.Error())
		return nil, err
	}

	return &user, nil
}
