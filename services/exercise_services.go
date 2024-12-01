package services

import (
	"workout-tracker/models"

	"github.com/jinzhu/gorm"
)

type ExerciseService interface {
	CreateExercise(workoutPlanID uint, exercise models.Exercise) error
	UpdateExercise(exerciseID uint, updatedExercise models.Exercise) (models.Exercise, error)
	DeleteExercise(exerciseID uint) error
}

type exerciseService struct {
	db *gorm.DB
}

func NewExerciseService(db *gorm.DB) ExerciseService {
	return &exerciseService{db: db}
}

func (e *exerciseService) CreateExercise(workoutPlanID uint, exercise models.Exercise) error {
	exercise.WorkoutPlanID = workoutPlanID

	if err := e.db.Create(&exercise).Error; err != nil {
		return err
	}

	return nil
}

func (e *exerciseService) DeleteExercise(exerciseID uint) error {
	// Attempt to delete the exercise by ID
	if err := e.db.Delete(&models.Exercise{}, exerciseID).Error; err != nil {
		return err
	}
	return nil
}

func (e *exerciseService) UpdateExercise(exerciseID uint, updatedExercise models.Exercise) (models.Exercise, error) {
	var existingExercise models.Exercise

	// Fetch the exercise by its ID
	if err := e.db.First(&existingExercise, exerciseID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.Exercise{}, err
		}
		return models.Exercise{}, err
	}

	existingExercise.Name = updatedExercise.Name
	existingExercise.Description = updatedExercise.Description
	existingExercise.Category = updatedExercise.Category
	existingExercise.UpdatedAt = updatedExercise.UpdatedAt

	// Save the updated record
	if err := e.db.Save(&existingExercise).Error; err != nil {
		return models.Exercise{}, err
	}

	return existingExercise, nil
}
