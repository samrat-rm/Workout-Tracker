package models

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
	Name               string
	Description        string
	Category           ExerciseCategory
	MuscleGroup        []Muscle
	PrimaryMuscleGroup Muscle
}

type Workout struct {
	Exercises []Exercise
}
