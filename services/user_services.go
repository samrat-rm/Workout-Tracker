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
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	AddWorkoutToUser(userID uint, workout *models.WorkoutPlan) error
	GetUserWorkouts(userID uint) ([]models.WorkoutPlan, error)
	RemoveWorkoutFromUser(userID uint, workoutID uint) error
}

type userService struct {
	db *gorm.DB
}

// AddWorkoutToUser implements UserService.
func (u *userService) AddWorkoutToUser(userID uint, workout *models.WorkoutPlan) error {
	panic("unimplemented")
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user *models.User) error { // !TODO checkif there is already user present
	if len(user.Workouts) > 0 {
		user.HasWorkoutPlan = true
	}
	if err := u.db.Create(&user).Error; err != nil {
		log.Printf("Error while saving user in DB : %s", err.Error())
		return err
	}
	return nil
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(id uint) error {
	panic("unimplemented")
}

// GetUser implements UserService.
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

// GetUserWorkouts implements UserService.
func (u *userService) GetUserWorkouts(userID uint) ([]models.WorkoutPlan, error) {
	panic("unimplemented")
}

// RemoveWorkoutFromUser implements UserService.
func (u *userService) RemoveWorkoutFromUser(userID uint, workoutID uint) error {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(user *models.User) error {
	panic("unimplemented")
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}
