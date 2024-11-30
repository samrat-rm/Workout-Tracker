package services

import (
	"fmt"
	"log"
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
	findUserByID(id uint) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (u *userService) CreateUser(user *models.User) error {
	if len(user.WorkoutPlans) > 0 {
		user.HasWorkoutPlan = true
	}
	if err := u.db.Create(&user).Error; err != nil {
		log.Printf("Error while saving user in DB : %s", err.Error())
		return err
	}
	return nil
}

func (u *userService) GetUser(id uint) (*models.User, error) {
	var user models.User
	// Use Preload to load associated WorkoutPlans
	if err := u.db.Preload("WorkoutPlans").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) GetUserByUsername(username string) (*models.User, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("invalid username")
	}
	var user models.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("Error while finding the user in DB : %s", err.Error())
		return nil, err
	}
	return &user, nil
}

func (u *userService) UpdateUser(id uint, user *models.User) error {
	existingUser, err := u.findUserByID(id)
	if err != nil {
		return err
	}

	if err := u.db.Model(existingUser).Updates(user).Error; err != nil {
		log.Printf("Error while updating user in DB: %s", err.Error())
		return err
	}

	return nil
}

func (u *userService) DeleteUser(id uint) error {
	user, err := u.findUserByID(id)
	if err != nil {
		return err
	}

	if err := u.db.Unscoped().Delete(&user, user.ID).Error; err != nil {
		log.Printf("Error while deleting user in DB: %s", err.Error())
		return err
	}

	return nil
}

func (u *userService) findUserByID(id uint) (*models.User, error) {
	if id == 0 {
		log.Printf("Invalid ID: %d", id)
		return nil, fmt.Errorf("invalid ID")
	}

	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		log.Printf("User not found: %s", err.Error())
		return nil, err
	}

	return &user, nil
}
