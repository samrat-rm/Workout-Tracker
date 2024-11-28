package services

import (
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

type WorkoutPlanService interface {
	AddWorkoutPlanToUser(userID uint, workoutPlan *models.WorkoutPlan) error
	UpdateWorkoutPlanforUser(userID uint, updatedWorkoutPlan *models.WorkoutPlan) error
	UpdateWorkoutPlanStatusForUser(userID uint, workoutPlanID uint) error
	RemoveWorkoutPlanForUser(userID uint, workoutPlanID uint) error
	GetAllWorkoutPlansForUser(userID uint) error
	GetWorkoutPlanForUser(userID uint, workoutPlanID uint) error
}

type workoutPlanService struct {
	db *gorm.DB
}

// AddWorkoutPlanToUser implements WorkoutPlanService.
func (w *workoutPlanService) AddWorkoutPlanToUser(userID uint, workoutPlan *models.WorkoutPlan) error {
	panic("unimplemented")
}

// GetAllWorkoutPlansForUser implements WorkoutPlanService.
func (w *workoutPlanService) GetAllWorkoutPlansForUser(userID uint) error {
	panic("unimplemented")
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

func NewWorkoutPlanService(db *gorm.DB) WorkoutPlanService {
	return &workoutPlanService{db: db}
}
