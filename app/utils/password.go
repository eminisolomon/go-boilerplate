package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string) error {
	// Minimum length requirement
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	// Check for at least one uppercase letter
	if !containsUppercase(password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	// Check for at least one lowercase letter
	if !containsLowercase(password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	// Check for at least one digit
	if !containsDigit(password) {
		return fmt.Errorf("password must contain at least one digit")
	}

	return nil
}

func containsUppercase(s string) bool {
	return regexp.MustCompile(`[A-Z]`).MatchString(s)
}

func containsLowercase(s string) bool {
	return regexp.MustCompile(`[a-z]`).MatchString(s)
}

func containsDigit(s string) bool {
	return regexp.MustCompile(`[0-9]`).MatchString(s)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Unable to hash the password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
