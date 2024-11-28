package services

import (
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

type ExerciseService interface {
	GetAllExercises(userID uint, workoutPlanID uint) ([]models.Exercise, error)
	GetExercise(exerciseID uint) (models.Exercise, error)
	UpdateExercise(exerciseID uint, updatedExercise models.Exercise) (models.Exercise, error)
	CreateExercise(userID uint, workoutPlanID uint) error
	DeleteExercise(exerciseID uint) error
}

type exerciseService struct {
	db *gorm.DB
}

// CreateExercise implements ExerciseService.
func (e *exerciseService) CreateExercise(userID uint, workoutPlanID uint) error {
	panic("unimplemented")
}

// DeleteExercise implements ExerciseService.
func (e *exerciseService) DeleteExercise(exerciseID uint) error {
	panic("unimplemented")
}

// GetAllExercises implements ExerciseService.
func (e *exerciseService) GetAllExercises(userID uint, workoutPlanID uint) ([]models.Exercise, error) {
	panic("unimplemented")
}

// GetExercise implements ExerciseService.
func (e *exerciseService) GetExercise(exerciseID uint) (models.Exercise, error) {
	panic("unimplemented")
}

// UpdateExercise implements ExerciseService.
func (e *exerciseService) UpdateExercise(exerciseID uint, updatedExercise models.Exercise) (models.Exercise, error) {
	panic("unimplemented")
}

func NewExerciseService(db *gorm.DB) ExerciseService {
	return &exerciseService{db: db}
}
