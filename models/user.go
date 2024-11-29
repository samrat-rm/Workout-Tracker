package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username       string        `gorm:"unique" json:"username" validate:"required"`
	Password       string        `json:"password" validate:"required"`
	WorkoutPlans   []WorkoutPlan `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"workout_plans"` // automatically remove related records when a User or WorkoutPlan is deleted
	HasWorkoutPlan bool          `json:"has_workout_plan"`
	ProgressLog    []ProgressLog `json:"progress_log"`
}

type Muscle int

const (
	Chest Muscle = iota
	Biceps
	Triceps
	Shoulders
	Delts
	BackMuscles
	ForeArms
	Abs
	Legs
)

type ExerciseCategory int

const (
	Cardio ExerciseCategory = iota
	Strength
	Flexibility
)

type Status int

const (
	NotStarted Status = iota
	InProgress
	Completed
	Quit
)

type Exercise struct {
	gorm.Model
	Name               string           `json:"name" validate:"required"`
	Description        string           `json:"description" validate:"required"`
	Category           ExerciseCategory `json:"category" validate:"required"`
	MuscleGroup        []Muscle         `json:"muscle_group" validate:"required"`
	PrimaryMuscleGroup Muscle           `json:"primary_muscle_group" validate:"required"`
}

type WorkoutPlan struct {
	gorm.Model
	UserID      uint       `json:"user_id" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Status      Status     `json:"status" validate:"required"`
	Exercises   []Exercise `json:"exercises"`
}

// TODO Add start and end date for workouts

type ProgressLog struct {
	gorm.Model
	UserID         uint      `json:"user_id" validate:"required"`         // Associate the progress log with a user
	WorkoutPlanID  uint      `json:"workout_plan_id" validate:"required"` // Link progress to a specific workout plan
	Date           time.Time `json:"date" validate:"required"`            // Log date
	ExerciseID     uint      `json:"exercise_id" validate:"required"`     // Log progress for a specific exercise
	Reps           int       `json:"reps"`                                // Number of repetitions
	Sets           int       `json:"sets" validate:"required"`            // Number of sets
	DurationInMins float64   `json:"duration_in_mins"`                    // Duration in minutes, for cardio or flexibility exercises
}

func (m Muscle) String() string {
	return [...]string{"Chest", "Biceps", "Triceps", "Shoulders", "Delts", "BackMuscles", "ForeArms", "Abs", "Legs"}[m]
}

func (e ExerciseCategory) String() string {
	return [...]string{"Cardio", "Strength", "Flexibility"}[e]
}
