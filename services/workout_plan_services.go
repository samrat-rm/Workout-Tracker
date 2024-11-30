package services

import (
	"fmt"
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

type WorkoutPlanService interface {
	CreateWorkoutPlanForUser(userID uint, workoutPlan *models.WorkoutPlan) error
	UpdateWorkoutPlanforUser(userID uint, workoutPlanID uint, updatedWorkoutPlan *models.WorkoutPlan) error
	UpdateWorkoutPlanStatusForUser(userID uint, workoutPlanID uint, status models.Status) error
	RemoveWorkoutPlanForUser(userID uint, workoutPlanID uint) error
	GetAllWorkoutPlansForUser(userID uint) ([]models.WorkoutPlan, error)
	GetWorkoutPlanForUser(userID uint, workoutPlanID uint) (models.WorkoutPlan, error)
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

func (w *workoutPlanService) GetWorkoutPlanForUser(userID uint, workoutPlanID uint) (models.WorkoutPlan, error) {
	var workoutPlan models.WorkoutPlan
	if err := w.db.Preload("Exercises").
		Where("user_id = ? AND id = ?", userID, workoutPlanID).First(&workoutPlan).Error; err != nil {
		return workoutPlan, err
	}
	return workoutPlan, nil
}

func (w *workoutPlanService) RemoveWorkoutPlanForUser(userID uint, workoutPlanID uint) error {
	var workoutPlan models.WorkoutPlan
	if err := w.db.Where("user_id = ? AND id = ?", userID, workoutPlanID).First(&workoutPlan).Error; err != nil {
		return err
	}

	if err := w.db.Delete(&workoutPlan).Error; err != nil {
		return err
	}

	var user models.User
	if err := w.db.First(&user, userID).Error; err != nil {
		return err
	}

	user.HasWorkoutPlan = false
	if err := w.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (w *workoutPlanService) UpdateWorkoutPlanStatusForUser(userID uint, workoutPlanID uint, status models.Status) error {
	var workoutPlan models.WorkoutPlan
	if err := w.db.Where("user_id = ? AND id = ?", userID, workoutPlanID).First(&workoutPlan).Error; err != nil {
		return err
	}

	workoutPlan.Status = status
	if err := w.db.Save(&workoutPlan).Error; err != nil {
		return err
	}

	return nil
}

func (w *workoutPlanService) UpdateWorkoutPlanforUser(userID uint, workoutPlanID uint, updatedWorkoutPlan *models.WorkoutPlan) error {
	var workoutPlan models.WorkoutPlan
	if err := w.db.Where("user_id = ? AND id = ?", userID, workoutPlanID).First(&workoutPlan).Error; err != nil {
		return err
	}

	workoutPlan.Description = updatedWorkoutPlan.Description
	workoutPlan.Exercises = updatedWorkoutPlan.Exercises

	if err := w.db.Save(&workoutPlan).Error; err != nil {
		return err
	}

	return nil
}
