package utils

import (
	"errors"
	"regexp"
	"workout-tracker/models"
)

func validatePassword(password string) error {
	// Minimum length of 8 characters
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// At least one uppercase letter
	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// At least one lowercase letter
	lowercase := regexp.MustCompile(`[a-z]`)
	if !lowercase.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// At least one digit
	digit := regexp.MustCompile(`[0-9]`)
	if !digit.MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	// At least one special character
	specialChar := regexp.MustCompile(`[^a-zA-Z0-9]`)
	if !specialChar.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func ValidateUserCredentials(user *models.User) bool {
	if user.Username == "" || user.Password == "" {
		return false
	}
	return true
}
