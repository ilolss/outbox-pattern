package task

import (
	"encoding/json"
)

type Task struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (t *Task) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Task) FromJSON(data []byte) error {
	return json.Unmarshal(data, t)
} 