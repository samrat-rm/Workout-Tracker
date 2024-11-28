package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username       string        `gorm:"unique" json:"username"`
	Password       string        `json:"password"`
	Workouts       []WorkoutPlan `json:"workouts"`
	HasWorkoutPlan bool          `json:"has_workout_plan"`
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

type Exercise struct {
	Name               string           `json:"name"`
	Description        string           `json:"description"`
	Category           ExerciseCategory `json:"category"`
	MuscleGroup        []Muscle         `json:"muscle_group"`
	PrimaryMuscleGroup Muscle           `json:"primary_muscle_group"`
}

type WorkoutPlan struct {
	gorm.Model
	Exercises []Exercise `json:"exercises"`
}

func (m Muscle) String() string {
	return [...]string{"Chest", "Biceps", "Triceps", "Shoulders", "Delts", "BackMuscles", "ForeArms", "Abs", "Legs"}[m]
}

func (e ExerciseCategory) String() string {
	return [...]string{"Cardio", "Strength", "Flexibility"}[e]
}
