package services

import (
	"fmt"
	"log"
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

// UserService defines the interface for user-related CRUD operations.
type UserService interface {
	CreateUser(user *models.User) error
	GetUser(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	db *gorm.DB
}

func (u *userService) DeleteUser(id uint) error {
	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		log.Printf("User not found: %s", err.Error())
		return err
	}

	if err := u.db.Delete(&user).Error; err != nil {
		log.Printf("Error while deleting user in DB: %s", err.Error())
		return err
	}

	return nil
}

func (u *userService) UpdateUser(id uint, user *models.User) error {

	var userInterface models.User
	if err := u.db.First(userInterface, user.ID).Error; err != nil {
		log.Printf("User not found: %s", err.Error())
		return err
	}
	if err := u.db.Save(&user).Error; err != nil {
		log.Printf("Error while deleting user in DB: %s", err.Error())
		return err
	}

	return nil
}

func (u *userService) AddWorkoutToUser(userID uint, workout *models.WorkoutPlan) error {
	panic("to do ")
}

func (u *userService) CreateUser(user *models.User) error { // !TODO checkif there is already user present
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
	if id == 0 {
		log.Printf("Invalid ID: %d", id)
		return nil, fmt.Errorf("invalid ID")
	}

	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		log.Printf("Error while finding the user in DB : %s", err.Error())
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

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}
