package models

import (
	"encoding/json"
	"fmt"
)

type Muscle int

const (
	Chest Muscle = iota
	Biceps
	Triceps
	Shoulders
	Delts
	BackMuscles
	ForeArms
	Core
	Legs
)

func (m Muscle) String() string {
	return [...]string{"Chest", "Biceps", "Triceps", "Shoulders", "Delts", "BackMuscles", "ForeArms", "Core", "Legs"}[m]
}

func (m Muscle) MarshalJSON() ([]byte, error) {
	muscleStr := m.String()
	if muscleStr == "" {
		return nil, fmt.Errorf("invalid muscle type: %d", m)
	}
	return json.Marshal(muscleStr)
}

func (m *Muscle) UnmarshalJSON(data []byte) error {
	var muscleStr string
	if err := json.Unmarshal(data, &muscleStr); err != nil {
		return err
	}

	switch muscleStr {
	case "Chest":
		*m = Chest
	case "Biceps":
		*m = Biceps
	case "Triceps":
		*m = Triceps
	case "Shoulders":
		*m = Shoulders
	case "Delts":
		*m = Delts
	case "BackMuscles":
		*m = BackMuscles
	case "ForeArms":
		*m = ForeArms
	case "Core":
		*m = Core
	case "Legs":
		*m = Legs
	default:
		return fmt.Errorf("invalid muscle: %s", muscleStr)
	}

	return nil
}
