package models

import (
	"encoding/json"
	"fmt"
)

type Status int

const (
	NotStarted Status = iota
	InProgress
	Completed
	Quit
)

func (s *Status) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}

	switch statusStr {
	case "NotStarted":
		*s = NotStarted
	case "InProgress":
		*s = InProgress
	case "Completed":
		*s = Completed
	case "Quit":
		*s = Quit
	default:
		return fmt.Errorf("invalid status: %s", statusStr)
	}

	return nil
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Status) String() string {
	return [...]string{"NotStarted", "InProgress", "Completed", "Quit"}[s]
}
