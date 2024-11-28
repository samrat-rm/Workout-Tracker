package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasssword(raw string) (string, error) {
	if err := validatePassword(raw); err != nil {
		log.Println("Invalid password, ", err.Error())
		return "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error while hashing password: ", err.Error())
		return "", err
	}
	return string(hashedPassword), nil
}
