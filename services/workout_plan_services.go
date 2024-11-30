package services

import (
	"fmt"
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

type WorkoutPlanService interface {
	CreateWorkoutPlanForUser(userID uint, workoutPlan *models.WorkoutPlan) error
	UpdateWorkoutPlanforUser(userID uint, updatedWorkoutPlan *models.WorkoutPlan) error
	UpdateWorkoutPlanStatusForUser(userID uint, workoutPlanID uint) error
	RemoveWorkoutPlanForUser(userID uint, workoutPlanID uint) error
	GetAllWorkoutPlansForUser(userID uint) ([]models.WorkoutPlan, error)
	GetWorkoutPlanForUser(userID uint, workoutPlanID uint) error
}

type workoutPlanService struct {
	db *gorm.DB
}

func NewWorkoutPlanService(db *gorm.DB) WorkoutPlanService {
	return &workoutPlanService{db: db}
}

func (w *workoutPlanService) CreateWorkoutPlanForUser(userID uint, workoutPlan *models.WorkoutPlan) error {
	var user models.User
	if err := w.db.First(&user, userID).Error; err != nil {
		fmt.Printf("Error User not found: %v\n", err)
		return err
	}

	workoutPlan.UserID = userID
	if err := w.db.Create(&workoutPlan).Error; err != nil {
		fmt.Printf("Error while creating workout plan: %v\n", err)
		return err
	}

	if err := w.db.Model(&user).Association("WorkoutPlans").Append(workoutPlan).Error; err != nil {
		fmt.Printf("Error while associating workout plan: %v\n", err)
		return err
	}

	user.HasWorkoutPlan = true
	if err := w.db.Save(&user).Error; err != nil {
		fmt.Printf("Error while saving workout plan: %v\n", err)
		return err
	}

	return nil
}

func (w *workoutPlanService) GetAllWorkoutPlansForUser(userID uint) ([]models.WorkoutPlan, error) {
	var workoutPlans []models.WorkoutPlan
	if err := w.db.Preload("Exercises").Where("user_id = ?", userID).Find(&workoutPlans).Error; err != nil {
		return nil, err
	}
	return workoutPlans, nil
}

// GetWorkoutPlanForUser implements WorkoutPlanService.
func (w *workoutPlanService) GetWorkoutPlanForUser(userID uint, workoutPlanID uint) error {
	panic("unimplemented")
}

// RemoveWorkoutPlanForUser implements WorkoutPlanService.
func (w *workoutPlanService) RemoveWorkoutPlanForUser(userID uint, workoutPlanID uint) error {
	panic("unimplemented")
}

// UpdateWorkoutPlanStatusForUser implements WorkoutPlanService.
func (w *workoutPlanService) UpdateWorkoutPlanStatusForUser(userID uint, workoutPlanID uint) error {
	panic("unimplemented")
}

// UpdateWorkoutPlanforUser implements WorkoutPlanService.
func (w *workoutPlanService) UpdateWorkoutPlanforUser(userID uint, updatedWorkoutPlan *models.WorkoutPlan) error {
	panic("unimplemented")
}
