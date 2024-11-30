package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type User struct {
	gorm.Model
	Username       string        `gorm:"unique" json:"username" validate:"required"`
	Password       string        `json:"password" validate:"required"`
	WorkoutPlans   []WorkoutPlan `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;" json:"workout_plans"` // Foreign Key relationship with WorkoutPlan
	HasWorkoutPlan bool          `json:"has_workout_plan"`
	ProgressLog    []ProgressLog `gorm:"foreignkey:UserID" json:"progress_log"` // Foreign Key relationship with ProgressLog
}

type ExerciseCategory int

const (
	Cardio ExerciseCategory = iota
	Strength
	Flexibility
)

type Exercise struct {
	gorm.Model
	Name               string           `json:"name" validate:"required"`
	Description        string           `json:"description" validate:"required"`
	Category           ExerciseCategory `json:"category" validate:"required"`
	MuscleGroup        pq.Int64Array    `gorm:"type:int[]" json:"muscle_group"`
	PrimaryMuscleGroup Muscle           `json:"primary_muscle_group" validate:"required"`
	WorkoutPlanID      uint             `json:"workout_plan_id"` // Foreign key reference to WorkoutPlan
}
type WorkoutPlan struct {
	gorm.Model
	UserID      uint       `json:"user_id" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Status      Status     `json:"status" validate:"required"`
	Exercises   []Exercise `gorm:"foreignkey:WorkoutPlanID" json:"exercises"` // Foreign key relationship with Exercise
}

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

func (e ExerciseCategory) String() string {
	return [...]string{"Cardio", "Strength", "Flexibility"}[e]
}
